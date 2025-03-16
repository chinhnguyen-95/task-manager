package middleware

import (
	"context"
	"crypto/rsa"
	"strings"

	"task-manager/pkg/jwtutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const userIDKey contextKey = "userID"

func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey).(string)
	return userID, ok
}

func NewJWTUnaryInterceptor(publicKey *rsa.PublicKey) grpc.UnaryServerInterceptor {
	skipAuth := map[string]bool{
		"/taskmanager.v1.AuthService/Login":    true,
		"/taskmanager.v1.AuthService/Register": true,
	}

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if skipAuth[info.FullMethod] {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		authHeaders := md["authorization"]
		if len(authHeaders) == 0 || !strings.HasPrefix(authHeaders[0], "Bearer ") {
			return nil, status.Error(codes.Unauthenticated, "missing or invalid Authorization header")
		}

		tokenStr := strings.TrimPrefix(authHeaders[0], "Bearer ")
		claims, err := jwtutil.ValidateToken(tokenStr, publicKey)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		userID, ok := claims["sub"].(string)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "invalid subject claim")
		}

		ctx = context.WithValue(ctx, userIDKey, userID)
		return handler(ctx, req)
	}
}
