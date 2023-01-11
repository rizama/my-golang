package main

/**
Package container/ring adalah implementasi struktur data circular list
Circular list adalah struktur data ring, dimana diakhir element akan kembali ke element awal (HEAD)
https://golang.org/pkg/container/ring/
*/

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
