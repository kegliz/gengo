package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kegliz/gengo/genstack"
	"github.com/kegliz/gengo/intstack"
	"github.com/kegliz/gengo/structstack"
)

var count int = 10000000

func main() {
	gus := genstack.NewStack[int]()
	is := intstack.NewStack()

	gss := genstack.NewStack[structstack.Item]()
	ss := structstack.NewStack()

	var start time.Time
	var elapsed time.Duration

	devNull, err := os.OpenFile("/dev/null", os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer devNull.Close()

	// Push int ========================================
	fmt.Printf("\nMeasuring %d int Push(s)\n", count)
	for i := 0; i < count; i++ {
		start = time.Now()
		gus.Push(i)
		elapsed += time.Since(start)
	}
	fmt.Printf("genstack Push time: %s\n", elapsed)

	elapsed = 0
	for i := 0; i < count; i++ {
		start = time.Now()
		is.Push(i)
		elapsed += time.Since(start)
	}
	fmt.Printf("intstack Push time: %s\n", elapsed)

	// Push struct ========================================
	fmt.Printf("\nMeasuring %d struct Push(s)\n", count)
	elapsed = 0
	for i := 0; i < count; i++ {
		start = time.Now()
		gss.Push(structstack.Item{Name: "Alice", Age: i})
		elapsed += time.Since(start)
	}
	fmt.Printf("genstack Push time: %s\n", elapsed)

	elapsed = 0
	for i := 0; i < count; i++ {
		start = time.Now()
		ss.Push(structstack.Item{Name: "Alice", Age: i})
		elapsed += time.Since(start)
	}
	fmt.Printf("structstack Push time: %s\n", elapsed)

	// Peek int ========================================
	fmt.Printf("\nMeasuring %d int Peek(s)\n", count)
	var num int
	elapsed = 0
	for i := 0; i < count; i++ {
		start = time.Now()
		num, _ = gus.Peek()
		elapsed += time.Since(start)
		touch(num, devNull)
	}
	fmt.Printf("genstack Peek time: %s\n", elapsed)

	elapsed = 0
	for i := 0; i < count; i++ {
		start = time.Now()
		num, _ = is.Peek()
		elapsed += time.Since(start)
		touch(num, devNull)
	}
	fmt.Printf("intstack Peek time: %s\n", elapsed)

	// Peek struct ========================================
	fmt.Printf("\nMeasuring %d struct Peek(s)\n", count)
	elapsed = 0
	for i := 0; i < count; i++ {
		start = time.Now()
		item, _ := gss.Peek()
		elapsed += time.Since(start)
		touch(item.Age, devNull)
	}
	fmt.Printf("genstack Peek struct time: %s\n", elapsed)

	elapsed = 0
	for i := 0; i < count; i++ {
		start = time.Now()
		item, _ := ss.Peek()
		elapsed += time.Since(start)
		touch(item.Age, devNull)
	}
	fmt.Printf("structstack Peek struct time: %s\n", elapsed)

	// Pop int ========================================
	fmt.Printf("\nMeasuring %d int Pop(s) while summing up\n", count)
	var sum int
	elapsed = 0
	for i := 0; i < count; i++ {
		start = time.Now()
		num, _ = gus.Pop()
		elapsed += time.Since(start)
		sum += num
	}
	fmt.Printf("genstack Pop time: %s\n", elapsed)
	touch(sum, devNull)

	elapsed, sum = 0, 0
	for i := 0; i < count; i++ {
		start = time.Now()
		num, _ = is.Pop()
		elapsed += time.Since(start)
		sum += num
	}
	fmt.Printf("intstack Pop time: %s\n", elapsed)
	touch(sum, devNull)

	// Pop struct ========================================
	fmt.Printf("\nMeasuring %d struct Pop(s) while summing up\n", count)
	elapsed, sum = 0, 0
	for i := 0; i < count; i++ {
		start = time.Now()
		item, _ := gss.Pop()
		elapsed += time.Since(start)
		sum += item.Age
	}
	fmt.Printf("genstack Pop struct time: %s\n", elapsed)
	touch(sum, devNull)
	elapsed, sum = 0, 0
	for i := 0; i < count; i++ {
		start = time.Now()
		item, _ := ss.Pop()
		elapsed += time.Since(start)
		sum += item.Age
	}
	fmt.Printf("structstack Pop struct time: %s\n", elapsed)
	touch(sum, devNull)
}

func touch(num int, devNull *os.File) {
	_, err := devNull.Write([]byte(string(num)))
	if err != nil {
		panic(err)
	}
}
