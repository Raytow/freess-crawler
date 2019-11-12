package main

import (
	"container/list"
	"fmt"
)

var db = list.New()

func store(a interface{}) {
	db.PushBack(a)
}

func readAll() {
	fmt.Println(db.Len())
}
