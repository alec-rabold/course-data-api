//go:build tools

package tools

import (
	_ "github.com/vektra/mockery/v2"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/cmd/goimports"
)
