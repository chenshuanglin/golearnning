package main

import (
    "reflect"
    "fmt"
)

type X int

type user struct {
    name string
    age int
}

type manager struct {
    user
    title string
}

func main() {
    var a X = 100
    t := reflect.TypeOf(a)
    println(t.Name(),t.Kind() == reflect.Int)
    
    fmt.Println(reflect.TypeOf(map[string]int{}).Elem())
    fmt.Println(reflect.TypeOf(map[string]int{}).Kind())
    
    var m manager
    t = reflect.TypeOf(&m)
    
    if t.Kind() == reflect.Ptr {
        fmt.Println(t.Elem())
        t = t.Elem()
    }
    
    for i := 0 ; i < t.NumField(); i++ {
        f := t.Field(i)
        fmt.Println(f.Name, f.Type, f.Offset)
        
        if f.Anonymous { //输出匿名字段结构
            for j := 0; j < f.Type.NumField(); j++ {
                af := f.Type.Field(j)
                fmt.Println(af.Name, af.Type, af.Offset)
            }
        }
    }
    name, _ := t.FieldByName("name")
    fmt.Println(name.Name, name.Type)
    
    age := t.FieldByIndex([]int{0,1})
    fmt.Println(age.Name, age.Type)
    
}
