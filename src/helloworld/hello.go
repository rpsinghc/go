package main

import "fmt"
import "encoding/base64"

func main() {
	data := []byte("any + old & data 444")
	str := base64.StdEncoding.EncodeToString(data)
	fmt.Println(str)	
}
