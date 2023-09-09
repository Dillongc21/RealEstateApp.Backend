package mockdata

import "github.com/Dillongc21/RealEstateApp.Backend/model"

func GetMockAddresses() []model.Address {
	return []model.Address {
		{ID: "0", Line1: "123 Fake Ln", StateID: "22", Zip: "12345"},
		{ID: "1", Line1: "54 Lemon St", Line2: "Apt. 102", StateID: "22", Zip: "54321"},
	}
}
