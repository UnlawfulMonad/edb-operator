package edb

import (
	// Import driver

	"fmt"
	"sync"

	// Import the MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"k8s.io/apimachinery/pkg/api/errors"

	// Import the PostgreSQL driver
	_ "github.com/lib/pq"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var (
	// ErrInvalidName is send out when the name provided doesn't comform to the
	// name restrictions.
	ErrInvalidName = errors.NewInternalError(fmt.Errorf("name is invalid"))

	// ErrUnsupportedPermission is what is returned when the user tries to pass a
	// a permission that isn't supported by the system.
	ErrUnsupportedPermission = errors.NewInternalError(fmt.Errorf("the passed permission is invalid"))
)

// An ExternalDB is the generic interface for handling database connections.
// It's effectively an easy to use wrapper around the various database drivers.
type ExternalDB interface {
	CreateUser(name, password string) error
	CreateDB(name string) error
	SetPassword(name, password string) error
	Grant(permission, to, on string) error
	Ping() error
	Close() error
}

var (
	externalDatabasesMutex = &sync.Mutex{}
	externalDatabases      = make(map[string]ExternalDB)
)

// LookupExternalDatabase finds a registered external database and outputs it
// if it exists. Returns nil if the named database does not exist.
func LookupExternalDatabase(name string) ExternalDB {
	externalDatabasesMutex.Lock()
	defer externalDatabasesMutex.Unlock()

	db, ok := externalDatabases[name]
	if !ok {
		return nil
	}

	return db
}

// AddOrUpdateExternalDatabase registers an external database.
// the passed external database cannot be nil. If the database already
// exists it will update the entry.
func AddOrUpdateExternalDatabase(name string, db ExternalDB) {
	if db == nil {
		panic("cannot add a nil ExternalDB. This is probably a bug. Please report it at https://github.com/UnlawfulMonad/edb-operator/issues")
	}

	externalDatabasesMutex.Lock()
	defer externalDatabasesMutex.Unlock()

	if old, ok := externalDatabases[name]; ok {
		if err := old.Close(); err != nil {
			logf.Log.Error(err, "failed to close old connection")
		}
	}

	externalDatabases[name] = db
}

// RemoveExternalDatabase closes a connection and deregisters an
// external database.
func RemoveExternalDatabase(name string) {
	externalDatabasesMutex.Lock()
	defer externalDatabasesMutex.Unlock()

	// Only delete if it actually exists
	value, ok := externalDatabases[name]
	if ok {
		if err := value.Close(); err != nil {
			logf.Log.Error(err, "failed to close database")
		}

		delete(externalDatabases, name)
	}
}
