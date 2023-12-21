package middleware

import (
	"apartment/internal/utils"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	allowedMethods []string
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(methods []string) *AuthInterceptor {
	return &AuthInterceptor{allowedMethods: methods}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		// Kiểm tra xem phương thức này có cần kiểm tra token hay không
		if !interceptor.isAllowedMethod(info.FullMethod) {
			return handler(ctx, req)
		}

		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	// accessibleRoles, ok := interceptor.accessibleRoles[method]

	// if !ok {
	// 	// everyone can access
	// 	return nil
	// }

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}
	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}
	accessToken := values[0]
	claims, err := utils.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	if claims == nil {
		return status.Error(codes.PermissionDenied, "no permission to access this RPC")
	}
	// for _, role := range accessibleRoles {
	// 	if role == cla {

	// 	}
	// }
	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
func (interceptor *AuthInterceptor) isAllowedMethod(method string) bool {
	for _, allowedMethod := range interceptor.allowedMethods {
		if method == allowedMethod {
			return true
		}
	}

	return false
}
