package service

// NOTE: The name of this service is temporal. The features have to be splitted into several files after the construction completed.

import (
	"fmt"
	"os"
	//"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/client"
)

type XenClient struct {}

func NewXenClient() *XenClient {
	return &XenClient{}
}

func (x *XenClient) Init() (*client.API, error) {
	xenSecret := os.Getenv("XENDIT_SECRET")
	if xenSecret == "" {
		return nil, fmt.Errorf("failed to load xendit secret")
	}

	xenCli := client.New(xenSecret)

	return xenCli, nil
}
