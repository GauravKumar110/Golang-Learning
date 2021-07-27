//lets marshal it
package main

import (
	"encoding/json"
	"fmt"
)

type response struct {
	Status  bool
	Message string
}

func main() {

	res := response{true, "All is well"}
	fmt.Println("lets Print Object", res)

	//lets marshal it

	data, err := json.Marshal(res)

	if err != nil {
		fmt.Println("Eror while Marshaling")
	}

	fmt.Println(res)
	fmt.Println(string(data))

}
