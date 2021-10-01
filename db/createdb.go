package db

import (
	"log"

	"github.com/teng231/workshop/pb"
	"xorm.io/xorm"
)

const (
	tblUser = "user"
)

func createTable(model interface{}, tblName string, engine *xorm.Engine) error {
	b, err := engine.IsTableExist(model)
	if err != nil {
		return err
	}
	if b {
		if err = engine.Sync2(model); err != nil {
			return err
		}
		return nil
	}
	if !b {
		if err := engine.CreateTables(model); err != nil {
			return err
		}
	}
	return nil
}

func (d *DB) CreateDb() error {
	var err error
	if err := createTable(&pb.User{}, tblUser, d.engine); err != nil {
		log.Print(err)
		return err
	}

	return err
}
