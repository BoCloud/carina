package runners

import (
	"carina/utils"
	"net"
	"os"

	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type gRPCServerRunner struct {
	srv            *grpc.Server
	sockFile       string
	leaderElection bool
}

var _ manager.LeaderElectionRunnable = gRPCServerRunner{}

// NewGRPCRunner creates controller-runtime's manager.Runnable for a gRPC server.
// The server will listen on UNIX domain socket at sockFile.
// If leaderElection is true, the server will run only when it is elected as leader.
func NewGRPCRunner(srv *grpc.Server, sockFile string, leaderElection bool) manager.Runnable {
	return gRPCServerRunner{srv, sockFile, leaderElection}
}

// Start implements controller-runtime's manager.Runnable.
func (r gRPCServerRunner) Start(ch <-chan struct{}) error {
	if !utils.FileExists(r.sockFile) {
		_ = os.MkdirAll(r.sockFile, os.ModeSocket)
	}
	_ = os.Remove(r.sockFile)
	lis, err := net.Listen("unix", r.sockFile)
	if err != nil {
		return err
	}

	go r.srv.Serve(lis)
	<-ch
	r.srv.GracefulStop()
	return nil
}

// NeedLeaderElection implements controller-runtime's manager.LeaderElectionRunnable.
func (r gRPCServerRunner) NeedLeaderElection() bool {
	return r.leaderElection
}
