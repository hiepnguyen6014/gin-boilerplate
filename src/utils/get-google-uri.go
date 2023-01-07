package utils

import (
	"mono-golang/src/constants"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

func getScope() (scope string) {
	scopes := []string{"profile", "email"}

	scope = strings.Join(scopes, " ")

	return scope
}

func GetGoogleUri() string {
	values := url.Values{"response_type": {"code"}, "prompt": {"login"}}

	values.Add("redirect_uri", constants.GOOGLE_REDIRECT_URI)
	values.Add("client_id", constants.GOOGLE_CLIENT_ID)
	values.Add("scope", getScope())
	values.Add("state", uuid.NewString())

	return constants.GOOGLE_AUTH_URI + "?" + values.Encode()
}
