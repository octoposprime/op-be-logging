package presentation

import (
	"net"

	pp "github.com/octoposprime/op-be-logging/internal/application/presentation/port"
	pb_error "github.com/octoposprime/op-be-shared/pkg/proto/pb/error"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tgrpc "github.com/octoposprime/op-be-shared/tool/grpc"
	"google.golang.org/grpc"
)

// Grpc is the gRPC API for the application
type Grpc struct {
	pb_error.UnimplementedErorrSvcServer
	pb_logging.UnimplementedLoggingSvcServer
	queryHandler   pp.QueryPort
	commandHandler pp.CommandPort
}

// NewGrpc creates a new instance of Grpc
func NewGrpc(qh pp.QueryPort, ch pp.CommandPort) *Grpc {
	api := &Grpc{
		queryHandler:   qh,
		commandHandler: ch,
	}
	return api
}

// Serve starts the API server
func (a *Grpc) Serve(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(tgrpc.Interceptor),
	)
	pb_error.RegisterErorrSvcServer(s, a)
	pb_logging.RegisterLoggingSvcServer(s, a)
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
