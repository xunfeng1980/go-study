package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
)

func TestMake(t *testing.T) {
	slice := make([]string, 5)
	assert.Equal(t, len(slice), 5)
	assert.Equal(t, cap(slice), 5)
}

func TestMakeLenNotEqCap(t *testing.T) {
	slice := make([]string, 3, 5)
	assert.Equal(t, len(slice), 3)
	assert.Equal(t, cap(slice), 5)
}

func TestMakeCapLTLen(t *testing.T) {
	// Compile Err:./app_test.go:22:26: invalid argument: length and capacity swapped
	//slice := make([]string, 5, 3)
	//assert.Equal(t, len(slice), 3)
	//assert.Equal(t, cap(slice), 5)
}

func TestNilSeAndEmptySe(t *testing.T) {
	var nilSe []int
	emptySe := make([]int, 0)
	assert.Equal(t, cap(nilSe), cap(emptySe))
	assert.Equal(t, len(nilSe), len(emptySe))
}

func TestCreateFromSe(t *testing.T) {
	se := []int{1, 2, 3, 4, 5}
	newSe := se[1:3]
	assert.Equal(t, []int{2, 3}, newSe)
	assert.Equal(t, cap(newSe), 5-1)
	assert.Equal(t, len(newSe), 3-1)

	newSe[1] = 3
	assert.Equal(t, newSe[1], se[2])
}

func TestNewMap(t *testing.T) {
	_ = make(map[string]string)

	//panic: assignment to entry in nil map
	//var d1 map[string]string
	//d1["Red"] = "#FF0000"

	d2 := map[string]string{}
	d2["Red"] = "#FF0000"
	d2["Green"] = "#00FF00"
	d2["Blue"] = "#00FF"

	for k, v := range d2 {
		fmt.Printf("k:%s v:%s\n", k, v)
	}
	delete(d2, "Blue")
	delete(d2, "Blue1")
	fmt.Println("delete...")
	for k, v := range d2 {
		fmt.Printf("k:%s v:%s\n", k, v)
	}

}

type User struct {
	Name string
	Age  int
}

func TestType(t *testing.T) {
	//type User struct {
	//	Name string
	//	Age  int
	//}

	u1 := User{
		Name: "A",
		Age:  20,
	}

	u2 := User{
		Name: "B",
		Age:  30,
	}
	assert.Equal(t, u1.Age+10, u2.Age)

}

func (u *User) notify() {
	fmt.Printf("User: %s %d\n", u.Name, u.Age)
}

type NestUser struct {
	User
	U User
}

//func (u *NestUser) notify() {
//	fmt.Printf("NestUser: %s %d\n", u.Name, u.Age)
//}

func TestEmbedType(t *testing.T) {
	u := NestUser{User{Name: "A", Age: 1}, User{Name: "A", Age: 1}}

	// ==
	u.User.notify()
	u.notify()

	u.U.notify()
}

func TestGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("A " + strconv.Itoa(i))
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 200; i++ {
			fmt.Println("B " + strconv.Itoa(i))
		}
	}()
	wg.Wait()
}

type LogicProvider1 struct {
}

func (lp *LogicProvider1) Process(data string) {
	fmt.Println("LogicProvider1: " + data)
}

type LogicProvider2 struct {
}

func (lp *LogicProvider2) Process(data string) {
	fmt.Println("LogicProvider2: " + data)
}

type Logic interface {
	Process(data string)
}

type Logic1 interface {
	Process(data string)
}

type Client struct {
	Logic
}

type MixClient struct {
	Logic
	Logic1
}

func TestInterface(t *testing.T) {
	l := LogicProvider1{}
	l.Process("test")
	c := Client{&LogicProvider1{}}
	c.Process("test")

	l1 := LogicProvider2{}
	l1.Process("test")
	c1 := Client{&LogicProvider2{}}
	c1.Process("test")

	//m := MixClient{&LogicProvider1{}, &LogicProvider2{}}
	//// ./app_test.go:181:4: ambiguous selector m.Process
	//m.Process("test")

}

func TestEmptyInterface(t *testing.T) {
	var i interface{}
	i = 20
	assert.Equal(t, i, 20)
}

func TestPanicAndRecover(t *testing.T) {
	div60 := func(i int) {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println(v)
			}
		}()
		fmt.Println(60 / i)
	}

	for _, v := range []int{1, 2, 0, 6} {
		div60(v)
	}
}
