package main

// package apiTokenControl

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

var Token sToken

var InitialParams sInitialParams

func initialParams(params sInitialParams) (sInitialParams, error) {

	if (params == sInitialParams{}) {
		err := errors.New("initial params not informed")
		return sInitialParams{}, err
	}

	return params, nil
}

func main() {

	InitialParams = sInitialParams{
		Grant_type:    "Grant",
		Client_id:     "Client",
		Client_secret: "secret",
		Code:          "code",
		Redirect_uri:  "uri",
	}
	new_token()
}
func new_token() (sToken, error) {

	params, err := initialParams(InitialParams)
	if err != nil {
		return sToken{}, err
	}

	body, _ := json.Marshal(map[string]string{
		"grant_type":    params.Grant_type,
		"client_id":     params.Client_id,
		"client_secret": params.Client_secret,
		"code":          params.Code,
		"redirect_uri":  params.Redirect_uri,
	})

	token, err_response := response_token(body)

	if err_response != nil {
		err := errors.New("error on get new Token, please provide a new CODE")
		return token, err
	}
	return token, nil
}

func response_token(body []byte) (sToken, error) {

	token := sToken{}
	errorToken := Errors{}

	var url string = "https://api.mercadolibre.com/oauth/token"

	payload := bytes.NewBuffer(body)

	resp, err := http.Post(url, "application/json", payload)
	if err != nil {
		return token, err
	}
	defer resp.Body.Close()

	_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}

	if err := json.Unmarshal(_body, &errorToken); err != nil {
		err := errors.New("error: " + errorToken.Error + "|" + strconv.Itoa(errorToken.Status) + "|" + errorToken.Message)
		return token, err
	}

	if err := json.Unmarshal(_body, &token); err != nil {
		return token, err
	}

	if (sToken{} == token) {
		err := errors.New("Error on call Token |" + string(body))
		return token, err

	} else {
		return token, nil
	}

}
