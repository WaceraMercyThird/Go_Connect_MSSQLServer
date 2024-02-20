package dtos

import (
	"time"
	"github.com/WaceraMercyThird/go_mssqldb/pkg/utils/enums"
)

// TransactionToDisplayDto represents the Go equivalent of the TransactionToDisplayDto C# class.
type TransactionToDisplayDto struct {
	TransactionTime   time.Time
	OrderId           string
	Rrn               string
	Currency          string
	InitialBalance    float64
	Amount            float64
	CurrentBalance    *float64
	TransactionType   enums.TransactionType
	Status            enums.TransactionStatus
	Narration         string
	ProviderRef       string
	TransactionStatus enums.TransactionStatus
	Commission        *float64
	TransactionFee    *float64
}

// MerchantInfo represents the Go equivalent of the MerchantInfo C# class.
type MerchantInfo struct {
	FirstName      string
	LastName       string
	Email          string
	AccountNo      string
	CurrentBalance float64
}
