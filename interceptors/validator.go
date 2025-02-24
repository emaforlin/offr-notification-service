package interceptors

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func UnaryValidator(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if err := protovalidate.Validate(req.(proto.Message)); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return handler(ctx, req)
}
