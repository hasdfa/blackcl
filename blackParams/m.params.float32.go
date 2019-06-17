package blackParams

import (
	"github.com/hasdfa/blackcl"
)

type (
	float32Func = func(in []float32) (out []float32)

	RequestParamsFloat32 struct {
		r           Request
		execDefault float32Func
	}

	RequestFloat32 struct {
		r  RequestParamsFloat32
		in []float32
	}
)

func (r Request) WithFloat32(f float32Func) *RequestParamsFloat32 {
	return &RequestParamsFloat32{
		r:           r,
		execDefault: f,
	}
}

func (r *RequestParamsFloat32) Make(args ...float32) *RequestFloat32 {
	return &RequestFloat32{
		r:  *r,
		in: args,
	}
}

func (r *RequestFloat32) Execute(d *blackcl.Device) (res []float32, err error) {
	var input *blackcl.VectorFloat32
	input, err = d.NewVectorFloat32With(r.in)
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

func (r *RequestFloat32) GO() []float32 {
	return r.r.execDefault(r.in)
}

func (r *RequestFloat32) CanGO() bool {
	return r.r.execDefault != nil
}
