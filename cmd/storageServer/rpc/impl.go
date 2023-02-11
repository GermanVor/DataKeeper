package datakeeperrpc

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/GermanVor/data-keeper/cmd/storageServer/storage"
	"github.com/GermanVor/data-keeper/internal/common"
	pb "github.com/GermanVor/data-keeper/proto/datakeeper"
)

func FormatStorageData(sD storage.DataType) pb.DataType {
	switch sD {
	case storage.BankCard:
		return pb.DataType_BANK_CARD
	case storage.LogPass:
		return pb.DataType_LOG_PASS
	case storage.Other:
		return pb.DataType_OTHER
	case storage.Text:
		return pb.DataType_TEXT
	default:
		return pb.DataType_UNSPECIFIED
	}
}

type DatakeeperServiceImpl struct {
	pb.UnimplementedDataKeeperServer
	stor storage.Interface
}

func (s *DatakeeperServiceImpl) New(ctx context.Context, in *pb.NewRequest) (*pb.NewResponse, error) {
	userId, _ := ctx.Value(common.USER_ID_CTX_NAME).(string)

	newData, err := in.Format(userId)
	if err != nil {
		return nil, err
	}

	id, err := s.stor.New(ctx, newData)
	if err != nil {
		return nil, err
	}

	return &pb.NewResponse{
		Id: id,
	}, nil
}

func (s *DatakeeperServiceImpl) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	userId, _ := ctx.Value(common.USER_ID_CTX_NAME).(string)

	data, err := s.stor.Get(ctx, &storage.GetData{UserID: userId, Id: in.Id})
	if err != nil {
		return nil, err
	}

	resp := &pb.GetResponse{
		Data: &pb.Data{
			Id:       data.Id,
			DataType: FormatStorageData(data.DataType),
			Data:     data.Data,
			Meta:     make(map[string]string),
		},
	}

	err = json.Unmarshal(data.Meta, &resp.Data.Meta)

	return resp, err
}

func (s *DatakeeperServiceImpl) GetBatch(ctx context.Context, in *pb.GetBatchRequest) (*pb.GetBatchResponse, error) {
	userId, _ := ctx.Value(common.USER_ID_CTX_NAME).(string)

	batch, err := s.stor.GetBatch(
		ctx,
		&storage.GetBatch{UserID: userId, Offset: in.Offset, Limit: in.Limit},
	)
	if err != nil {
		return nil, err
	}

	resp := &pb.GetBatchResponse{
		DataArray: make([]*pb.Data, len(batch)),
	}

	for i, row := range batch {
		resp.DataArray[i] = &pb.Data{
			Id:       row.Id,
			DataType: pb.DataType(row.DataType),
			Data:     row.Data,
			Meta:     make(map[string]string),
		}

		err := json.Unmarshal(row.Meta, &resp.DataArray[i].Meta)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func (s *DatakeeperServiceImpl) Set(ctx context.Context, in *pb.SetRequest) (*pb.SetResponse, error) {
	userId, _ := ctx.Value(common.USER_ID_CTX_NAME).(string)

	req, err := in.Format(userId)
	if err != nil {
		return nil, err
	}

	data, err := s.stor.Set(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &pb.SetResponse{
		Data: &pb.Data{
			Id:       data.Id,
			DataType: FormatStorageData(data.DataType),
			Data:     data.Data,
			Meta:     make(map[string]string),
		},
	}

	return resp, json.Unmarshal(data.Meta, &resp.Data.Meta)
}

func (s *DatakeeperServiceImpl) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	userId, _ := ctx.Value(common.USER_ID_CTX_NAME).(string)

	ok, err := s.stor.Delete(ctx, &storage.DeleteData{
		UserID: userId,
		Id:     in.Id,
	})
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("todo")
	}

	return &pb.DeleteResponse{}, nil
}

func Init(stor storage.Interface) *DatakeeperServiceImpl {
	return &DatakeeperServiceImpl{
		stor: stor,
	}
}
