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

//VectorFloat32 is a memory buffer on device that holds []float32
type VectorFloat32 struct {
	buf *buffer
}

//Length the length of the vector
func (v *VectorFloat32) Length() int {
	return v.buf.size / 4
}

//Release releases the buffer on the device
func (v *VectorFloat32) Release() error {
	return v.buf.Release()
}

//NewVector allocates new vector buffer with specified length
func (d *Device) NewVectorFloat32(length int) (*VectorFloat32, error) {
	size := length * 4
	buf, err := newBuffer(d, size)
	if err != nil {
		return nil, err
	}
	return &VectorFloat32{buf: &buffer{memobj: buf, device: d, size: size}}, nil
}

//Copy copies the float32 data from host data to device buffer
//it's a non-blocking call, channel will return an error or nil if the data transfer is complete
func (v *VectorFloat32) Copy(data []float32) <-chan error {
	if v.Length() != len(data) {
		ch := make(chan error, 1)
		ch <- errors.New("vector length not equal to data length")
		return ch
	}
	return v.buf.copy(len(data)*4, unsafe.Pointer(&data[0]))
}

//Data gets float32 data from device, it's a blocking call
func (v *VectorFloat32) Data() ([]float32, error) {
	data := make([]float32, v.buf.size/4)
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
func (v *VectorFloat32) Map(k *Kernel) <-chan error {
	return k.Global(v.Length()).Local(1).Run(v)
}
