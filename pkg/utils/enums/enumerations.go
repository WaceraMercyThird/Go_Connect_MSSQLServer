package enums

import (
	//"encoding/json"
	"fmt"
)

// Enumerations

type IdType int

const (
	NationalId IdType = iota + 1
	Passport
	Alien
	Military
)

func (t IdType) String() string {
	return [...]string{"NationalId", "Passport", "Alien", "Military"}[t-1]
}

type ProviderType int

const (
	Bank ProviderType = iota + 1
	MOMO
	Service
)

func (t ProviderType) String() string {
	return [...]string{"Bank", "MOMO", "Service"}[t-1]
}

type ChargeValueType int

const (
	Percentage ChargeValueType = iota + 1
	Fixed
)

func (t ChargeValueType) String() string {
	return [...]string{"Percentage", "Fixed"}[t-1]
}

type OfflineTopuSource int

const (
	Internal OfflineTopuSource = iota + 1
	BankDeposit
	PayoutReversal
)

func (t OfflineTopuSource) String() string {
	return [...]string{"Internal", "BankDeposit", "PayoutReversal"}[t-1]
}

type ServiceType int

const (
	KPLC ServiceType = iota + 1
	AIRTIME
	WATER
	PAYTV
	//IPRS
	DATA
	INTERNET
	BILLER
	P2P
	MERCHANTPAYMENT
	TOPUP
	PAYOUT
)

func (t ServiceType) String() string {
	return [...]string{"KPLC", "AIRTIME", "WATER", "PAYTV", "IPRS", "DATA", "INTERNET", "BILLER", "P2P", "MERCHANTPAYMENT", "TOPUP", "PAYOUT"}[t-1]
}

type ChargeCategory int



func (t ChargeCategory) String() string {
	return [...]string{"Commission", "TransactionFee"}[t-1]
}

type CostBurden int

const (
	Receiver CostBurden = iota + 1
	Sender
	Shared
)

func (t CostBurden) String() string {
	return [...]string{"Receiver", "Sender", "Shared"}[t-1]
}

type TransactionStatus int

const (
	Pending TransactionStatus = iota + 1
	Success
	Canceled
	Fail
	TimeOut
	Reversed
	Reversing
	Processing
	NotFound
)

func (t TransactionStatus) String() string {
	return [...]string{"Pending", "Success", "Canceled", "Fail", "TimeOut", "Reversed", "Reversing", "Processing", "NotFound"}[t-1]
}

type Gender int

const (
	Male Gender = iota + 1
	Female
	Others
)

func (t Gender) String() string {
	return [...]string{"Male", "Female", "Others"}[t-1]
}

type TransactionType int

func (t TransactionType) String() string {
	return [...]string{"Withdrawal", "Transfer", "Reversal", "Balance", "Debit", "Credit", "MobileTransaction", "WalletTransaction", "Card"}[t-1]
}

type OldProvider int


func (t OldProvider) String() string {
	return [...]string{"Mpesa", "TKash", "AirtelMoney", "Wallet", "Card", "Equity", "NBK", "COOP", "IPLS", "KCB", "MPESAB2B", "GOTV", "DSTV", "STARTIME", "ZUKU", "SAFARICOM", "TELKOM", "AIRTEL", "SMS", "IPRS", "FAIBA", "NAIROBIWATER", "OFFLINE"}[t-1]
}

type ProviderTable int

const (
	Wallet ProviderTable = iota + 1
	Mpesa
	Tkash
	AirtelMoney
	Card
)

func (t ProviderTable) String() string {
	return [...]string{"Wallet", "Mpesa", "Tkash", "AirtelMoney", "Card"}[t-1]
}

type AccountType int

const (
	Whitelabel AccountType = iota + 1
	Individual
	Business
	Main
)

func (t AccountType) String() string {
	return [...]string{"Whitelabel", "Individual", "Business", "Main"}[t-1]
}

type ChargedFor int

const (
	Initial ChargedFor = iota + 1
	Commission
	TransactionFee
	Reversal
)

func (t ChargedFor) String() string {
	return [...]string{"Initial", "Commission", "TransactionFee", "Reversal"}[t-1]
}

type Currency int

const (
	KES Currency = iota + 1
	USD
	TZS
	UGX
	SSP
	SOS
	NGN
)

func (t Currency) String() string {
	return [...]string{"KES", "USD", "TZS", "UGX", "SSP", "SOS", "NGN"}[t-1]
}

type ReportType int

const (
	Daily ReportType = iota + 1
	Weekly
	Monthly
	Annually
)

func (t ReportType) String() string {
	return [...]string{"Daily", "Weekly", "Monthly", "Annually"}[t-1]
}

type Format int

const (
	CSV Format = iota + 1
	Excel
	Pdf
)

func (t Format) String() string {
	return [...]string{"CSV", "Excel", "Pdf"}[t-1]
}

type ReturnType int

const (
	Download ReturnType = iota + 1
	Email
)

func (t ReturnType) String() string {
	return [...]string{"Download", "Email"}[t-1]
}

// Example Usage:

func main() {
	// Example using IdType
	idType := NationalId
	fmt.Println(idType) // Output: NationalId

	// Example using ProviderType
	providerType := Bank
	fmt.Println(providerType) // Output: Bank

	// ... and so on for other enums
}
