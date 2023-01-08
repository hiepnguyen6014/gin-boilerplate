package services

import (
	"encoding/json"
	"io"

	"mono-golang/src/constants"
	"mono-golang/src/models"
	"mono-golang/src/repos"
	"mono-golang/src/types"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

func RedirectGoogleLoginPage() string {
	values := url.Values{"response_type": {"code"}, "prompt": {"login"}}

	values.Add("redirect_uri", constants.GOOGLE_REDIRECT_URI)
	values.Add("client_id", constants.GOOGLE_CLIENT_ID)
	values.Add("scope", getScope())
	values.Add("state", uuid.NewString())

	return constants.GOOGLE_AUTH_URI + "?" + values.Encode()
}

func GoogleCallback(queries url.Values) (int, interface{}) {
	code := queries.Get("code")
	state := queries.Get("state")

	if state == "" {
		return http.StatusOK, map[string]string{
			"message": "session error",
		}
	}

	tokens, err := getTokensByCode(code)

	if err != nil {
		return http.StatusBadGateway, map[string]string{
			"message": err.Error(),
		}
	}

	return http.StatusOK, tokens
}

func SaveUserProfile(googleResponse types.GoogleResponse) {
	access_token, id_token := googleResponse.Access_token, googleResponse.Id_token
	user, _ := getUserProfile(access_token)
	user.Id_token = id_token

	repos.CreateUser(user)
}

// RedirectLoginPage
func getScope() (scope string) {
	scopes := []string{"profile", "email"}

	scope = strings.Join(scopes, " ")

	return scope
}

// GoogleCallback
func getTokensByCode(code string) (googleResponse types.GoogleResponse, err error) {

	values := url.Values{"grant_type": {"authorization_code"}}
	values.Add("client_id", constants.GOOGLE_CLIENT_ID)
	values.Add("redirect_uri", constants.GOOGLE_REDIRECT_URI)
	values.Add("client_secret", constants.GOOGLE_CLIENT_SECRET)
	values.Add("code", code)

	response, err := http.PostForm(constants.GOOGLE_TOKEN_URI, values)
	response.Header.Set("Accept", "application/json")

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

// SaveUserProfile
func getUserProfile(access_token string) (user models.User, err error) {
	url_values := url.Values{"access_token": {access_token}}.Encode()
	uri := constants.GOOGLE_AUTH_USER_INFO + "?" + url_values

	response, err := http.Get(uri)
	response.Header.Set("Accept", "application/json")

	defer response.Body.Close()

	if err != nil {
		return user, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return user, err
	}

	json.Unmarshal(body, &user)

	return user, err
}
