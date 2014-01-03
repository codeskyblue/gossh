package main

import (
	"fmt"
	"os"
	"strings"
)

func Debugf(format string, args ...interface{}) {
	if !*verbose {
		return
	}
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Fprintf(os.Stderr, format, args...)
}
