package helper

import (
	"strings"
	"github.com/pborman/getopt/v2"
)

type Helper []string

func (h *Helper) Set(str string, opt getopt.Option) error {
	*h = append(*h, str)
	_ = opt
	return nil
}

func (h *Helper) String() string {
	return strings.Join(h.Array(), ", ")
}

func (h *Helper) Array() []string {
	return *h
}

func (h *Helper) ParseMultipleOptions() []string {
	return h.Array()
}
