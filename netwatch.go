package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type Service struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {

	var services []Service

	client := &http.Client{
		Timeout: time.Second * 3,
	}

	wg := new(sync.WaitGroup)

	GetJSONData("services.json", &services)

	for _, service := range services {
		wg.Add(1)

		go func(service Service, wg *sync.WaitGroup) {
			defer wg.Done()
			statusCode, responseTime := getStatusCode(client, service.Url)
			fmt.Println("Result for:", service.Name)
			switch {
			case statusCode == 200:
				fmt.Println(statusCode, "- Service is up and healthy! Response time:", responseTime.Milliseconds(), "ms")
			case statusCode == 401 || statusCode == 403:
				fmt.Println(statusCode, "- Client isn't authorized. Response time:", responseTime.Milliseconds(), "ms")
			case statusCode == 404:
				fmt.Println(statusCode, "- Resource not found. Response time:", responseTime.Milliseconds(), "ms")
			case statusCode >= 500 && statusCode <= 599:
				fmt.Println(statusCode, "- Server-side error.  Response time:", responseTime.Milliseconds(), "ms")
			case statusCode == 0:
				fmt.Println("No HTTP response")
			default:
				fmt.Println(statusCode, "- Not implemented")
			}
		}(service, wg)
	}
	wg.Wait()
}

func getStatusCode(client *http.Client, url string) (int, time.Duration) {

	start := time.Now()

	resp, err := client.Get(url)

	elapsed := time.Since(start)

	if err != nil {
		return 0, elapsed // In the future, adapt to provide more information on why
	}

	defer resp.Body.Close() // Close connection to save sys resources

	return resp.StatusCode, elapsed
}

func GetJSONData(filepath string, s *[]Service) {

	servicesJSON, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("This error occured while trying to read the json file:", err)
		return
	}

	// json.Unmarshal returns an error, data is saved not my assignment but by passing a pointer!
	if err := json.Unmarshal(servicesJSON, &s); err != nil {
		fmt.Println("This error occured while translating json to go objects:", err)
		return
	}
}
