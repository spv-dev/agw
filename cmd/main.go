package main

import (
	"fmt"
	"runtime/debug"

	"github.com/spv-dev/agw/internal/buildinfo"
)

func init() {
	bi, ok := debug.ReadBuildInfo()
	if ok {
		buildinfo.ServiceName = bi.Path
		buildinfo.ServiceVersion = bi.Main.Version
	}
}

func main() {
	fmt.Printf("Name:%s Version:%s", buildinfo.ServiceName, buildinfo.ServiceVersion)
}
