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

func TestType(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

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

type User struct {
	Name string
	Age  int
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
