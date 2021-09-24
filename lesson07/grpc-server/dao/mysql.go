package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var eg *xorm.EngineGroup

func NewEngine() (*xorm.EngineGroup, error) {
	cons := []string{
		"mysql://mysql:root@localhost:3306/to_do?sslmode=disable;",
		"mysql://mysql:root@localhost:3306/to_do?sslmode=disable;",
		"mysql://mysql:root@localhost:3306/to_do?sslmode=disable",
	}

	var err error
	eg, err = xorm.NewEngineGroup("mysql", cons, xorm.RoundRobinPolicy())
	return eg, err
}
