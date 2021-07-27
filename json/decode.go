//lets discuss unmarshaling
package main

import (
	"encoding/json"
	"fmt"
)

type request struct {
	Name string
	Age  int
}

func main() {

	//lets discuss unmarshaling
	var req request

	str := `{"name":"Atul","age":35}`

	byteCode := []byte(str)
	json.Unmarshal(byteCode, &req)

	fmt.Println(req.Name)
	fmt.Println(req.Age)

}
