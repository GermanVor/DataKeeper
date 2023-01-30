package storage

import (
	"context"
	"log"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
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

var (
	// INSERT INTO main_storage (userID, data, dataType, meta)
	// VALUES (($1, $2, $3, $4) RETURNING main_storage.id;
	insertNewDataSQL = "INSERT INTO main_storage (userID, data, dataType, meta) " +
		"VALUES ($1, $2, $3, $4) RETURNING main_storage.id"

	// SELECT data, dataType, meta FROM main_storage WHERE userID=$1 AND id=$2;
	getDataSQL = "SELECT data, dataType, meta FROM main_storage WHERE userID=$1 AND id=$2"

	// UPDATE main_storage SET (data, meta) = ($3, $4) WHERE userID=$1 AND id=$2 RETURNING main_storage.dataType;
	setDataSQL = "UPDATE main_storage SET (data, meta) = " +
		"($3, $4) WHERE userID=$1 AND id=$2 RETURNING main_storage.dataType"

	// DELETE FROM main_storage WHERE userID=$1 AND id=$2;
	deleteSQL = "DELETE FROM main_storage WHERE userID=$1 AND id=$2"

	// SELECT id, data, dataType, meta FROM main_storage WHERE userID=$1 LIMIT $2 OFFSET $3;
	getBatchSQL = "SELECT id, data, dataType, meta FROM main_storage WHERE userID=$1 LIMIT $2 OFFSET $3"

	createTableSQL = "CREATE TABLE IF NOT EXISTS main_storage (" +
		"userID text, " +
		"id SERIAL, " +
		"data bytea, " +
		"dataType smallint," +
		"meta bytea " +
		");"
)

type Impl struct {
	dbPool *pgxpool.Pool
}

func (s *Impl) New(ctx context.Context, newData *NewData) (string, error) {
	id := 0
	err := s.dbPool.QueryRow(
		ctx,
		insertNewDataSQL,
		newData.UserID,
		newData.Data,
		newData.DataType,
		newData.Meta,
	).Scan(&id)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(id), nil
}

func (s *Impl) Get(ctx context.Context, getData *GetData) (*Data, error) {
	resp := &Data{
		Id: getData.Id,
	}

	err := s.dbPool.QueryRow(
		ctx,
		getDataSQL,
		getData.UserID,
		getData.Id,
	).Scan(&resp.Data, &resp.DataType, &resp.Meta)

	return resp, err
}

func (s *Impl) Set(ctx context.Context, setData *SetData) (*Data, error) {
	resp := &Data{
		Id:   setData.Id,
		Data: setData.Data,
		Meta: setData.Meta,
	}

	err := s.dbPool.QueryRow(
		ctx,
		setDataSQL,
		setData.UserID,
		setData.Id,
		setData.Data,
		setData.Meta,
	).Scan(&resp.DataType)

	return resp, err
}

func (s *Impl) Delete(ctx context.Context, deleteData *DeleteData) (bool, error) {
	tag, err := s.dbPool.Exec(ctx, deleteSQL, deleteData.UserID, deleteData.Id)
	return tag.RowsAffected() == 1, err
}

func (s *Impl) GetBatch(ctx context.Context, getBatch *GetBatch) ([]*Data, error) {
	rows, err := s.dbPool.Query(ctx,
		getBatchSQL,
		getBatch.UserID,
		getBatch.Limit,
		getBatch.Offset,
	)
	if err != nil {
		return nil, err
	}

	resp := make([]*Data, getBatch.Limit)

	i := int32(0)
	for rows.Next() {
		row := &Data{}

		id := 0
		err := rows.Scan(&id, &row.Data, &row.DataType, &row.Meta)
		if err != nil {
			return nil, err
		}

		row.Id = strconv.Itoa(id)
		resp[i] = row
		i++
	}

	if i != getBatch.Limit {
		resp = resp[0:i]
	}

	return resp, nil
}

var _ Interface = (*Impl)(nil)

func Init(databaseURI string) Interface {
	conn, err := pgxpool.Connect(context.TODO(), databaseURI)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = conn.Exec(context.TODO(), createTableSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &Impl{
		dbPool: conn,
	}
}
