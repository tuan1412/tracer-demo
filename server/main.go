package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tracer-demo/proto"
)

type server struct {
	locations map[string]uint64
}

func (s *server) GetCurrentWeather(c context.Context, in *proto.WeatherRequest) (*proto.WeatherResponse, error) {
	temperature, ok := s.locations[in.Location]

	if !ok {
		err := status.Error(codes.NotFound, "location not found")
		return nil, err
	}

	return &proto.WeatherResponse{
		Temperature: temperature,
	}, nil
}

func main() {
	server := &server{
		locations: map[string]uint64{
			"HaNoi":  30,
			"HCM":    35,
			"DaNang": 22,
		},
	}

	s := grpc.NewServer()
}
