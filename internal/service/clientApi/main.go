package clientapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"notification-api/pkg/exceptions"
)

// SendSignUpData -> send auth data to the main client API
func SendSignUpData(dto any) error {

	client := &http.Client{}
	requestURL := "str"  // config.ClientApiAuth()
	accessToken := "str" // config.GetClientAPIAccessToken()

	body, err := json.Marshal(&dto)
	if err != nil {
		exceptions.HandleAnError("unmarshal err:" + err.Error())
		return err
	}

	reader := bytes.NewReader(body)

	req, err := http.NewRequest(http.MethodPost, requestURL, reader)
	if err != nil {
		exceptions.HandleAnError("creation a request was failed: " + err.Error())
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("AccessToken", accessToken)

	res, err := client.Do(req)
	if err != nil || res.StatusCode != 201 {
		exceptions.HandleAnError("error making http request: " + err.Error())
		return err
	}

	defer res.Body.Close()
	// fmt.Println("res status -> ", res.StatusCode)
	return nil
}
