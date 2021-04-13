package main

import (
	"fmt"
	"strconv"
)

type Key struct {
	Row, Col int
}

type TableSample struct {
	Id   int
	Name string
	Age  int
}

func main() {

	// テーブル構造のデータを持ちたい場合は、構造体でキーを作ると便利感
	records := map[Key]string{}
	var count int = 10

	for i := 0; i < count; i++ {

		records[Key{i, 0}] = "Hello 0:" + strconv.Itoa(i)
		records[Key{i, 1}] = "Hello 1:" + strconv.Itoa(i)
	}

	fmt.Println(records)
	fmt.Println(records[Key{0, 0}])

	// 構造体でMapを使う場合は、ポインタにすると良い
	records2 := map[int]*TableSample{}

	for i := 0; i < count; i++ {

		records2[i] = &TableSample{
			Id:   i,
			Name: "Struct Test : " + strconv.Itoa(i),
			Age:  i,
		}
	}

	fmt.Println(records2)
}
