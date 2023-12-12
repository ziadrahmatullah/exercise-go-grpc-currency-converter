package middleware

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/apperror"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	res, err := handler(ctx, req)
	if err != nil {
		// if errors.Is(err, context.DeadlineExceeded) {
		// 	return nil, status.Error(codes.DeadlineExceeded, err.Error())
		// }
		switch e := err.(type) {
		case *apperror.CustomError:
			return nil, status.Error(codes.Code(e.Code), e.Message)
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return res, nil
}
