package datakeeperrpc

import (
	"context"
	"errors"

	"github.com/GermanVor/data-keeper/cmd/storageServer/storage"
	pb "github.com/GermanVor/data-keeper/proto/datakeeper"
)

func FormatStorageData(sD storage.DataType) pb.DataType {
	switch sD {
	case storage.BankCard:
		return pb.DataType_BANK_CARD
	case storage.LogPass:
		return pb.DataType_LOGG_PASS
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
	userId, _ := ctx.Value("userId").(string)

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
	userId, _ := ctx.Value("userId").(string)

	data, err := s.stor.Get(ctx, &storage.GetData{UserId: userId, Id: in.Id})
	if err != nil {
		return nil, err
	}

	return &pb.GetResponse{
		Id:       data.Id,
		DataType: FormatStorageData(data.DataType),
		Data:     data.Data,
		Meta:     data.Meta,
	}, nil
}

func (s *DatakeeperServiceImpl) Set(ctx context.Context, in *pb.SetRequest) (*pb.SetResponse, error) {
	userId, _ := ctx.Value("userId").(string)

	data, err := s.stor.Set(ctx, in.Format(userId))
	if err != nil {
		return nil, err
	}

	return &pb.SetResponse{
		Id:       data.Id,
		DataType: FormatStorageData(data.DataType),
		Data:     data.Data,
		Meta:     data.Meta,
	}, nil
}

func (s *DatakeeperServiceImpl) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	userId, _ := ctx.Value("userId").(string)

	ok, err := s.stor.Delete(ctx, &storage.DeleteData{
		UserId: userId,
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
