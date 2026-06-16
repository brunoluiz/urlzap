// Package xtest provides lightweight test helpers as a replacement for testify.
package xtest

import (
	"reflect"
	"testing"
)

// NoError fails the test if err is not nil.
func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error:\n%v", err)
	}
}

// Equal fails the test if expected and actual are not deeply equal.
func Equal(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("not equal:\nexpected: %+v\nactual:   %+v", expected, actual)
	}
}

// NotEqual fails the test if a and b are deeply equal.
func NotEqual(t *testing.T, a, b interface{}) {
	t.Helper()
	if reflect.DeepEqual(a, b) {
		t.Fatalf("unexpected equal values:\n%+v", a)
	}
}
