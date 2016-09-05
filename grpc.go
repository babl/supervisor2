//go:generate go-bindata data/...

package main

import (
	"crypto/tls"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/transport"
)

func opts() (opts []grpc.ServerOption) {
	certPEMBlock, _ := Asset("data/server.pem")
	keyPEMBlock, _ := Asset("data/server.key")
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	check(err)
	creds := credentials.NewServerTLSFromCert(&cert)
	opts = append(opts, grpc.Creds(creds))
	opts = append(opts, grpc.MaxMsgSize(MaxGrpcMessageSize))
	return
}

func NewServer() *grpc.Server {
	return grpc.NewServer(opts()...)
}

func MethodFromContext(ctx context.Context) string {
	stream, ok := transport.StreamFromContext(ctx)
	if !ok {
		panic("babl-server: grpc: MethodFromContext: No Stream in Context")
	}
	return stream.Method()
}
