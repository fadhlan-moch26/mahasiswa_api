package service

import (
	"context"
	"errors"
	"fmt"
	codec "mahasiswa/models/codec"
	model "mahasiswa/models/sql"
	"mahasiswa/repository"
)

func GetStudents(ctx context.Context) ([]*codec.Mahasiswa, error) {
	students, err := repository.GetStudents(ctx)
	if err != nil {
		// log.Error(ctx, err)
		return nil, err
	}
	// if len(students) < 0 {
	// 	return nil, log.Errorln("Data not found")
	// }
	v := make([]*codec.Mahasiswa, 0)
	for _, student := range students {
		v = append(v, &codec.Mahasiswa{
			Id:                student.Id,
			Nama:              student.Nama,
			Usia:              student.Usia,
			Gender:            student.Gender,
			TanggalRegistrasi: student.TanggalRegistrasi,
		})
	}
	return v, nil
}

func GetStudent(ctx context.Context, id string) (*codec.Mahasiswa, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	student, err := repository.GetStudent(ctx, id)
	if err != nil {
		return nil, errors.New("repository error")
	}
	return &codec.Mahasiswa{
		Id:                student.Id,
		Nama:              student.Nama,
		Usia:              student.Usia,
		Gender:            student.Gender,
		TanggalRegistrasi: student.TanggalRegistrasi,
	}, nil
}

func DeleteStudent(ctx context.Context, req *codec.Mahasiswa) (int, error) {
	id, err := repository.DeleteStudent(ctx, *req)
	if err != nil {
		return 0, err
	}
	req.Id = id
	return req.Id, nil
}

func UpdateStudent(ctx context.Context, req *codec.Mahasiswa) (int, error) {
	id, err := repository.UpdateStudent(ctx, *req)
	if err != nil {
		return 0, err
	}
	req.Id = id
	return req.Id, nil
}

func CreateStudent(ctx context.Context, req *codec.Mahasiswa) (int, error) {
	id, err := repository.CreateStudent(ctx, model.Mahasiswa{
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

func CreateMajor(ctx context.Context, req *codec.Major) (int, error) {
	id, err := repository.CreateMajor(ctx, model.Major{
		Id:          req.Id,
		NamaJurusan: req.NamaJurusan,
	})
	if err != nil {
		fmt.Println(ctx, err)
		// log.Error(ctx, err)
		return id, err
	}
	return id, nil
}

func CreateHobby(ctx context.Context, req *codec.Hobby) (int, error) {
	id, err := repository.CreateHobby(ctx, model.Hobby{
		Id:       req.Id,
		NamaHobi: req.NamaHobi,
	})
	if err != nil {
		fmt.Println(ctx, err)
		// log.Error(ctx, err)
		return id, err
	}
	return id, nil
}
