package clientapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"notification-api/excepriton"
)

// SendSignUpData -> send auth data to the main client API
func SendSignUpData(dto any) error {
	// http example ->
	// https://thedevelopercafe.com/articles/make-post-request-in-go-d9756284d70b

	fmt.Println("dto is ->\n", dto)

	client := http.Client{}
	requestURL := "str"  // config.ClientApiAuth()
	accessToken := "str" // config.GetClientAPIAccessToken()

	body, err := json.Marshal(&dto)
	if err != nil {
		excepriton.HandleAnError("unmarshal err:", err)
		return err
	}

	reader := bytes.NewReader(body)

	req, err := http.NewRequest(http.MethodPost, requestURL, reader)
	if err != nil {
		excepriton.HandleAnError("creation a request was failed: ", err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("AccessToken", accessToken)

	res, err := client.Do(req)
	if err != nil || res.StatusCode != 201 {
		excepriton.HandleAnError("error making http request: ", err)
		return err
	}

	defer res.Body.Close()
	fmt.Println("res status -> ", res.StatusCode)
	return nil
}
