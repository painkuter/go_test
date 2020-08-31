package interview

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	// Что выведет?
	names := []string{"john", "dave", "mike"}

	for i := range names {
		go func() {
			println(names[i])
		}()
	}
	//	time.Sleep(time.Second)
}

func Test2(t *testing.T) {
	// Что выведет?
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}

func Test3(t *testing.T) {
	// Что выведет?
	s := "mañana"
	println(len(s))
}

// ====================

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		println("from test()")
	}
	return nil
}

func Test4(t *testing.T) {
	// Что выведет?
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

// =============
func add(arr []int, v int) {
	arr[0] = 100
	arr = append(arr, v)
}

func Test5(t *testing.T) {
	// Что выведет?
	arr := make([]int, 2)
	fmt.Printf("%v %p\n", arr, &arr)
	add(arr, 10)
	fmt.Printf("%v %p\n", arr, &arr)
}

// =================

func Test6(t *testing.T) {
	i := 1
	defer println(func() int { return i * i }())
	i++
}

func Test7(t *testing.T) {
	// Что, в какой последовательности и ПОЧЕМУ выведет данный тест?
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(1))

	go func() {
		for {
			fmt.Println("hello")
		}
	}()

	time.Sleep(10 * time.Millisecond)

	go func() {
		for {
			fmt.Println("world")
		}
	}()

	time.Sleep(10 * time.Millisecond)
}

// ==========

/*type User struct{}

func NewUser() *User {
	user = &User{}

	return user
}

var user = NewUser()

func Test8(t *testing.T) {
	fmt.Println(user)
}
*/

func BenchmarkSumForward(b *testing.B) {
	nums := []int{}
	for i := 0; i < 5; i++ {
		nums = append(nums, i)
	}
	for n := 0; n < b.N; n++ {
		sum := nums[0] + nums[1] + nums[2] + nums[3] + nums[4]
		_ = sum
	}
}
func BenchmarkSumBackward(b *testing.B) {
	nums := []int{}
	for i := 0; i < 5; i++ {
		nums = append(nums, i)
	}
	for n := 0; n < b.N; n++ {
		sum := nums[4] + nums[3] + nums[2] + nums[1] + nums[0]
		_ = sum
	}
}

//
