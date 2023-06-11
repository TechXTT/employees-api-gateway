package data

import (
	"fmt"

	"github.com/TechXTT/employees-api-gateway/pkg/config"
	"github.com/TechXTT/employees-api-gateway/pkg/data/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.DataServiceClient
}

func InitServiceClient(c *config.Config) pb.DataServiceClient {
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial(c.DataSvcUrl, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("Error connecting to data service:", err)
	}

	return pb.NewDataServiceClient(cc)
}
