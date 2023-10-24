package mock

import (
	"errors"

	"github.com/Dillongc21/RealEstateApp.Backend/constants"
	"github.com/Dillongc21/RealEstateApp.Backend/model"
)

var addresses = []model.Address {
    {
        ID: "0", Line1: "123 Fake Ln", City: "Edgewood", State: "Washington",
        Country: "U.S", Zip: "12345",
    },
	{
        ID: "1", Line1: "54 Lemon St", Line2: "Apt. 102", City: "Bethel",
        State: "Washington", Country: "U.S.", Zip: "54321",
    },
}

func getAddressIndex(id string) int {
	for index, address := range addresses {
		if address.ID == id {
			return index
		}
	}
	return -1
}

func deleteAddressAtSliceIndex(slice []model.Address, index int) []model.Address {
    newSlice := append(slice[:index], slice[index + 1:]...)
    return newSlice
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

func UpdateAddressLine1(id string, newLine1 string) (model.Address, error) {
	index := getAddressIndex(id)

	if index == -1 {
		return model.Address{}, errors.New(constants.ERROR_ADDRESS_NOT_FOUND)
	}
    
	addresses[index].Line1 = newLine1
	return addresses[index], nil
}

func DeleteAddress(id string) (bool, error) {
    index := getAddressIndex(id)

    if index == -1 {
        return false, errors.New(constants.ERROR_ADDRESS_NOT_FOUND)
    }

    addresses = deleteAddressAtSliceIndex(addresses, index)

    index = getAddressIndex(id)
    if index != -1 {
        return false, errors.New(constants.ERROR_DELETE_ADDRESS_FAILED)
    }
    return true, nil
}
    
