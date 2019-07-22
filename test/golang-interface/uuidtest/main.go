package main

import (
	"fmt"
	"github.com/oklog/ulid"
	"golang-interface/hashset"
	"math/rand"
	"time"
)

func getUUID() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}

func main() {
	set := hashset.New()
	for i := 1; i <= 10000000; i++ {
		set.Add(getUUID())
	}
	fmt.Println(set.Size() == 10000000)
	//fmt.Println("----")
	//fmt.Println(set.String())
}
