package authenticate

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type AuthenticationServiceInteractor struct {
	client *http.Client
}

func (a *AuthenticationServiceInteractor) GetTokenUser(ctx context.Context, userName string,
	password string) (string, error) {

	requestDataByte, errReq := json.Marshal(&RequestGetTokenUser{
		UserName: userName,
		Password: password,
	})
	if errReq != nil {
		return "", nil
	}

	req, _ := http.NewRequest("POST", os.Getenv("GET-TOKEN-USER"), bytes.NewReader(requestDataByte))
	req.WithContext(ctx).Header.Add("Content-Type", "application/json")
	req.WithContext(ctx).Header.Add("Accept", "application/json")

	client := a.client
	response, errRes := client.Do(req)
	if errRes != nil {
		return "", errRes
	}
	defer response.Body.Close()
	data, readErr := ioutil.ReadAll(response.Body)

	if readErr != nil {
		return "", readErr
	}

	responseToken := &ResponseGetTokenUser{}
	jsonErr := json.Unmarshal(data, &responseToken)
	if jsonErr != nil {
		return "", jsonErr
	}

	return responseToken.Token, nil
}

func NewAuthService(clientHttp *http.Client) AuthenthicationService {
	return &AuthenticationServiceInteractor{client: clientHttp}
}
