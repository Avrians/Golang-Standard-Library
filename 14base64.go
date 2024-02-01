package main

import (
	"encoding/base64"
	"fmt"
)

func main14()  {
	value := "Kevin Aldo"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println("Encoded:", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Decoded:", string(decoded))
	}
}