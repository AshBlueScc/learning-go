package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)



func main() {
	session, err := mgo.Dial("")

	fmt.Println(session, err)
}
