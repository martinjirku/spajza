//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type StorageItemsQuantity string

const (
	StorageItemsQuantity_Mass        StorageItemsQuantity = "mass"
	StorageItemsQuantity_Length      StorageItemsQuantity = "length"
	StorageItemsQuantity_Volume      StorageItemsQuantity = "volume"
	StorageItemsQuantity_Temperature StorageItemsQuantity = "temperature"
	StorageItemsQuantity_Time        StorageItemsQuantity = "time"
	StorageItemsQuantity_Count       StorageItemsQuantity = "count"
)

func (e *StorageItemsQuantity) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "mass":
		*e = StorageItemsQuantity_Mass
	case "length":
		*e = StorageItemsQuantity_Length
	case "volume":
		*e = StorageItemsQuantity_Volume
	case "temperature":
		*e = StorageItemsQuantity_Temperature
	case "time":
		*e = StorageItemsQuantity_Time
	case "count":
		*e = StorageItemsQuantity_Count
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for StorageItemsQuantity enum")
	}

	return nil
}

func (e StorageItemsQuantity) String() string {
	return string(e)
}