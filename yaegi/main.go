package main

import (
	"func.xyz/tastes/yaegi/pkg"
	"github.com/traefik/yaegi/interp"
	"reflect"
)

const src = `
package foo

import "func.xyz/tastes/yaegi/pkg"

func Bar(s string) int {
	*pkg.Count++
	return *pkg.Count
}
`

const src0 = `
package foo

import "func.xyz/tastes/yaegi/pkg"

func Bar(s string) int {
	*pkg.Count++
	return *pkg.Count
}
`

func main() {
	i := interp.New(interp.Options{})

	i.Use(map[string]map[string]reflect.Value{
		"func.xyz/tastes/yaegi/pkg": {
			"Count": reflect.ValueOf(pkg.Count),
		},
	})

	_, err := i.Eval(src)
	if err != nil {
		panic(err)
	}

	v, err := i.Eval("foo.Bar")
	if err != nil {
		panic(err)
	}

	bar := v.Interface().(func(string) int)

	r := bar("Kung")
	println(r)

	*pkg.Count++

	s := bar("Kung")
	println(s)

	i0 := interp.New(interp.Options{})

	i0.Use(map[string]map[string]reflect.Value{
		"func.xyz/tastes/yaegi/pkg": {
			"Count": reflect.ValueOf(pkg.Count),
		},
	})

	_, err = i0.Eval(src0)
	if err != nil {
		panic(err)
	}

	v, err = i0.Eval("foo.Bar")
	if err != nil {
		panic(err)
	}

	bar = v.Interface().(func(string) int)

	r = bar("Kung")
	println(r)
}
