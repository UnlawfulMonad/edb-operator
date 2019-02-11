package edb

import (
	// Import driver
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"sync"
)

type ExternalDB interface {
	CreateUser(name, password string) error
	CreateDB(name, owner string) error
	Ping() error
	Close() error
}

var (
	externalDatabasesMutex = &sync.Mutex{}
	externalDatabases      = make(map[string]ExternalDB)
)

func LookupExternalDatabase(name, namespace string) ExternalDB {
	name = nsString(name, namespace)

	externalDatabasesMutex.Lock()
	defer externalDatabasesMutex.Unlock()

	db, ok := externalDatabases[name]
	if !ok {
		return nil
	}

	return db
}

func AddOrUpdateExternalDatabase(name, namespace string, db ExternalDB) {
	if db == nil {
		panic("cannot add a nil ExternalDB. This is probably a bug. Please report it at https://github.com/UnlawfulMonad/edb-operator/issues")
	}

	externalDatabasesMutex.Lock()
	externalDatabases[nsString(name, namespace)] = db
	externalDatabasesMutex.Unlock()
}

func RemoveExternalDatabase(name, namespace string) {
	name = nsString(name, namespace)

	externalDatabasesMutex.Lock()
	delete(externalDatabases, name)
	externalDatabasesMutex.Unlock()
}

func nsString(name, namespace string) string {
	return name + "/" + namespace
}

var (
	userValidateRegexp = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)
)

func isValidUsername(name string) bool {
	return userValidateRegexp.MatchString(name)
}
