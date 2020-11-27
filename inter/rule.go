package inter

import "github.com/lanvard/support"

type Rule interface {
	Valid(present bool, value support.Value) bool
	Error(present bool, value support.Value) error
}
