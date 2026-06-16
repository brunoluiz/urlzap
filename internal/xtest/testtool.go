package xtest

import (
	"reflect"
	"testing"
)

func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error:\n%v", err)
	}
}

func Equal(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("not equal:\nexpected: %+v\nactual:   %+v", expected, actual)
	}
}

func NotEqual(t *testing.T, a, b interface{}) {
	t.Helper()
	if reflect.DeepEqual(a, b) {
		t.Fatalf("unexpected equal values:\n%+v", a)
	}
}
