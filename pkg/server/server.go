// package server exposes a GRPC server builder for zanzi
package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	zanzi "github.com/sourcenetwork/zanzi"
	"github.com/sourcenetwork/zanzi/pkg/api"
)

// Server initializes a GRPC server for Zanzi
type Server struct {
	server   *grpc.Server
	listener *net.TCPListener
	addr     string
}

func (s *Server) registerGRPCServices(z *zanzi.Zanzi) {
	s.server.RegisterService(
		&api.RelationGraph_ServiceDesc,
		z.GetRelationGraphService(),
	)
	s.server.RegisterService(
		&api.PolicyService_ServiceDesc,
		z.GetPolicyService(),
	)
}

func (s *Server) initSocket() error {
	addr, err := net.ResolveTCPAddr("tcp", s.addr)
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	s.listener = listener

	return nil
}

func NewServer(address string) Server {
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	return Server{
		addr:   address,
		server: server,
	}
}

func (s *Server) Init(z *zanzi.Zanzi) error {
	s.registerGRPCServices(z)

	if err := s.initSocket(); err != nil {
		return fmt.Errorf("failed initializing sever: %w", err)
	}
	z.GetLogger().Infof("GRPC Service listening on %v", s.listener.Addr())

	return nil
}

// Run starts serving requests by accepting connections.
// Run is a blocking call
func (s *Server) Run() {
	s.server.Serve(s.listener)
}

func NewGRPCGatewayServer(grpcAddress, restAddress string) (*http.Server, error) {

	ctx := context.Background()
	// Create a client connection to the gRPC server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	conn, err := grpc.DialContext(
		ctx,
		grpcAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("dial to gRPC server %s, %w", grpcAddress, err)
	}

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err = api.RegisterPolicyServiceHandler(ctx, mux, conn)
	if err != nil {
		return nil, fmt.Errorf("register ring service handler, %w", err)
	}

	err = api.RegisterRelationGraphHandler(ctx, mux, conn)
	if err != nil {
		return nil, fmt.Errorf("register utility service handler, %w", err)
	}

	// Register Zanzi Services.
	err = api.RegisterPolicyServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return nil, fmt.Errorf("register ring service handler from endpoint, %w", err)
	}

	err = api.RegisterRelationGraphHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return nil, fmt.Errorf("register utility service handler from endpoint, %w", err)
	}

	gw := &http.Server{
		Addr:    restAddress,
		Handler: mux,
	}

	return gw, nil
}
