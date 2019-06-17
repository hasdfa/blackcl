package blackParams

import (
	"github.com/hasdfa/blackcl"
)

type (
	int64Func = func(in []int64) (out []int64)

	RequestParamsInt64 struct {
		r           Request
		execDefault int64Func
	}

	RequestInt64 struct {
		r  RequestParamsInt64
		in []int64
	}
)

func (r Request) WithInt64(f int64Func) *RequestParamsInt64 {
	return &RequestParamsInt64{
		r:           r,
		execDefault: f,
	}
}

func (r *RequestParamsInt64) Make(args ...int64) *RequestInt64 {
	return &RequestInt64{
		r:  *r,
		in: args,
	}
}

func (r *RequestInt64) Execute(d *blackcl.Device) (res []int64, err error) {
	var input *blackcl.VectorInt64
	input, err = d.NewVectorInt64With(r.in)
	if err != nil {
		return
	}

	if err = <-d.Kernel(r.r.r.r.r.name).
		Global(r.r.r.r.global...).
		Local(r.r.r.local...).
		Run(input); err != nil {
		return
	}
	return input.Data()
}

func (r *RequestInt64) GO() []int64 {
	return r.r.execDefault(r.in)
}

func (r *RequestInt64) CanGO() bool {
	return r.r.execDefault != nil
}
