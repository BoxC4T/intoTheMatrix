package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Example
func main() {

	resp, _ := http.Get("https://api.mcsrvstat.us/3/boxc4t.net")

	// convert resp to json
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result["ip"])

	defer resp.Body.Close()

}
