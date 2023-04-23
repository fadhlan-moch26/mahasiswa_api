package repository

import (
	"context"
	"errors"
	"fmt"
	codec "mahasiswa/models/codec"
	model "mahasiswa/models/sql"
	"strconv"

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

func GetStudent(ctx context.Context, id string) (*model.Student, error) {
	id_int, _ := strconv.Atoi(id)
	student := model.Student{Id: id_int}
	o := orm.NewOrm()
	err := o.Read(&student)
	if errors.Is(orm.ErrNoRows, err) {
		return nil, errors.New("no data found")
	} else if err != nil {
		return nil, errors.New("general error")
	}
	return &student, nil
}

func DeleteStudent(ctx context.Context, req codec.Student) (int, error) {
	o := orm.NewOrm()
	err := o.Read(&req)
	if err != nil {
		return 0, errors.New("There is no id match in database")
	}
	id, err := o.Delete(&req)
	if err != nil {
		// log.Println(ctx, err, "Error occured when try to delete model", id)
		return 0, errors.New("Error occured when try to delete model")
	}
	fmt.Println(id, " Succesfully delete data")
	return int(id), nil
}
