package dicts_test

import (
	"github.com/vovkasm/go-benchdicts"
	"testing"
)

func testDict(t *testing.T, d dicts.Dict) {
	d.Set(dicts.Key{"a", "b", "c"}, 5)
	checkGet := func(k dicts.Key, ev int, eok bool) {
		v, ok := d.Get(k)
		if ok != eok {
			t.Errorf("d.Get(%v) expected (%v, %v), but got (%v, %v)", k, ev, eok, v, ok)
		} else {
			if v != ev {
				t.Errorf("d.Get(%v) expected (%v, %v), but got (%v, %v)", k, ev, eok, v, ok)
			}
		}
	}
	checkGet(dicts.Key{"a", "b", "c"}, 5, true)
	checkGet(dicts.Key{"a", "b", "x"}, 0, false)
	checkGet(dicts.Key{"a", "x", "c"}, 0, false)
	checkGet(dicts.Key{"x", "b", "c"}, 0, false)
}
func TestDict1(t *testing.T) {
	d := dicts.NewDict1()
	testDict(t, d)
}
func TestDict2(t *testing.T) {
	d := dicts.NewDict2()
	testDict(t, d)
}
func TestDict3(t *testing.T) {
	d := dicts.NewDict3()
	testDict(t, d)
}

var result int
var resultOk bool

var sKeys = []dicts.Key{
	dicts.Key{"aaa", "bbb", "ccc"},
	dicts.Key{"bbb", "ccc", "dddd"},
	dicts.Key{"bbb", "ccc", "xxxxxxx"},
	dicts.Key{"bbb", "cccddd", "dddd"},
	dicts.Key{"bbbaaa", "ccc", "dddd"},
}
var sVals = []int{1, 2, 3, 4, 5}

var ckeys = []dicts.Key{
	dicts.Key{"aaa", "bbb", "ccc"},
	dicts.Key{"bbb", "ccc", "dddd"},
	dicts.Key{"a", "b", "c"},
}

func benchDict(b *testing.B, d dicts.Dict) {
	for i, k := range sKeys {
		d.Set(k, sVals[i])
	}
	for n := 0; n < b.N; n++ {
		for _, k := range ckeys {
			result, resultOk = d.Get(k)
		}
	}
}
func BenchmarkDict1(b *testing.B) {
	d := dicts.NewDict1()
	benchDict(b, d)
}
func BenchmarkDict2(b *testing.B) {
	d := dicts.NewDict2()
	benchDict(b, d)
}
func BenchmarkDict3(b *testing.B) {
	d := dicts.NewDict3()
	benchDict(b, d)
}
