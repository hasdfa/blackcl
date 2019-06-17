package blackParams

import (
	"github.com/hasdfa/blackcl"
)

type (
	float64Func = func(in []float64) (out []float64)

	RequestParamsFloat64 struct {
		r           Request
		execDefault float64Func
	}

	RequestFloat64 struct {
		r  RequestParamsFloat64
		in []float64
	}
)

func (r Request) WithFloat64(f float64Func) *RequestParamsFloat64 {
	return &RequestParamsFloat64{
		r:           r,
		execDefault: f,
	}
}

func (r *RequestParamsFloat64) Make(args ...float64) *RequestFloat64 {
	return &RequestFloat64{
		r:  *r,
		in: args,
	}
}

func (r *RequestFloat64) Execute(d *blackcl.Device) (res []float64, err error) {
	var input *blackcl.VectorFloat64
	input, err = d.NewVectorFloat64With(r.in)
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

func (r *RequestFloat64) GO() []float64 {
	return r.r.execDefault(r.in)
}

func (r *RequestFloat64) CanGO() bool {
	return r.r.execDefault != nil
}
