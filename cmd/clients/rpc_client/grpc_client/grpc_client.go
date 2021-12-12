package grpc_client

import (
	//"context"
	//"fmt"
	//bitbucket.org/simbiligosi/user_auth_api/users/repository/grpc_repository
	"google.golang.org/grpc"
)

type GrpcClient struct {
}

type IGrpcClient interface {
}

func New() *GrpcClient {
	return &GrpcClient{}
}

func StartGRPCClient() {
	serverAddress := "localhost:7000"

	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())

	if e != nil {
		panic(e)
	}
	defer conn.Close()

	// client := service.NewRepositoryServiceClient(conn)

	// for i := range [10]int{} {
	// 	repositoryModel := grpc_repository.Repository{
	// 		Id:        int64(i),
	// 		IsPrivate: true,
	// 		Name:      string("Grpc-Demo"),
	// 		UserId:    1245,
	// 	}

	// 	if responseMessage, e := client.Add(context.Background(), &repositoryModel); e != nil {
	// 		panic(fmt.Sprintf("Was not able to insert Record %v", e))
	// 	} else {
	// 		fmt.Println("Record Inserted..")
	// 		fmt.Println(responseMessage)
	// 		fmt.Println("=============================")
	// 	}
	// }
}
