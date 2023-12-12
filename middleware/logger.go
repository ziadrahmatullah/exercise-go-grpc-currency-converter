package middleware

import (
	"context"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	log := logger.NewLogger()
	startTime := time.Now()

	res, err := handler(ctx, req)
	s := status.Code(err)

	param := map[string]interface{}{
		"status_code": s,
		"method":      info.FullMethod,
		"latency":     time.Since(startTime),
	}
	if err == nil {
		log.Info(param)
	} else {
		param["errors"] = err.Error()
		log.Errorf("", param)
	}

	return res, err
}
