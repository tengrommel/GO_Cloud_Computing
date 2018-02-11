package models

import (
	_ "github.com/jinzhu/gorm"
	"github.com/rs/xid"
	"time"
)

type Base struct {
	ID 					string		`gorm:"primary_key"`
	CreatedAt		time.Time
	UpdateAt		time.Time
	DeletedAt		*time.Time		`sql:"index"`
}

func (b *Base)BeforeCreate() (err error) {
	b.ID = xid.New().String()
	return
}

type Todo struct {
	Base
	Title	string
	Done  bool
}
