# BlackCL

Black magic with black magic. These are highly opinionated OpenCL bindings for Go. It tries to make GPU computing easy, with some sugar abstraction, Go's concurency and channels.

Fork of https://github.com/microo8/blackcl with some extensions

Simple usage
```go
package main

import (
	"fmt"
	"log"
	
	"github.com/hasdfa/blackcl/blackParams"
	"github.com/hasdfa/blackcl/blackService"
)

//an complicated kernel
const kernelSource = `
__kernel void addOne(__global float* data) {
	const int i = get_global_id (0);
	data[i] += 1;
}
`

func main() {
	service := blackService.New()
	if err := service.Init(); err != nil {
		panic("could not initialize service: " + err.Error())
	}

	if err := service.ParseRaw(kernelSource); err != nil {
		panic("could not add program: " + err.Error())
	}

    // prints 
    //  [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
	runKernel(service, "addOne")
	
	// prints:
	//  could not execute: kernel with name 'addTwo' not found
	//  [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
	runKernel(service, "addTwo")
}

func golangFunction(in []float32) (out []float32) {
	log.Println("executing")
	
	for i := 0; i < len(in); i++ {
		in[i] = in[i] + 1
	}
	return in
}

func runKernel(service blackService.IBlackService, name string) {
	data := []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	request := blackParams.NewRequest(name).
		WithGlobal(len(data)).
		WithLocal(1).
		WithFloat32(golangFunction).
		Make(data...)

	newData, err := service.ExecF32(request)
	if err != nil {
		log.Println("could not execute: " + err.Error())
	}

	//prints out [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
	fmt.Println(newData)
}
```