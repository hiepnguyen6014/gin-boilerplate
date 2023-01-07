package utils

import (
	"encoding/json"
	"io"
	"mono-golang/src/constants"
	"mono-golang/src/types"
	"net/http"
	"net/url"
)

func GetTokenByCode(code string) (googleResponse types.GoogleResponse, err error) {
	values := url.Values{"grant_type": {"authorization_code"}}
	values.Add("client_id", constants.GOOGLE_CLIENT_ID)
	values.Add("redirect_uri", constants.GOOGLE_REDIRECT_URI)
	values.Add("client_secret", constants.GOOGLE_CLIENT_SECRET)
	values.Add("code", code)

	response, err := http.PostForm(constants.GOOGLE_TOKEN_URI, values)
	response.Header.Set("Accept", "application/xml")

	defer response.Body.Close()

	if err != nil {
		return googleResponse, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return googleResponse, err
	}

	json.Unmarshal(body, &googleResponse)

	return googleResponse, err
}
