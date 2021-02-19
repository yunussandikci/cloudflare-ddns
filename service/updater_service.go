package service

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	log "github.com/sirupsen/logrus"
	"strings"
)

type UpdaterService interface {
	Start()
}

type updaterService struct {
	cloudFlareService CloudFlareService
	ipAddressService IPAddressService
	lastIpAddress string
	domains []string
	interval uint64
}

func NewUpdaterService(cfService CloudFlareService, ipService IPAddressService, domains []string, interval uint64) UpdaterService {
	return &updaterService{
		cloudFlareService: cfService,
		ipAddressService:  ipService,
		domains: domains,
		interval: interval,
	}
}

func (u *updaterService) Start() {
	u.update()
	s := gocron.NewScheduler()
	scheduleErr := s.Every(u.interval).Minutes().Do(u.update)
	if scheduleErr != nil {
		panic(scheduleErr)
	}
	<- s.Start()
}

func (u *updaterService) update() {
	if u.updateIpAddress() {
		log.Infof("New IP address found: %s", u.lastIpAddress)
		u.updateDomains()
	} else {
		log.Infof("IP Address is not changed.")
	}
}

func (u *updaterService) updateIpAddress() bool {
	ipAddress, ipAddressErr := u.ipAddressService.GetMyIpAddress()
	if ipAddressErr != nil {
		panic(ipAddressErr)
	}
	if ipAddress == u.lastIpAddress {
		return false
	}
	u.lastIpAddress = ipAddress
	return true
}

func (u *updaterService) updateDomains() {
	for _, domain := range u.domains {
		zoneName := fmt.Sprintf("%s.%s",strings.Split(domain, ".")[1], strings.Split(domain, ".")[2])
		zoneId, zoneIdErr := u.cloudFlareService.GetZoneId(zoneName)
		if zoneIdErr != nil {
			panic(zoneIdErr)
		}
		recordId, recordIdErr := u.cloudFlareService.GetDnsRecordId(zoneId, domain)
		if recordIdErr != nil {
			panic(recordIdErr)
		}
		recordUpdateErr := u.cloudFlareService.UpdateDnsRecord(zoneId, recordId, u.lastIpAddress)
		if recordUpdateErr != nil {
			panic(recordUpdateErr)
		}
		log.Infof("DNS Record %s updated with %s", domain, u.lastIpAddress)
	}
}