package main

import (
	"fmt"
	"io"
	"net/http"
)

// const url = "https://lco.dev"
const url = "https://web.learncodeonline.in/"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	checkNilErr(err)

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // Caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)

	checkNilErr(err)

	content := string(databytes)
	fmt.Println(content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
