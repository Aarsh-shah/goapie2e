package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main2() {
	grpc_middleware.ChainStreamClient()
	listener, err := net.Listen("tcp", ":8085")
	if err != nil {
		fmt.Println(err)

	}
	server := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer()))
	service := NewAarshServer()
	server.RegisterService(&TardisFeatureService_ServiceDesc, service)
	reflection.Register(server)
	fmt.Println("Serving on port 8080")
	err = server.Serve(listener)
	if err != nil {
		fmt.Println(err)
	}
}

type AarshIf interface {
	GetFeatures(context.Context, *TardisGetFeatureRequest) (*TardisGetFeaturesResponse, error)
	GetCounters(context.Context, *TardisGetCountersRequest) (*TardisGetCountersResponse, error)
}

type AarshServerImpl struct {
	// You can embed the UnimplementedTardisFeatureServiceServer to ensure compatibility
	// with future versions of the generated code.
	UnimplementedTardisFeatureServiceServer
}

func NewAarshServer() *AarshServerImpl {
	return &AarshServerImpl{}
}

// Implement the GetFeatures method as defined in the interface
func (a AarshServerImpl) GetFeatures(ctx context.Context, req *TardisGetFeatureRequest) (*TardisGetFeaturesResponse, error) {
	//time.Sleep(20 * time.Millisecond)
	return &TardisGetFeaturesResponse{Features: req.FeatureName + req.FeatureName}, nil
}

// Implement the GetCounters method as defined in the interface
func (a AarshServerImpl) GetCounters(ctx context.Context, req *TardisGetCountersRequest) (*TardisGetCountersResponse, error) {
	return &TardisGetCountersResponse{
		Counters: int32(len(req.Counter)),
	}, nil
}
