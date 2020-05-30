package db

import (
	"fmt"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

func InitDB(user string, pass string, host string, data string) (sqlbuilder.Database, error) {
	settings := mysql.ConnectionURL{
		User:     user,
		Password: pass,
		Host:     host,
		Database: data,
	}
	sess, err := mysql.Open(settings)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func VerifyDB(sess sqlbuilder.Database, tables []string) error {
	if _, err := sess.Query("SELECT * FROM Users"); err != nil {
		return err
	}
	for _, table := range tables {
		if _, err := sess.Query(fmt.Sprintf("SELECT * FROM %s", table)); err != nil {
			return err
		}
	}
	return nil
}