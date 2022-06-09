// Golang program to illustrate how to create
// an empty file in the default directory
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	var i = 1
	for i < 40 {
		str := fmt.Sprintf("日志%d.doc", i)
		myfile, e := os.Create(str)
		if e != nil {
			log.Fatal(e)
		}
		log.Println(myfile)
		myfile.Close()
		i += 1
	}
}
