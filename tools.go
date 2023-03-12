//go:build tools

package k8s

import (
	_ "github.com/gogo/protobuf/protoc-gen-gogo"
	_ "golang.org/x/tools/cmd/goimports"
	_ "k8s.io/code-generator"
)
