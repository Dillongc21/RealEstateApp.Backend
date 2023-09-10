package mockdata

import (
	"errors"
	"strconv"

	"github.com/Dillongc21/RealEstateApp.Backend/model"
)

var addresses = []model.Address {
	{ID: "0", Line1: "123 Fake Ln", StateID: "22", Zip: "12345"},
	{ID: "1", Line1: "54 Lemon St", Line2: "Apt. 102", StateID: "22", Zip: "54321"},
}

func GetMockAddresses() []model.Address { return addresses }

func GetMockAddressByID(id string) model.Address {
	for _, address := range addresses {
		if address.ID == id {
			return address
		}
	}
	return model.Address{}
}

func AppendMockAddress(newAddress model.Address) {
	addresses = append(addresses, newAddress)
}

func GetAddressIndex(id string) (int, error) {
	for _, address := range addresses {
		if address.ID == id {
			return strconv.Atoi(id)
		}
	}
	return -1, nil
}

func UpdateAddressLine1(id string, newLine1 string) (model.Address, error) {
	index, _ := GetAddressIndex(id)
	if index == -1 {
		return model.Address{}, errors.New("Address not Found")
	}

	addresses[index].Line1 = newLine1
	return addresses[index], nil
}
