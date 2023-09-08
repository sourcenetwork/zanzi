// package server exposes a GRPC server builder for zanzi
package server

import (
    "net"
    "fmt"

    "google.golang.org/grpc"

    zanzi "github.com/sourcenetwork/zanzi"
    "github.com/sourcenetwork/zanzi/pkg/api"
)

// Server initializes a GRPC server for Zanzi
type Server struct {
    server *grpc.Server
    listener *net.TCPListener
    addr string
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
        addr: address,
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
