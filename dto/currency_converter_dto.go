package dto

type CurrencyConversionRequest struct {
	Amount         float64 `validate:"required,min=0"`
	SourceCurrency string  `validate:"required"`
	TargetCurrency string  `validate:"required"`
}
