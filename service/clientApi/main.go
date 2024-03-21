package clientapi

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"notification_service/config"
// 	"os"
// )

// // SendSignUpData -> send auth data to the client API
// func SendSignUpData(dto any) bool {
// 	// http example ->
// 	// https://thedevelopercafe.com/articles/make-post-request-in-go-d9756284d70b

// 	fmt.Println("dto is ->\n", dto)

// 	client := http.Client{}
// 	requestURL := config.ClientApiAuth()
// 	accessToken := config.GetClientAPIAccessToken()

// 	body, err := json.Marshal(&dto)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		return false
// 	}

// 	reader := bytes.NewReader(body)

// 	req, err := http.NewRequest(http.MethodPost, requestURL, reader)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		return false
// 	}

// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("AccessToken", accessToken)

// 	res, err := client.Do(req)
// 	if err != nil || res.StatusCode != 201 {
// 		fmt.Printf("error making http request: %s\n", err)
// 		return false
// 	}

// 	defer res.Body.Close()
// 	fmt.Println("res status -> ", res.StatusCode)
// 	return true
// }
