// Package hummingbird is an API client for the https://hummingbird.me website.
package hummingbird

import (
	"encoding/json"
	"net/url"

	"github.com/parnurzeal/gorequest"
)

// API represents the API client
type API struct {
	endpoint string
	token    string
	request  *gorequest.SuperAgent
}

// Search allows you to search an anime by title.
//
// https://github.com/hummingbird-me/hummingbird/wiki/API-v1-Methods#anime--search-by-title
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

// Library allows you to retrieve all the anime in the library of an user of the given type.
//
// https://github.com/hummingbird-me/hummingbird/wiki/API-v1-Methods#library--get-all-entries
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

// UserAuthenticate returns an user's authentication token if the credentials are correct.
//
// https://github.com/hummingbird-me/hummingbird/wiki/API-v1-Methods#user--authenticate
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

// UserInformation allows you to retrieve all the informations about an user.
//
// https://github.com/hummingbird-me/hummingbird/wiki/API-v1-Methods#user--get-activity-feed
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

// UserFavorites allows you to retrieve the favorite animes of a given user.
//
// https://github.com/hummingbird-me/hummingbird/wiki/API-v1-Methods#user--get-favorite-anime
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

// Instantiates the API and return a new API instance.
func NewAPI() *API {
	api := new(API)
	api.endpoint = "https://hummingbird.me/api"
	api.request = gorequest.New()
	return api
}
