package storage

import (
	"context"
	"log"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
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

var _ Storager = (*Impl)(nil)

func Init(databaseURI string) Storager {
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
