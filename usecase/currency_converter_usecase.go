package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-currency-converter/pb"
)

type CurrencyConverterUsecase interface {
	ConvertCurrency(dto.CurrencyConversionRequest) (*pb.CurrencyConversionResponse, error)
}

type currencyConverterUsecase struct{}

func NewCurrencyConverterUsecase() CurrencyConverterUsecase {
	return &currencyConverterUsecase{}
}

func (u *currencyConverterUsecase) ConvertCurrency(req dto.CurrencyConversionRequest) (*pb.CurrencyConversionResponse, error) {
	rate1, ok1 := constant.ConversionRate[req.SourceCurrency]
	if !ok1 {
		return nil, apperror.ErrUnknownSource
	}
	rate2, ok2 := constant.ConversionRate[req.TargetCurrency]
	if !ok2 {
		return nil, apperror.ErrUnknownTarget
	}
	result := req.Amount / rate1 * rate2
	return &pb.CurrencyConversionResponse{ConvertedAmount: result}, nil
}
