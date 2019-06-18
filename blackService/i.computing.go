package blackService

import (
	"github.com/hasdfa/blackcl/blackParams"
)

type IBlackService interface {
	Init() error
	Release() error

	ParseFolder(folder string) error
	ParseRaw(source string) error

	ExecI32(params *blackParams.RequestInt32) (data []int32, err error)
	ExecI64(params *blackParams.RequestInt64) (data []int64, err error)

	ExecF32(params *blackParams.RequestFloat32) (data []float32, err error)
	ExecF64(params *blackParams.RequestFloat64) (data []float64, err error)

	ExecMx32(params *blackParams.RequestMatrix32) (data []int32, err error)
	ExecMx64(params *blackParams.RequestMatrix64) (data []int64, err error)
}
