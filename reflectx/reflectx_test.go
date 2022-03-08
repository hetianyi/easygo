package reflectx

import (
	json "github.com/json-iterator/go"
	"testing"
)

func TestConfigure(t *testing.T) {
	a := &Config{
		Server: &Server{},
	}
	Start(&a)
	toString, _ := json.MarshalToString(a)
	println(toString)
}
