package blackParams

import (
	"github.com/hasdfa/blackcl"
)

type (
	matrix64WithParamsFunc = func(matrix []int64, in []int64) (out []int64)

	RequestParamsMatrix64 struct {
		Request
		execDefault matrix64WithParamsFunc
	}

	RequestMatrix64 struct {
		RequestParamsMatrix64
		matrix []int64
		in     []int64
	}
)

func (r Request) WithMatrix64(f matrix64WithParamsFunc) *RequestParamsMatrix64 {
	return &RequestParamsMatrix64{
		Request:     r,
		execDefault: f,
	}
}

func (r *RequestParamsMatrix64) Make(matrix []int64, in ...int64) *RequestMatrix64 {
	return &RequestMatrix64{
		RequestParamsMatrix64: *r,
		matrix:                matrix,
		in:                    in,
	}
}

func (r *RequestMatrix64) Execute(d *blackcl.Device) (res []int64, err error) {
	var matrix, input *blackcl.VectorInt64
	matrix, err = d.NewVectorInt64With(r.matrix)
	if err != nil {
		return
	}
	input, err = d.NewVectorInt64With(r.in)
	if err != nil {
		return
	}

	if err = <-d.Kernel(r.r.r.name).
		Global(r.r.global...).
		Local(r.local...).
		Run(matrix, input); err != nil {
		return
	}
	return input.Data()
}

func (r *RequestMatrix64) GO() []int64 {
	return r.execDefault(r.matrix, r.in)
}

func (r *RequestMatrix64) CanGO() bool {
	return r.execDefault != nil
}
