package service

// NOTE: The name of this service is temporal. The features have to be splitted into several files after the construction completed.

import (
	"fmt"
	"os"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/$product$"
)

type xenCli struct {}

func NewXenCli() xenCli {
	return &xenCli{}
}

func (x *XenCli) init() {
	xendit.Opt.SecretKey = os.Getenv("XENDIT_SECRET")
}

func (x *XenCli) readAvailbaleBalance() {
	data := balance.GetParams{
		AccountType: "CASH",
	}

	resp, err := balance.Get(&data)
	if err != nil {
		fmt.Errorf("unable to read balance")
		return
	}

	fmt.Printf("balance: %+v\n", resp)
}
