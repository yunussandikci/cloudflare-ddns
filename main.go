package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/yunussandikci/cloudflare-dnynamic-dns/service"
	"os"
	"strconv"
	"strings"
)

func main() {
	cloudFlareToken := os.Getenv("CLOUDFLARE_TOKEN")
	domains := strings.Split(strings.ReplaceAll(os.Getenv("DOMAINS"), " ", ""), ",")
	interval, intervalErr := strconv.Atoi(os.Getenv("INTERVAL"))
	if intervalErr != nil {
		panic(intervalErr)
	}
	cloudFlareService := service.NewCloudFlareService(cloudFlareToken)
	ipAddressService := service.NewIpAddressService()
	updaterService := service.NewUpdaterService(cloudFlareService, ipAddressService, domains, uint64(interval))

	log.Infof("Cloudflare Dynamic DNS Service started.")
	log.Infof("Domains: %s", strings.Join(domains,", "))
	log.Infof("Interval: %d minutes", interval)
	updaterService.Start()
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:             true,
	})
}