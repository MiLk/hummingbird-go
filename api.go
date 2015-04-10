package hummingbird

import (
	"encoding/json"
	"net/url"

	"github.com/parnurzeal/gorequest"
)

type API struct {
	endpoint string
	token    string
	request  *gorequest.SuperAgent
}

func (api *API) Search(title string) (errs []error, results []Anime) {
	_url := api.endpoint + "/v1/search/anime?query=" + url.QueryEscape(title)
	_, body, errs := api.request.Get(_url).End()
	if len(errs) != 0 {
		return
	}

	err := json.Unmarshal([]byte(body), &results)
	if err != nil {
		errs = append(errs, err)
	}
	return
}

func (api *API) Library(username, status string) (errs []error, library []LibraryEntry) {
	_url := api.endpoint + "/v1/users/" + url.QueryEscape(username) + "/library"
	if len(status) > 0 {
		_url += "?status=" + url.QueryEscape(status)
	}
	_, body, errs := api.request.Get(_url).End()
	if len(errs) != 0 {
		return
	}

	err := json.Unmarshal([]byte(body), &library)
	if err != nil {
		errs = append(errs, err)
	}
	return
}

func (api *API) UserAuthenticate(username, email, password string) (errs []error, body string) {
	data := map[string]string{
		"username": username,
		"email":    email,
		"password": password,
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
		Get(api.endpoint + "/v1/users/" + url.QueryEscape(username)).
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

func (api *API) UserFavorites(username string) (errs []error, animes []Anime) {
	_, body, errs := api.request.
		Get(api.endpoint + "/v1/users/" + url.QueryEscape(username) + "/favorite_anime").
		End()
	if len(errs) != 0 {
		return
	}

	err := json.Unmarshal([]byte(body), &animes)
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
