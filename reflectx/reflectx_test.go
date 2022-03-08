package reflectx

import (
	"testing"
)

func TestConfigure(t *testing.T) {
	a := &Config{
		Server: &Server{},
	}
	Start(a)
}
