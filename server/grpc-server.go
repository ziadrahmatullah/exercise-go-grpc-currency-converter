package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/appvalidator"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/middleware"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/usecase"
	"google.golang.org/grpc"
)

func StarGRPCServer() {
	validator := appvalidator.NewAppValidatorImpl()
	appvalidator.SetValidator(validator)

	cu := usecase.NewCurrencyConverterUsecase()
	ch := handler.NewCurrencyConverterHandler(cu, validator)

	list, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Error().Err(err).Msg("error starting tcp server")
	}

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		middleware.LoggerInterceptor,
		middleware.ErrorInterceptor,
	))

	pb.RegisterCurrencyConverterServer(server, ch)

	log.Info().Msg("starting grpc server")

	signCh := make(chan os.Signal, 1)
	signal.Notify(signCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.Serve(list); err != nil {
			signCh <- syscall.SIGINT
		}
	}()
	log.Info().Msg("server started")
	<-signCh
	log.Info().Msg("server stopped")
}
