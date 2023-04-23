package sql

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

const tableNameStudent = "student"
const tableNameHobby = "Hobby"
const tableNameMajor = "Major"

func init() {
	orm.RegisterModel(new(Student))
	orm.RegisterModel(new(Hobby))
	orm.RegisterModel(new(Major))
}

type Student struct {
	Id                int `orm:"pk"`
	Nama              string
	Usia              int
	Gender            int
	TanggalRegistrasi time.Time
}

type Major struct {
	Id          int `orm:"pk"`
	NamaJurusan string
}

type Hobby struct {
	Id       int `orm:"pk"`
	NamaHobi string
}

func (m *Student) TableName() string {
	return tableNameStudent
}

func (m *Hobby) TableName() string {
	return tableNameHobby
}

func (m *Major) TableName() string {
	return tableNameMajor
}
