package service

import (
	"context"
	"errors"
	"fmt"
	codec "mahasiswa/models/codec"
	model "mahasiswa/models/sql"
	"mahasiswa/repository"
)

func GetStudents(ctx context.Context) ([]*codec.Student, error) {
	students, err := repository.GetStudents(ctx)
	if err != nil {
		// log.Error(ctx, err)
		return nil, err
	}
	// if len(students) < 0 {
	// 	return nil, log.Errorln("Data not found")
	// }
	v := make([]*codec.Student, 0)
	for _, student := range students {
		v = append(v, &codec.Student{
			Id:                student.Id,
			Nama:              student.Nama,
			Usia:              student.Usia,
			Gender:            student.Gender,
			TanggalRegistrasi: student.TanggalRegistrasi,
		})
	}
	return v, nil
}

func GetStudent(ctx context.Context, id string) (*codec.Student, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	student, err := repository.GetStudent(ctx, id)
	if err != nil {
		return nil, errors.New("repository error")
	}
	return &codec.Student{
		Id:                student.Id,
		Nama:              student.Nama,
		Usia:              student.Usia,
		Gender:            student.Gender,
		TanggalRegistrasi: student.TanggalRegistrasi,
	}, nil
}

func DeleteStudent(ctx context.Context, req *codec.Student) (int, error) {
	id, err := repository.DeleteStudent(ctx, *req)
	if err != nil {
		return 0, err
	}
	req.Id = id
	return req.Id, nil
}

func UpdateStudent(ctx context.Context, req *codec.Student) (int, error) {
	id, err := repository.UpdateStudent(ctx, *req)
	if err != nil {
		return 0, err
	}
	req.Id = id
	return req.Id, nil
}

func CreateStudent(ctx context.Context, req *codec.Student) (int, error) {
	id, err := repository.CreateStudent(ctx, model.Student{
		Id:                req.Id,
		Nama:              req.Nama,
		Usia:              req.Usia,
		Gender:            req.Gender,
		TanggalRegistrasi: req.TanggalRegistrasi,
	})
	if err != nil {
		fmt.Println(ctx, err)
		// log.Error(ctx, err)
		return id, err
	}
	return id, nil
}
