package blackService

import (
	"errors"
	"fmt"
	"github.com/hasdfa/blackcl"
	"io/ioutil"
	"os"
	"path/filepath"
)

type blackService struct {
	device *blackcl.Device
}

func New() IBlackService {
	return &blackService{}
}

func (s *blackService) Init() (err error) {
	s.device, err = blackcl.GetDefaultDevice()
	if err != nil {
		return errors.New("no opencl device: " + err.Error())
	}
	return
}

func (s *blackService) ParseRaw(source string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()

	s.device.AddProgram(string(source))
	return
}

func (s *blackService) ParseFolder(folder string) (err error) {
	return filepath.Walk(folder, func(path string, info os.FileInfo, inputErr error) (err error) {
		if info.IsDir() {
			return nil
		}

		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("%s", r)
			}
		}()

		bts, err := ioutil.ReadFile(path)
		if err != nil {
			return
		}

		s.device.AddProgram(string(bts))
		return
	})
}

func (s *blackService) Release() error {
	return s.device.Release()
}

func (s *blackService) parseFile(p string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()

	bts, err := ioutil.ReadFile(p)
	if err != nil {
		return
	}

	s.device.AddProgram(string(bts))
	return
}
