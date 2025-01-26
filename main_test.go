package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func Test(t *testing.T) {
	client, _ := grpc.Dial("localhost:8085", grpc.WithInsecure())
	conn := NewTardisFeatureServiceClient(client)
	cancelContext, _ := context.WithTimeout(context.Background(), 15*time.Millisecond)
	resp, err := conn.GetFeatures(cancelContext, &TardisGetFeatureRequest{
		FeatureName: "Aarsh",
	})
	fmt.Println(resp.Features)
	fmt.Println("Err ", err)
}
