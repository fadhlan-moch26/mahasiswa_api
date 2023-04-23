package service

import (
	"context"
	codec "mahasiswa/models/codec"
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
