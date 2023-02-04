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
	Meta     []byte
}

type NewData struct {
	UserID   string
	DataType DataType
	Data     []byte
	Meta     []byte
}

type GetData struct {
	UserID string
	Id     string
}

type SetData struct {
	UserID string
	Id     string
	Data   []byte
	Meta   []byte
}

type DeleteData struct {
	UserID string
	Id     string
}

type GetBatch struct {
	UserID string
	Offset int32
	Limit  int32
}

type Interface interface {
	New(ctx context.Context, newData *NewData) (string, error)
	Get(ctx context.Context, getData *GetData) (*Data, error)
	Set(ctx context.Context, setData *SetData) (*Data, error)
	GetBatch(ctx context.Context, getBatch *GetBatch) ([]*Data, error)
	Delete(ctx context.Context, deleteData *DeleteData) (bool, error)
}
