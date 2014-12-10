// Package testsweet provides a comprehensive testing framework.
package testsweet

import (
	"reflect"
	"testing"
)

var Equals = reflect.DeepEqual

func FailUnlessEqual(t *testing.T, a, b interface{}) {
	if !Equals(a, b) {
		t.Fatalf("!= %#v and %#v\n", a, b)
	}
}
