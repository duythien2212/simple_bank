package gapi

import (
	"fmt"

	db "github.com/duythien2212/simple_bank/db/sqlc"
	"github.com/duythien2212/simple_bank/pb"
	"github.com/duythien2212/simple_bank/token"
	"github.com/duythien2212/simple_bank/util"
	"github.com/duythien2212/simple_bank/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributer worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store, taskDistributer worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributer: taskDistributer,
	}

	return server, nil
}
