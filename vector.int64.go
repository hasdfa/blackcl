package blackcl

/*
#cgo CFLAGS: -I CL
#cgo !darwin LDFLAGS: -lOpenCL
#cgo darwin LDFLAGS: -framework OpenCL

#ifdef __APPLE__
#include <OpenCL/opencl.h>
#else
#include <CL/cl.h>
#endif
*/
import "C"
import (
	"errors"
	"unsafe"
)

//VectorInt64 is a memory buffer on device that holds []int64
type VectorInt64 struct {
	buf *buffer
}

//Length the length of the vector
func (v *VectorInt64) Length() int {
	return v.buf.size / int64CLSize
}

//Release releases the buffer on the device
func (v *VectorInt64) Release() error {
	return v.buf.Release()
}

//NewVectorInt64 allocates new vector buffer with specified length
func (d *Device) NewVectorInt64(length int) (*VectorInt64, error) {
	size := length * int64CLSize
	buf, err := newBuffer(d, size)
	if err != nil {
		return nil, err
	}
	return &VectorInt64{buf: &buffer{memobj: buf, device: d, size: size}}, nil
}

//NewVectorInt64 allocates new vector buffer with specified length
func (d *Device) NewVectorInt64With(data []int64) (*VectorInt64, error) {
	size := len(data) * int64CLSize
	buf, err := newBuffer(d, size)
	if err != nil {
		return nil, err
	}
	v := &VectorInt64{buf: &buffer{memobj: buf, device: d, size: size}}
	if err = <-v.Copy(data); err != nil {
		return nil, err
	}
	return v, nil
}

//Copy copies the float64 data from host data to device buffer
//it's a non-blocking call, channel will return an error or nil if the data transfer is complete
func (v *VectorInt64) Copy(data []int64) <-chan error {
	if v.Length() != len(data) {
		ch := make(chan error, 1)
		ch <- errors.New("vector length not equal to data length")
		return ch
	}
	return v.buf.copy(len(data)*int64CLSize, unsafe.Pointer(&data[0]))
}

//Data gets int64 data from device, it's a blocking call
func (v *VectorInt64) Data() ([]int64, error) {
	data := make([]int64, v.buf.size/int64CLSize)
	err := toErr(C.clEnqueueReadBuffer(
		v.buf.device.queue,
		v.buf.memobj,
		C.CL_TRUE,
		0,
		C.size_t(v.buf.size),
		unsafe.Pointer(&data[0]),
		0,
		nil,
		nil,
	))
	if err != nil {
		return nil, err
	}
	return data, nil
}

//Map applies an map kernel on all elements of the vector
func (v *VectorInt64) Map(k *Kernel) <-chan error {
	return k.Global(v.Length()).Local(1).Run(v)
}
