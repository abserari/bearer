package types

import (
	"github.com/bearer/bearer/new/detector/types"
	"github.com/bearer/bearer/new/language/tree"
)

type Object struct {
	Properties []Property
	// IsVirtual describes whether this object actually exists, or has
	// been detected as part of a variable name
	IsVirtual bool
}

type Property struct {
	Name   string
	Node   *tree.Node
	Object *types.Detection
}

type String struct {
	Value     string
	IsLiteral bool
}
