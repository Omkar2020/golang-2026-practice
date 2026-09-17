package main

import (
	"fmt"
)

var config map[string]string

func init() {
	config = map[string]string{"env": "production",
		"version": "1.0.0", "region": "us-east-1"}

	fmt.Println("Configuration initialized:", config)
}

func main() {
	fmt.Println("Main function started")
	fmt.Println("Running In:", config["env"])
}
