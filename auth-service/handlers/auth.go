package handlers

import (
	"context"

	pb "github.com/murillio4/qa-micro/auth-service/proto"
	"github.com/murillio4/qa-micro/auth-service/repositories/permission"
)

type AuthService struct {
	perm permission.Repository
}

// CreateTokens Signs in a user by returning jwt tokens if LoginRequest is valid
func (h *AuthService) CreateTokens(ctx context.Context, in *pb.LoginRequest, out *pb.Tokens) error {

}

func (h *AuthService) ValidateAuthToken(ctx context.Context, in *pb.Token, out *pb.Claims) error {

}

func (h *AuthService) RefreshAuthToken(ctx context.Context, in *pb.Token, out *pb.Tokens) error {

}

func (h *AuthService) DeleteRefreshToken(ctx context.Context, in *pb.Token, out *pb.Empty) error {

}
