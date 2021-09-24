package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var eg *xorm.EngineGroup

func NewEngine() (*xorm.EngineGroup, error) {
	cons := []string{
		"root:root@(127.0.0.1:3306)/grpc_todolist?charset=utf8",
		"root:root@(127.0.0.1:3306)/grpc_todolist?charset=utf8",
		"root:root@(127.0.0.1:3306)/grpc_todolist?charset=utf8",
	}

	var err error
	eg, err = xorm.NewEngineGroup("mysql", cons, xorm.RoundRobinPolicy())
	return eg, err
}
