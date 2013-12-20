package petfinder

import (
	"encoding/json"
	"net/http"
	"net/url"
	"github.com/golang/glog"
)

type api struct {
	BaseURL     string
	Key, Secret string
}

// NewAPI creates a new api object from an API key and secret key
// TODO: secret is ignored for now
func NewAPI(key, secret string) (*api, error) {
	if key == "" {
		return nil, ErrMissingAPIKey
	}
	return &api{"http://api.petfinder.com/",
		key, secret}, nil
}

func (a *api) getResponse(endpoint string, args map[string]string) (*apiResponse, error) {
	values, _ := url.ParseQuery("")
	for key, value := range args {
		values.Add(key, value)
	}
	// add universal params
	values.Add("key", a.Key)
	values.Add("format", "json")

	u, err := url.Parse(a.BaseURL + endpoint + "?" + values.Encode())
	if err != nil {
		return nil, err
	}

	r, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	glog.Infoln("petfinder getResponse:", u.String(), r.StatusCode)
	// TODO if a totally bad response we get XML back, handle it if so

	resp := new(apiResponse)
	err = json.NewDecoder(r.Body).Decode(resp)
	if err != nil {
		return nil, err
	}
	// TODO: handle standard status fields in response body ala
	// err = resp.getError()

	return resp, nil
}

// Breeds fetches the available breeds for an animal
func (a *api) Breeds(animalName string) ([]Breed, error) {
	r, err := a.getResponse("breed.list", map[string]string{
		"animal": animalName,
	})
	if err != nil {
		return []Breed{}, err
	}
	// TODO: check for nils along this path
	return r.Petfinder.Breeds.breeds(), nil

}

type apiResponseBreeds struct {
	Animal string `json:"@animal"`
	Breed  []struct {
		T string `json:"$t"`
	} `json:"breed"`
}

func (b apiResponseBreeds) breeds() []Breed {
	result := []Breed{}
	for _, breed := range b.Breed {
		result = append(result, Breed(breed.T))
	}
	return result
}

type apiResponse struct {
	Petfinder struct {
		Breeds apiResponseBreeds `json:"breeds"`
		Header struct {
			Status struct {
				Code struct {
					T string `json:"$t"`
				} `json:"code"`
				Message struct{} `json:"message"`
			} `json:"status"`
			Timestamp struct {
				_T string `json:"$t"`
			} `json:"timestamp"`
			Version struct {
				_T string `json:"$t"`
			} `json:"version"`
		} `json:"header"`
	} `json:"petfinder"`
}
