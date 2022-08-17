package apiTokenControl

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

//	type SInitialParams struct {
//	    Grant_type    string `json:"grant_type"`
//	    Client_id     string `json:"client_id"`
//	    Client_secret string `json:"client_secret"`
//	    Code          string `json:"code"`
//	    Redirect_uri  string `json:"redirect_uri"`
//	}
var InitialParams SInitialParams

func New() (Token, error) {

	params, err := initialParams(InitialParams)
	if err != nil {
		return Token{}, err
	}

	body, _ := json.Marshal(map[string]string{
		"grant_type":    "authorization_code",
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

func MakeRefreshToken(s Token) (Token, error) {

	if strings.Trim(s.Refresh_token, " ") == "" {
		err := errors.New("impossible generate refresh token, trying get new token")
		return Token{}, err
	}

	params, err := initialParams(InitialParams)
	if err != nil {
		return Token{}, err
	}

	body, _ := json.Marshal(map[string]string{
		"grant_type":    "refresh_token",
		"client_id":     params.Client_id,
		"client_secret": params.Client_secret,
		"refresh_token": s.Refresh_token,
	})

	token, err_response := response_token(body)

	if err_response != nil {
		err := errors.New("error on make_refresh_token")
		return Token{}, err
	}

	return token, nil
}

func response_token(body []byte) (Token, error) {

	token := Token{}
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

	if (Token{} == token) {
		err := errors.New("Error on call Token |" + string(body))
		return token, err

	} else {
		return token, nil
	}

}

func initialParams(params SInitialParams) (SInitialParams, error) {

	if (params == SInitialParams{}) {
		err := errors.New("initial params not informed")
		return SInitialParams{}, err
	}

	return params, nil
}
