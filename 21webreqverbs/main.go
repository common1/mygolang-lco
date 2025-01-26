package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	PeformGetRequest()
}

func PeformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	checkNilErr(err)

	defer response.Body.Close()

	fmt.Println("Status code: ", response.Status)
	fmt.Println("Content length is: ", response.ContentLength)

	// fmt.Println(content)
	// fmt.Println(string(content))

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Bytecount is: ", byteCount)
	fmt.Println(responseString.String())
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
