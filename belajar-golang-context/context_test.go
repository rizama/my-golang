package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)

}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// Get Value Context
	// ketika membaca data dari Context, pertama yang akan dibaca adalah child paling bawah,
	// jika tidak ada mundur ke parent, jika tidak ada mundur lagi ke parent lebih tinggi dst sampai parent paling akhir.
	// parent1 -> parent2 -> parent3 -> child

	fmt.Println(contextF.Value("f")) // dapat data, karena pada contextF memang ada value dengan key "f"
	fmt.Println(contextF.Value("c")) // dapat data parent, karena contextF adalah child dari contextC dimana contextC memiliki key "c"
	fmt.Println(contextF.Value("b")) // tidak mendapatkan data, karena dari contextF sampai ke parent tertinggi tidak ada parent yang memiliki key "b" / disini kasusnya beda parent
	fmt.Println(contextA.Value("b")) // tidak mendapatkan data, karena tidak bisa mengambil data dari child

}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)

		counter := 0

		for {
			select {
			case <-ctx.Done():
				{
					return
				}
			default:
				{
					destination <- counter
					counter++
					time.Sleep(1 * time.Second) // simulasi slow proses
				}
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	// Memberi signal cancel pada context (ctx) yang dikirim ke CreateConter()
	cancel()

	time.Sleep(5 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

}
