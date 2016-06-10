package main

import (
	pb "github.com/larskluge/babl/protobuf/messages"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/transport"
)

//
// SERVER
//

type BinaryServer interface {
	IO(context.Context, string, *pb.BinRequest) (*pb.BinReply, error)
	Ping(context.Context, *pb.Empty) (*pb.Pong, error)
}

func RegisterBinaryServer(serviceName string, s *grpc.Server, srv BinaryServer) {
	desc := grpc.ServiceDesc{
		ServiceName: serviceName,
		HandlerType: (*BinaryServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "IO",
				Handler:    _Binary_IO_Handler,
			},
			{
				MethodName: "Ping",
				Handler:    _Binary_Ping_Handler,
			},
		},
		Streams: []grpc.StreamDesc{},
	}
	s.RegisterService(&desc, srv)
}

func _Binary_IO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, _ grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.BinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}

	stream, _ := transport.StreamFromContext(ctx)
	method := stream.Method()

	return srv.(BinaryServer).IO(ctx, method, in)
}

func _Binary_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, _ grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(BinaryServer).Ping(ctx, in)
}
