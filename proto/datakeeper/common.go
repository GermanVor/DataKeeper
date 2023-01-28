package proto

import (
	"errors"

	"github.com/GermanVor/data-keeper/cmd/storageServer/storage"
)

func (d *DataType) Format() storage.DataType {
	switch *d {
	case DataType_BANK_CARD:
		return storage.BankCard
	case DataType_LOGG_PASS:
		return storage.LogPass
	case DataType_OTHER:
		return storage.Other
	case DataType_TEXT:
		return storage.Text
	default:
		return storage.Unknown
	}
}

func (r *NewRequest) Format(userId string) (*storage.NewData, error) {
	dataType := r.DataType.Format()

	if dataType == storage.Unknown {
		return nil, errors.New("")
	}

	return &storage.NewData{
		UserId:   userId,
		DataType: dataType,
		Data:     r.Data,
		Meta:     r.Meta,
	}, nil
}

func (r *SetRequest) Format(userId string) *storage.SetData {
	return &storage.SetData{
		UserId: userId,
		Id:     r.Id,
		Data:   r.Data,
		Meta:   r.Meta,
	}
}
