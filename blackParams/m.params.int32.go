package blackParams

import (
	"github.com/hasdfa/blackcl"
)

type (
	int32Func = func(in []int32) (out []int32)

	RequestParamsInt32 struct {
		r           Request
		execDefault int32Func
	}

	RequestInt32 struct {
		r  RequestParamsInt32
		in []int32
	}
)

func (r Request) WithInt32(f int32Func) *RequestParamsInt32 {
	return &RequestParamsInt32{
		r:           r,
		execDefault: f,
	}
}

func (r *RequestParamsInt32) Make(args ...int32) *RequestInt32 {
	return &RequestInt32{
		r:  *r,
		in: args,
	}
}

func (r *RequestInt32) Execute(d *blackcl.Device) (res []int32, err error) {
	var input *blackcl.VectorInt32
	input, err = d.NewVectorInt32With(r.in)
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

func (r *RequestInt32) GO() []int32 {
	return r.r.execDefault(r.in)
}

func (r *RequestInt32) CanGO() bool {
	return r.r.execDefault != nil
}
