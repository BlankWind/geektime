package gorm

import (
	"time"

	"github.com/jinzhu/gorm"
)

type GormClient = gorm.DB

func NewGORM(info *GormConfig) (*GormClient, error) {
	db, err := gorm.Open(info.DBType, info.DSN)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	//db.SetLogger(gorm.Logger{})
	//db.LogMode(info.LogMode)
	db.DB().SetMaxOpenConns(info.MaxOpen)
	db.DB().SetMaxIdleConns(info.MaxIdle)
	db.DB().SetConnMaxLifetime(time.Second * time.Duration(info.ConnMaxLifetime))
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
