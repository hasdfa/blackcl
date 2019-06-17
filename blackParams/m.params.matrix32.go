package blackParams

import (
	"github.com/hasdfa/blackcl"
)

type (
	matrix32WithParamsFunc = func(matrix []int32, in []int32) (out []int32)

	RequestParamsMatrix32 struct {
		r           Request
		execDefault matrix32WithParamsFunc
	}

	RequestMatrix32 struct {
		r      RequestParamsMatrix32
		matrix []int32
		in     []int32
	}
)

func (r Request) WithMatrix32(f matrix32WithParamsFunc) *RequestParamsMatrix32 {
	return &RequestParamsMatrix32{
		r:           r,
		execDefault: f,
	}
}

func (r *RequestParamsMatrix32) Make(matrix []int32, in ...int32) *RequestMatrix32 {
	return &RequestMatrix32{
		r:      *r,
		matrix: matrix,
		in:     in,
	}
}

func (r *RequestMatrix32) Execute(d *blackcl.Device) (res []int32, err error) {
	var matrix, input *blackcl.VectorInt32
	matrix, err = d.NewVectorInt32With(r.matrix)
	if err != nil {
		return
	}
	input, err = d.NewVectorInt32With(r.in)
	if err != nil {
		return
	}

	if err = <-d.Kernel(r.r.r.r.r.name).
		Global(r.r.r.r.global...).
		Local(r.r.r.local...).
		Run(matrix, input); err != nil {
		return
	}
	return input.Data()
}

func (r *RequestMatrix32) GO() []int32 {
	return r.r.execDefault(r.matrix, r.in)
}

func (r *RequestMatrix32) CanGO() bool {
	return r.r.execDefault != nil
}
