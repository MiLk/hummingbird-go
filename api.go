package hummingbird

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

type API struct {
	endpoint string
	token    string
	request  *gorequest.SuperAgent
}

func (api *API) UserAuthenticate(username, email, password string) (errs []error, body string) {

	type UserAuthenticateData struct {
		Username string `json:"username,omitempty"`
		Email    string `json:"email,omitempty"`
		Password string `json:"password"`
	}

	data := UserAuthenticateData{
		Username: username,
		Email:    email,
		Password: password,
	}
	_, body, errs = api.request.
		Post(api.endpoint + "/v1/users/authenticate").
		Send(data).
		End()
	if len(errs) == 0 {
		api.token = body
	}
	return
}

func (api *API) UserInformation(username string) (errs []error, user User) {
	_, body, errs := api.request.
		Get(api.endpoint + "/v1/users/" + username).
		End()
	if len(errs) != 0 {
		return
	}

	err := json.Unmarshal([]byte(body), &user)
	if err != nil {
		errs = append(errs, err)
	}
	return
}

func NewAPI() *API {
	api := new(API)
	api.endpoint = "https://hummingbird.me/api"
	api.request = gorequest.New()
	return api
}
