package thetypes

import "strings"

func IsSysThermalTemp(s string) bool {
	return strings.HasPrefix(s, "/sys/class/thermal/") &&
		strings.HasSuffix(s, "temp")
}

func IsDS18B20(s string) bool {
	return s=="/sys/bus/w1/devices/28-0417c1ca72ff/w1_slave"
}

type Autodetect = func(string) bool
type Read = func(string) (int,error)
type BundlePiece struct {
	Autodetect
	Read
}