package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
	fmt.Printf("%T\n", w)

	v := reflect.ValueOf(3)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())

	t = v.Type()
	fmt.Println(t.String())

	v = reflect.ValueOf(3)
	x := v.Interface()
	fmt.Println(x.(int))

	x := 2
	d := reflect.ValueOf(&x).Elem()
	px := d.Addr().Interface().(*int)
	*px = 3
	fmt.Println(x)

	d.Set(reflect.ValueOf(10))
	fmt.Println(x)
}
