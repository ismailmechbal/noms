// Generated by: main
// TypeWriter: noms
// Directive: +gen on Float32

package types

// DO NOT EDIT
//
// This file was generated by a tool.
// See http://clipperhouse.github.io/gen for details.

func (self Float32) Equals(other Value) bool {
	if other, ok := other.(Float32); ok {
		return self == other
	} else {
		return false
	}
}
