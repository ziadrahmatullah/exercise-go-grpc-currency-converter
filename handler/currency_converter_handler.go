package handler

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/appvalidator"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/usecase"
)

type CurrencyConverterHandler struct {
	pb.UnimplementedCurrencyConverterServer
	usecase   usecase.CurrencyConverterUsecase
	validator appvalidator.AppValidator
}

func NewCurrencyConverterHandler(u usecase.CurrencyConverterUsecase, valid appvalidator.AppValidator) *CurrencyConverterHandler {
	return &CurrencyConverterHandler{
		usecase:   u,
		validator: valid,
	}
}

func (h *CurrencyConverterHandler) ConvertCurrency(ctx context.Context, req *pb.CurrencyConversionRequest) (*pb.CurrencyConversionResponse, error) {
	currencyReq := dto.CurrencyConversionRequest{
		Amount:         req.Amount,
		SourceCurrency: req.SourceCurrency,
		TargetCurrency: req.TargetCurrency,
	}
	err := h.validator.Validate(currencyReq)
	if err != nil {
		return nil, apperror.ErrInvalidBody
	}
	res, err := h.usecase.ConvertCurrency(currencyReq)
	if err != nil {
		return nil, err
	}
	return res, nil
}
