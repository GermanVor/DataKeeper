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
	Id       string   // ID of the Data
	DataType DataType // Type of the Data
	Data     []byte   // Serialized data
	Meta     []byte   // Meta information
}

type NewData struct {
	UserID   string
	DataType DataType // Type of the Data
	Data     []byte   // Serialized data
	Meta     []byte   // Meta information
}

type GetData struct {
	UserID string
	Id     string
}

type SetData struct {
	UserID string
	Id     string
	Data   []byte // new serialized data
	Meta   []byte // new meta information
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
	// Creates a new entry. Returns its `ID`
	New(ctx context.Context, newData *NewData) (string, error)

	// Requesting an entry
	Get(ctx context.Context, getData *GetData) (*Data, error)

	// Modifies an existing entry
	Set(ctx context.Context, setData *SetData) (*Data, error)

	// Returns a chunk of records of size `GetBatch.Limit` or less with offset `GetBatch.Offset`
	GetBatch(ctx context.Context, getBatch *GetBatch) ([]*Data, error)

	// Deletes an existing entry. Returns `true` if the entry is deleted
	Delete(ctx context.Context, deleteData *DeleteData) (bool, error)
}
