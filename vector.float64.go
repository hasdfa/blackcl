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

//VectorFloat64 is a memory buffer on device that holds []float64
type VectorFloat64 struct {
	buf *buffer
}

//Length the length of the vector
func (v *VectorFloat64) Length() int {
	return v.buf.size / float64CLSize
}

//Release releases the buffer on the device
func (v *VectorFloat64) Release() error {
	return v.buf.Release()
}

//NewVector allocates new vector buffer with specified length
func (d *Device) NewVectorFloat64(length int) (*VectorFloat64, error) {
	size := length * float64CLSize
	buf, err := newBuffer(d, size)
	if err != nil {
		return nil, err
	}
	return &VectorFloat64{buf: &buffer{memobj: buf, device: d, size: size}}, nil
}

//NewVector allocates new vector buffer with specified length
func (d *Device) NewVectorFloat64With(data []float64) (*VectorFloat64, error) {
	size := len(data) * float64CLSize
	buf, err := newBuffer(d, size)
	if err != nil {
		return nil, err
	}
	v := &VectorFloat64{buf: &buffer{memobj: buf, device: d, size: size}}
	if err = <-v.Copy(data); err != nil {
		return nil, err
	}
	return v, nil
}

//Copy copies the float64 data from host data to device buffer
//it's a non-blocking call, channel will return an error or nil if the data transfer is complete
func (v *VectorFloat64) Copy(data []float64) <-chan error {
	if v.Length() != len(data) {
		ch := make(chan error, 1)
		ch <- errors.New("vector length not equal to data length")
		return ch
	}
	return v.buf.copy(len(data)*float64CLSize, unsafe.Pointer(&data[0]))
}

//Data gets float64 data from device, it's a blocking call
func (v *VectorFloat64) Data() ([]float64, error) {
	data := make([]float64, v.buf.size/float64CLSize)
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
func (v *VectorFloat64) Map(k *Kernel) <-chan error {
	return k.Global(v.Length()).Local(1).Run(v)
}
