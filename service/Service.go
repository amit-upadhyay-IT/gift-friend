package service

type Service interface {
	GetVoucher(string) (string, error)
	GiftView() string
}
