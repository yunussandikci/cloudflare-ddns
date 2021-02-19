package service

import (
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/pkg/errors"
)

type CloudFlareService interface {
	GetZoneId(zoneName string) (string, error)
	GetDnsRecordId(zoneId string, recordName string) (string, error)
	UpdateDnsRecord(zoneId string, recordId, content string) error
}

type cloudFlareService struct {
	client *cloudflare.API
}

func NewCloudFlareService(apiToken string) CloudFlareService {
	client, clientErr := cloudflare.NewWithAPIToken(apiToken)
	if clientErr != nil {
		panic(clientErr)
	}
	return &cloudFlareService{
		client: client,
	}
}

func (c *cloudFlareService) GetZoneId(zoneName string) (string, error) {
	return c.client.ZoneIDByName(zoneName)
}

func (c *cloudFlareService) GetDnsRecordId(zoneId string, recordName string) (string, error) {
	records, dnsRecordsErr := c.client.DNSRecords(zoneId, cloudflare.DNSRecord{
		Name: recordName,
	})
	if dnsRecordsErr != nil {
		return "", dnsRecordsErr
	}
	if len(records) == 0 {
		return "", errors.New(fmt.Sprintf("There is no record found for %s", recordName))
	}
	return records[0].ID, nil
}

func (c *cloudFlareService) UpdateDnsRecord(zoneId string, recordId, content string) error {
	updateDnsRecordErr := c.client.UpdateDNSRecord(zoneId, recordId, cloudflare.DNSRecord{Content: content})
	if updateDnsRecordErr != nil {
		return updateDnsRecordErr
	}
	return nil
}
