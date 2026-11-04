package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// http
// 1. HTTP stands for hyper text transfer protocol
// 2. Provides set of rules defining how data is exchanged between web clients and services
// 3. It comes under application layer
// 4. It is a layer 7 protocol of the OSI network model

// net/http package
// 1. The standard lib net/http package provides all the functionality to work with HTTP
// 2. You can make HTTP requests with the help of http client provided by the package

type boringResponse struct {
	Participants  int     `json:"participants"`
	Price         float64 `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
	Accessibility float64 `json:"accessibility"`
}

func main() {
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://www.boredapi.com/api/activity", nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error doing request: ", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error invalid status code", resp.Status)
	}

	var boringResponse boringResponse
	if err := json.NewDecoder(resp.Body).Decode(&boringResponse); err != nil {
		fmt.Println("Error reading response body: ", err)
		return
	}

	fmt.Printf("%+v\n", boringResponse)

}
