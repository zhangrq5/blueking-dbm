// Package dbutil TODO
package dbutil

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	"dbm-services/common/dbha/ha-module/log"

	_ "github.com/go-sql-driver/mysql"
)

// ConnMySQL connParam format: user:password@(ip:port)/dbName, %s:%s@(%s:%d)/%s
func ConnMySQL(connParam string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connParam)
	if err != nil {
		log.Logger.Errorf("connect mysql failed. err:%s", err.Error())
		return nil, nil
	}

	return db, nil
}

// GetMySQLErrorCode return 0 for success
func GetMySQLErrorCode(err error) uint16 {
	var sqlErr *mysql.MySQLError
	if errors.As(err, &sqlErr) {
		return sqlErr.Number
	}
	return 0
}
