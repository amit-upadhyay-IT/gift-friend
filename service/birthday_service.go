package service

import (
	"birthdayGiftVoucher/util"
	"fmt"
	"strings"
)

type BirthdayService struct {

}

func NewBirthdayService() *BirthdayService {
	return &BirthdayService{}
}

func (svc *BirthdayService) GetVoucher(name string) (string, error) {
	if validateName(name) {
		return util.InHTMLLinkTag("https://www.amazon.in/gp/css/order-history?ref_=nav_AccountFlyout_orders"), nil
	} else {
		return "", fmt.Errorf("gift is not for you")
	}
}

func (svc *BirthdayService) GiftView() string {
	return util.GetEnterNameHTML()
}

func validateName(name string) bool {
	name = strings.ToLower(name)
	fmt.Println(name)
	switch name {
	case "amit upadhyay":
		return true
	case "amit":
		return true
	}
	return false
}