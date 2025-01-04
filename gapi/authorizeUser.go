package gapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/duythien2212/simple_bank/token"
	"google.golang.org/grpc/metadata"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (server *Server) authorizeUser(ctx context.Context) (payload *token.Payload, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}
	value := md.Get(authorizationHeader)
	if len(value) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	authHeader := value[0]
	field := strings.Fields(authHeader)
	if len(field) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(field[0])
	if authType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type")
	}
	accessToken := field[1]
	payload, err = server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token")
	}

	return payload, err
}
