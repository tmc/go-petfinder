package petfinder

import

// Animal is a petfinder Animal
"fmt"

type Animal string

// Breed is a petfinder Breed
type Breed string

type Pet struct {
	ID      int
	Age     string
	Animal  Animal
	Breeds  []Breed
	Contact struct {
		Email                                string
		Address1, Address2, City, State, Zip string
		Fax, Phone                           string
	}
	Description string
	LastUpdate  string // TODO: make time.Time
	Media       struct {
		Photos []struct {
			ID   string
			T    string
			Size string
		}
	}
	Mix          string
	Name         string
	Options      string
	Sex          string
	ShelterId    int
	ShelterPetId int
	Size         string
	Status       string
}

func (p Pet) String() string {
	return fmt.Sprintf("%s (id:%d)", p.Name, p.ID)
}
