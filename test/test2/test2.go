package main

import "crypto/sha1"
import "fmt"

func main() {
	s := `<request type="TRAINRETRIEVE"><retrieveKey>14192</retrieveKey></request>55fd79a7-652a-48be-94ff-1`
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
