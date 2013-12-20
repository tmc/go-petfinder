package petfinder

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

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

// Breeds fetches the available breeds for an animal
func (a *api) Breeds(animalName string) ([]Breed, error) {
	r, err := a.getResponse("breed.list", map[string]string{
		"animal": animalName,
	})
	if err != nil {
		return []Breed{}, err
	}
	// TODO: check for nils along this path
	return r.breeds(), nil

}

// RandomPet fetches a list of random Pets
func (a *api) RandomPets() ([]Pet, error) {
	r, err := a.getResponse("pet.getRandom", map[string]string{
		"output": "full",
	})
	if err != nil {
		return []Pet{}, err
	}
	// TODO: check for nils along this path
	return r.pets(), nil

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

type apiResponse struct {
	Petfinder struct {
		Breeds apiResponseBreeds `json:"breeds"`
		Pets   []apiResponsePet  `json:"pet"`
		Header struct {
			Status struct {
				Code struct {
					T string `json:"$t"`
				} `json:"code"`
				Message struct{} `json:"message"`
			} `json:"status"`
			Timestamp struct {
				T string `json:"$t"`
			} `json:"timestamp"`
			Version struct {
				T string `json:"$t"`
			} `json:"version"`
		} `json:"header"`
	} `json:"petfinder"`
}

type apiResponseBreeds struct {
	Animal string `json:"@animal"`
	Breeds []struct {
		T string `json:"$t"`
	} `json:"breed"`
}

func (r apiResponse) breeds() []Breed {
	result := []Breed{}
	for _, breed := range r.Petfinder.Breeds.Breeds {
		result = append(result, Breed(breed.T))
	}
	return result
}

type apiResponsePet struct {
	Age struct {
		T string `json:"$t"`
	} `json:"age"`
	Animal interface{} `json:"animal"`
	Breeds interface{} `json:"breeds"`
	//Breeds struct {
	//	Breed []struct {
	//		T string `json:"$t"`
	//	} `json:"breed"`
	//} `json:"breeds"`
	Contact struct {
		Address1 interface{} `json:"address1"`
		Address2 interface{} `json:"address2"`
		City     struct {
			T string `json:"$t"`
		} `json:"city"`
		Email interface{} `json:"email"`
		Fax   interface{} `json:"fax"`
		Phone interface{} `json:"phone"`
		State struct {
			T string `json:"$t"`
		} `json:"state"`
		Zip struct {
			T string `json:"$t"`
		} `json:"zip"`
	} `json:"contact"`
	ID struct {
		T string `json:"$t"`
	} `json:"id"`
	Name struct {
		T string `json:"$t"`
	} `json:"name"`
	Description struct {
		T string `json:"$t"`
	} `json:"description"`
	LastUpdate struct{} `json:"lastUpdate"`
	Media      struct {
		Photos struct {
			Photo []struct {
				T     string `json:"$t"`
				_Id   string `json:"@id"`
				_Size string `json:"@size"`
			} `json:"photo"`
		} `json:"photos"`
	} `json:"media"`
	Mix struct {
		T string `json:"$t"`
	} `json:"mix"`
	Options struct{} `json:"options"`
	Sex     struct {
		T string `json:"$t"`
	} `json:"sex"`
	ShelterId struct {
		T string `json:"$t"`
	} `json:"shelterId"`
	ShelterPetId struct{} `json:"shelterPetId"`
	Size         struct {
		T string `json:"$t"`
	} `json:"size"`
	Status struct {
		T string `json:"$t"`
	} `json:"status"`
}

func (p apiResponsePet) pet() Pet {

	// TODO: flesh out
	var (
		id int
	)
	id, _ = strconv.Atoi(p.ID.T)
	pet := Pet{
		ID:          id,
		Age:         p.Age.T,
		Description: p.Description.T,
		Name:        p.Name.T,
		Status:      p.Status.T,
	}
	return pet
}

func (r apiResponse) pets() []Pet {
	result := []Pet{}
	for _, p := range r.Petfinder.Pets {
		result = append(result, p.pet())
	}
	return result
}
