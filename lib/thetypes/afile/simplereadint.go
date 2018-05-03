package afile

import (
	"io/ioutil"
	"fmt"
	"bytes"
)

func SimpleReadInt(path string) (r int, err error) {
	var b []byte
	b, err = ioutil.ReadFile(path)
	if err==nil {
		fmt.Sscan(string(b),&r)
	}
	return
}

func ScanTheEndOfTheFilePrefixedBy(prefix []byte, path string, maxlen uint) (r int, err error) {
	var b []byte
	b, err = ioutil.ReadFile(path)
	if err!=nil {
		return
	}
	if maxlen<uint(len(b)) {
		b = b[uint(len(b))-maxlen:]
	}
	i := bytes.Index(b, prefix)
	b = b[i+len(prefix):]
	fmt.Sscan(string(b), &r)
	return
}