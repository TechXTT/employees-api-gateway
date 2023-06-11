package auth

import (
	"fmt"

	"github.com/TechXTT/employees-api-gateway/pkg/auth/pb"
	"github.com/TechXTT/employees-api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("Error connecting to auth service:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
