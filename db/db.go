package db

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/teng231/workshop/pb"
	"xorm.io/xorm"
)

type DB struct {
	engine *xorm.Engine
}

func (d *DB) ConnectDb(sqlPath, dbName string) error {
	sqlConnStr := fmt.Sprintf("%s/%s?charset=utf8", sqlPath, dbName)
	engine, err := xorm.NewEngine("mysql", sqlConnStr)
	if err != nil {
		return err
	}
	log.Print("Connected to: ", sqlConnStr)
	d.engine = engine
	d.engine.ShowSQL(false)
	return err
}

// -------- USER ---------

// InsertUser normal insert to db
func (d *DB) InsertUser(user *pb.User) error {
	affected, err := d.engine.Insert(user)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("can_not_insert")
	}
	return nil
}

// UpdateUser normal update to db
func (d *DB) UpdateUser(updator, selector *pb.User) error {
	affected, err := d.engine.Update(updator, selector)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("can_not_update")
	}
	return nil
}

// FindUser get single row
func (d *DB) FindUser(in *pb.User) (*pb.User, error) {
	ss := d.engine.Table(tblUser)
	found, err := ss.Get(in)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("not found")
	}

	return in, nil
}

// listUsersQuery list partner query with id, state, name, from, to usereter
func (d *DB) listUsersQuery(req *pb.UserRequest) *xorm.Session {
	ss := d.engine.Table(tblUser)
	return ss
}

// CountListUsers count by conditions
func (d *DB) CountListUsers(in *pb.UserRequest) (int64, error) {
	return d.listUsersQuery(in).Count()
}

// ListUsers list by conditions
func (d *DB) ListUsers(in *pb.UserRequest) ([]*pb.User, error) {
	var users []*pb.User
	ss := d.listUsersQuery(in)
	// if in.GetLimit() != 0 {
	// 	ss.Limit(int(in.GetLimit()), int(in.GetLimit())*int(in.GetOffset()))
	// }
	if err := ss.Desc("id").Find(&users); err != nil {
		return nil, err
	}

	return users, nil
}

// -------- END USER ---------
