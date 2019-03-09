package edb

import (
	"database/sql"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
)

var (
	ErrInvalidName = errors.NewInternalError(fmt.Errorf("name is invalid"))
)

type mySqlConn struct {
	conn *sql.DB
}

func NewMySQL(adminUser, adminPassword, adminHost string) (ExternalDB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/?timeout=30s", adminUser, adminPassword, adminHost)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	return &mySqlConn{conn: db}, nil
}

func (c *mySqlConn) listUsers() ([]string, error) {
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

func (c *mySqlConn) CreateUser(user, password string) error {
	users, err := c.listUsers()
	if err != nil {
		return err
	}

	for _, existing := range users {
		if existing == user {
			return nil
		}
	}

	_, err = c.conn.Exec(`INSERT INTO mysql.user(User, Host, Password) VALUES (?, ?, PASSWORD(?))`, user, "%", password)
	if err != nil {
		return err
	}

	return nil
}

func (c *mySqlConn) SetPassword(user, password string) error {
	_, err := c.conn.Exec(`UPDATE mysql.user SET Password = PASSWORD(?) WHERE user = ?`, password, user)
	if err != nil {
		return err
	}

	return nil
}

func (c *mySqlConn) CreateDB(name, owner string) error {
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
		}
	}

	if !haveUser {
		// FIX ME
		return errors.NewServiceUnavailable("user does not exist")
	}

	if _, err := c.conn.Exec(fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS %s`, name)); err != nil {
		return err
	}

	_, err = c.conn.Exec( fmt.Sprintf(`GRANT ALL ON %s TO %s@'%%'`, name, owner))

	if err != nil {
		return err
	}

	_, err = c.conn.Exec(`FLUSH PRIVILEGES`)
	if err != nil {
		return err
	}

	return nil
}

func (c *mySqlConn) Ping() error {
	return c.conn.Ping()
}

func (c *mySqlConn) Close() error {
	defer func() { c.conn = nil }()
	return c.conn.Close()
}
