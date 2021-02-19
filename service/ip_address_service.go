package service

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const ipAddressSite = "https://checkip.amazonaws.com"

type IPAddressService interface {
	GetMyIpAddress() (string, error)
}

type ipAddressService struct {
}

func NewIpAddressService() IPAddressService {
	return &ipAddressService{}
}

func (i *ipAddressService) GetMyIpAddress() (string, error) {
	resp, respErr := http.Get(ipAddressSite)
	if respErr != nil {
		return "", respErr
	}
	body, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		return "", respErr
	}
	ipResponse := strings.TrimSuffix(string(body), "\n")
	defer resp.Body.Close()
	return ipResponse, nil
}
