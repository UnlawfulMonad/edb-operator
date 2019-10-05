package edb

import (
	"database/sql"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var (
	// ErrInvalidName is send out when the name provided doesn't comform to the
	// name restrictions.
	ErrInvalidName = errors.NewInternalError(fmt.Errorf("name is invalid"))

	mlog = logf.Log.WithName("edb_mysql")
)

type mySQLConn struct {
	conn *sql.DB
}

// NewMySQL creates a connection to a MySQL database.
func NewMySQL(adminUser, adminPassword, adminHost, dbName string) (ExternalDB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", adminUser, adminPassword, adminHost, dbName)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	return &mySQLConn{conn: db}, nil
}

func (c *mySQLConn) listDatabases() ([]string, error) {
	dbs := make([]string, 0)
	rows, err := c.conn.Query("SHOW DATABASES")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		dbs = append(dbs, name)
	}

	return dbs, nil
}

func (c *mySQLConn) listUsers() ([]string, error) {
	users := make([]string, 0)

	rows, err := c.conn.Query("SELECT User FROM mysql.user")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		users = append(users, name)
	}

	return users, nil
}

func (c *mySQLConn) CreateUser(user, password string) error {

	// Validate username
	if !isValidIdentifier(user) {
		return ErrInvalidName
	}

	users, err := c.listUsers()
	if err != nil {
		return err
	}

	for _, existing := range users {
		if existing == user {
			return nil
		}
	}

	// We've already validated the username so this is safe.
	fullUsername := `"` + user + `"@%`
	_, err = c.conn.Exec(`CREATE USER ` + fullUsername)
	if err != nil {
		return err
	}

	// Set the user's password.
	err = c.SetPassword(user, password)
	if err != nil {
		return err
	}

	c.flushPrivs()

	return nil
}

func (c *mySQLConn) SetPassword(user, password string) error {
	_, err := c.conn.Exec(`UPDATE mysql.user SET Password = PASSWORD(?) WHERE user = ?`, password, user)
	if err != nil {
		return err
	}

	c.flushPrivs()

	return nil
}

func (c *mySQLConn) CreateDB(name string) error {
	log := mlog.WithValues("Database", name)

	if !isValidIdentifier(name) {
		return ErrInvalidName
	}

	dbs, err := c.listDatabases()
	if err != nil {
		return err
	}

	haveDb := false
	for _, db := range dbs {
		if db == name {
			haveDb = true
			break
		}
	}

	if haveDb {
		return nil
	}

	log.Info("creating database")
	if _, err := c.conn.Exec(fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS %s`, name)); err != nil {
		return err
	}

	//_, err = c.conn.Exec(fmt.Sprintf(`GRANT ALL ON %s.* TO '%s'@'%%'`, name, owner))
	//if err != nil {
	//	return err
	//}

	log.Info("flushing privilages")
	c.flushPrivs()

	return nil
}

func (c *mySQLConn) flushPrivs() {
	_, err := c.conn.Exec(`FLUSH PRIVILEGES`)
	if err != nil {
		panic(err)
	}
}

func (c *mySQLConn) Ping() error {
	return c.conn.Ping()
}

func (c *mySQLConn) Close() error {
	defer func() { c.conn = nil }()
	return c.conn.Close()
}
