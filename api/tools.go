//go:build tools
// +build tools

package tools

// tools is a dummy package that will be ignored for builds, but included for dependencies
import (
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
)
