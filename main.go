package main

import (
	"context"
	"fmt"
	"grpc-sample/envRouting"
	"log"
	"net"

	"grpc-sample/person"

	"google.golang.org/grpc"
)

type personServer struct {
	person.UnimplementedPersonServiceServer
}

func (s *personServer) GetAll(ctx context.Context, in *person.Empty) (*person.AllPerson, error) {

	return &person.AllPerson{
		People: []*person.Person{
			{
				Id:   1,
				Name: "John",
				Age:  22,
			},
		},
	}, nil
}

func main() {
	envRouting.LoadEnv()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", envRouting.Port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	person.RegisterPersonServiceServer(s, &personServer{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
