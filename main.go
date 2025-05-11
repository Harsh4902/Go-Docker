package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// cmd := os.Args[1]

	// if cmd == "start" {
	// 	imageName := os.Args[2]
	// 	driver := os.Args[3]
	// 	cfg, err := LoadConfig(ConfigPath)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	cfg.Instance.Driver = driver
	// 	SaveConfig(ConfigPath, cfg)
	// 	StartDocker(imageName)
	// }

	// if cmd == "stop" {
	// 	containerId := os.Args[2]
	// 	StopDocker(containerId)
	// }

	endpoint := "http://localhost:18080/realms/microcks/protocol/openid-connect/token"

	// Prepare form data
	data := url.Values{}
	data.Set("client_id", "microcks-serviceaccount")
	data.Set("client_secret", "ab54d329-e435-41ae-a900-ec6b3fe15c54")
	data.Set("username", "admin")
	data.Set("password", "microcks123")
	data.Set("grant_type", "password")

	// Create HTTP request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}
