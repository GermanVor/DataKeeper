package datakeeper

import (
	"context"
	"errors"

	storage "github.com/GermanVor/data-keeper/internal/storage"
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
	newData, err := in.Format()
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
	data, err := s.stor.Get(ctx, &storage.GetData{UserId: in.UserId, Id: in.Id})
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
	data, err := s.stor.Set(ctx, in.Format())
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
	ok, err := s.stor.Delete(ctx, &storage.DeleteData{
		UserId: in.UserId,
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
