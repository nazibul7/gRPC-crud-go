package interceptor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	log.Printf("--> request: %s", info.FullMethod)

	resp, err := handler(ctx, req)
	log.Printf(
		"<-- response: %s duration=%s err=%v",
		info.FullMethod,
		time.Since(start),
		err,
	)

	return resp, err
}
