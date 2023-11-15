package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	result, err := http.Get("https://www.sofascore.com/team/football/barcelona/2817")

	if err != nil {
		return
	}

	body, err := io.ReadAll(result.Body)

	fmt.Println(string(body))
}
