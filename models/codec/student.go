package codec

import "time"

type Student struct {
	Id                string    `json:"id"`
	Nama              string    `json:"nama"`
	Usia              int       `json:"usia"`
	Gender            int       `json:"gender"`
	TanggalRegistrasi time.Time `json:"tanggalRegistrasi"`
}
