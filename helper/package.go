package helper

import (
	"runtime"
	"strings"
)

func GetPackageName() string {
	pc, _, _, _ := runtime.Caller(1)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	pkg := strings.Split(parts[len(parts)-1], ".")
	return pkg[0]
}
