package repository

import (
	"context"
	model "mahasiswa/models/sql"

	"github.com/beego/beego/v2/adapter/orm"
)

func GetStudents(ctx context.Context) ([]model.Student, error) {
	// students := make([]*model.Student, 0)
	// res, err := o.Raw("SELECT * FROM student.mahasiswa").Exec()

	students := []model.Student{}
	o := orm.NewOrm()
	qs := o.QueryTable("mahasiswa")
	_, err := qs.All(&students)

	if err != nil {
		// log.Error(ctx, err)
		return nil, err
	}
	// log.Infoln(ctx, "Succesfully get data")
	return students, nil
}
