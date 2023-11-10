package model

import "github.com/go-playground/validator/v10"

type Transaction struct {
	ID           string  `json:"id" validate:"required,uuid4"`
	AccountID    string  `json:"accountId" validate:"required,uuid4"`
	Amount       float64 `json:"amount" validate:"required,numeric"`
	PixKeyTo     string  `json:"pixKeyTo" validate:"required"`
	PixKeyKindTo string  `json:"pixKeyKindTo" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Status       string  `json:"status" validate:"required"`
	Error        string  `json:"description"`
}

func (t *Transaction) isValid() error {
	v := validator.New()
	err := v.Struct(t)
	if err != nil {
		log.Errorf("Error during Transaction validation: %s", err.Error())
		return err
	}
	return nil
}
