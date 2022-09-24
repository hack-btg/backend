package grpc

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	CorrelationIDFieldName = "X-Correlation-Id"
)

func CorrelationIDInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	corrID := CorrelationIDFromIncomingMeta(ctx)
	if corrID == "" {
		corrID = uuid.New().String()
	}
	ctx = context.WithValue(ctx, 0, corrID)
	ctx = metadata.AppendToOutgoingContext(ctx, CorrelationIDFieldName, corrID)
	return handler(ctx, req)
}

func CorrelationIDFromIncomingMeta(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if id := md[strings.ToLower(CorrelationIDFieldName)]; len(id) > 0 {
			return id[0]
		}
	}
	return ""
}
