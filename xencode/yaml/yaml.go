package yaml

import (
	y "gopkg.in/yaml.v2"
)

func Unmarshal(in []byte, out interface{}) (err error) {
	return y.Unmarshal(in, out)
}

func Marshal(in interface{}) (out []byte, err error) {
	return y.Marshal(in)
}
