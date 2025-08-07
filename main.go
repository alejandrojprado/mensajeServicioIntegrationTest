package main

import (
	"fmt"
	"mensajeServiceIntegrationTests/componets/client"
	"os"
)

var (
	APIClient  *client.Client
	ServiceURL string
)

func init() {
	ServiceURL = os.Getenv("SERVICE_URL")
	if ServiceURL == "" {
		panic("SERVICE_URL environment variable is required")
	}
	APIClient = client.NewClient(ServiceURL)
}

func main() {
	fmt.Printf("Running tests against: %s\n", ServiceURL)
}
