// Code generated by gop (Go+); DO NOT EDIT.

package main

import (
	"github.com/goplus/gop/test"
	"strconv"
	"testing"
)

type case_foo struct {
	test.Case
}
//line foo/foo_test.gox:1
func (this *case_foo) Main() {
//line foo/foo_test.gox:1:1
	if
//line foo/foo_test.gox:1:1
	v := foo(50); v != 100 {
//line foo/foo_test.gox:2:1
		this.T().Error("foo(50) ret: " + strconv.Itoa(v))
	}
//line foo/foo_test.gox:5:1
	this.T().Run("foo -10", func(t *testing.T) {
//line foo/foo_test.gox:6:1
		if foo(-10) != -20 {
//line foo/foo_test.gox:7:1
			t.Fatal("foo(-10) != -20")
		}
	})
//line foo/foo_test.gox:11:1
	this.T().Run("foo 0", func(t *testing.T) {
//line foo/foo_test.gox:12:1
		if foo(0) != 0 {
//line foo/foo_test.gox:13:1
			t.Fatal("foo(0) != 0")
		}
	})
}
func Test_foo(t *testing.T) {
//line foo/foo.gop:2:1
	test.Gopt_Case_TestMain(new(case_foo), t)
}
