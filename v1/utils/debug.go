package utils

import (
	"fmt"
	"os"
	"strings"
)

const (
	DebugText = "[KOALA-DEBUG]"
)

func DebugPrint(format string, a ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	fmt.Fprint(os.Stderr, DebugText + format, a)
}

func DebugPrintf(format string, a ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	fmt.Fprintf(os.Stdout, DebugText + format, a)
}