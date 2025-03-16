package auth

import (
	"context"

	"task-manager/dto"
	"task-manager/pkg/keycloak"
)

type Service struct {
	kc *keycloak.Client
}

func NewService(kc *keycloak.Client) *Service {
	return &Service{kc: kc}
}

func (s *Service) Register(ctx context.Context, req dto.RegisterRequest) error {
	return s.kc.CreateUser(ctx, req.Name, req.Email, req.Password)
}

func (s *Service) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	return s.kc.Login(ctx, req.Email, req.Password)
}
