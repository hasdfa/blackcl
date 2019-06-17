package blackService

import (
	"fmt"
	"github.com/hasdfa/blackcl/blackParams"
)

func (s *blackService) ExecI32(params *blackParams.RequestInt32) (data []int32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
			if params != nil && params.CanGO() {
				data = params.GO()
			}
		}
	}()
	if data, err = params.Execute(s.device); err != nil {
		data = params.GO()
	}
	return
}

func (s *blackService) ExecI64(params *blackParams.RequestInt64) (data []int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
			if params != nil && params.CanGO() {
				data = params.GO()
			}
		}
	}()
	if data, err = params.Execute(s.device); err != nil {
		data = params.GO()
	}
	return
}

func (s *blackService) ExecF32(params *blackParams.RequestFloat32) (data []float32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
			if params != nil && params.CanGO() {
				data = params.GO()
			}
		}
	}()
	if data, err = params.Execute(s.device); err != nil {
		data = params.GO()
	}
	return
}

func (s *blackService) ExecF64(params *blackParams.RequestFloat64) (data []float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
			if params != nil && params.CanGO() {
				data = params.GO()
			}
		}
	}()
	if data, err = params.Execute(s.device); err != nil {
		data = params.GO()
	}
	return
}

// specific

func (s *blackService) ExecMx32(params *blackParams.RequestMatrix32) (data []int32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
			if params != nil && params.CanGO() {
				data = params.GO()
			}
		}
	}()
	if data, err = params.Execute(s.device); err != nil {
		data = params.GO()
	}
	return
}

func (s *blackService) ExecMx64(params *blackParams.RequestMatrix64) (data []int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
			if params != nil && params.CanGO() {
				data = params.GO()
			}
		}
	}()
	if data, err = params.Execute(s.device); err != nil {
		data = params.GO()
	}
	return
}
