package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, 14)

	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

func CheckPassword(hpass string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hpass), []byte(pass))
	return err == nil
}

func GenerateUserID(uid string) string {
	return uid
}

func GetApiID(UserName string, Password string, Email string) string {
	
	fmt.Println(UserName, Password)
	// ----------------- Fetch Access Token  ----------------
	requestBody := map[string]string{
		"username": "interview",
		"password": "Testpass123@",
	}
	body, err := Curl("https://api.proxybatch.com/api/v1/login", &requestBody, "")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	type FirstApi struct {
		Refresh string
		Token string
	}
	var firstApi FirstApi
	err = json.Unmarshal(body, &firstApi)
	if err != nil { return "" }
	fmt.Println(firstApi.Token)

	//----------- create sub user --------------------
	requestBody = map[string]string{
		"email": Email,
		"username": UserName,
		"password": Password,
	}
	body, err = Curl("https://api.proxybatch.com/api/v1/users", &requestBody, firstApi.Token)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	type SecondApi struct {
		Id string
	}
	var secondApi SecondApi
	err = json.Unmarshal(body, &secondApi)
	if err != nil { return "" }
	fmt.Println(secondApi.Id)
	return secondApi.Id
}

func Curl(url string, reqData *map[string]string, token string) ([]byte, error) {
	// Define the API endpoint

	// Create the request body as a map
	// requestBody := map[string]string{
	// 	"username": UserName,
	// 	"password": Password,
	// }
	requestBody := reqData

	// Convert the request body to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return nil, err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer " + token)
	}
	// Make the HTTP request using the http.Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}
	fmt.Println(string(body))
	return body, nil
}