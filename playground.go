package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/GermanVor/data-keeper/proto/datakeeper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	// ctx := context.Background()

	conn, err := grpc.Dial("localhost:5678", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewDataKeeperClient(conn)

	ctx := metadata.AppendToOutgoingContext(
		context.Background(),
		"jwt",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIn0.TfeIcncXPMAsIX9LUq6voCF8EY-tys5s_fJL3ZYnfXU",
	)

	// for i := range []int{1, 2, 3, 4, 5} {
	// 	client.New(ctx, &pb.NewRequest{
	// 		DataType: pb.DataType_TEXT,
	// 		Data:     []byte("qwe" + strconv.Itoa(i)),
	// 		Meta:     map[string]string{"qwe": strconv.Itoa(10 * i)},
	// 	})
	// }

	resp, err := client.GetBatch(ctx, &pb.GetBatchRequest{Offset: 1, Limit: 1000})

	// Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIn0.TfeIcncXPMAsIX9LUq6voCF8EY-tys5s_fJL3ZYnfXU"
	// resp, err := client.Delete(ctx, &pb.DeleteRequest{Id: "1"})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.DataArray)
	// fmt.Println(resp)
}
