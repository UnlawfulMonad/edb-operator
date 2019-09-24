package edb

import (
	"database/sql"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
)

var (
	ErrInvalidName = errors.NewInternalError(fmt.Errorf("name is invalid"))
)

type mySQLConn struct {
	conn *sql.DB
}

func NewMySQL(adminUser, adminPassword, adminHost string) (ExternalDB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/?timeout=30s", adminUser, adminPassword, adminHost)

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
	if !isValidUsername(user) {
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

	return nil
}

func (c *mySQLConn) CreateDB(name, owner string) error {
	if !isValidUsername(owner) {
		return ErrInvalidName
	}

	if !isValidUsername(name) {
		return ErrInvalidName
	}

	users, err := c.listUsers()
	if err != nil {
		return err
	}

	haveUser := false
	for _, user := range users {
		if user == owner {
			haveUser = true
			break
		}
	}

	if !haveUser {
		// FIX ME
		return errors.NewServiceUnavailable("user does not exist")
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

	if _, err := c.conn.Exec(fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS %s`, name)); err != nil {
		return err
	}

	_, err = c.conn.Exec(fmt.Sprintf(`GRANT ALL ON %s.* TO '%s'@'%%'`, name, owner))
	if err != nil {
		return err
	}

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
