package storage

import (
	"context"
)

type DataType int64

const (
	LogPass DataType = iota
	Text
	BankCard
	Other
	Unknown
)

type MetaType map[string]string

type Data struct {
	Id       string
	DataType DataType
	Data     []byte
	Meta     MetaType
}

type NewData struct {
	UserId   string
	DataType DataType
	Data     []byte
	Meta     MetaType
}

type GetData struct {
	UserId string
	Id     string
}

type SetData struct {
	UserId string
	Id     string
	Data   []byte
	Meta   MetaType
}

type DeleteData struct {
	UserId string
	Id     string
}

type Interface interface {
	New(ctx context.Context, newData *NewData) (string, error)
	Get(ctx context.Context, getData *GetData) (*Data, error)
	Set(ctx context.Context, setData *SetData) (*Data, error)
	Delete(ctx context.Context, deleteData *DeleteData) (bool, error)
}
