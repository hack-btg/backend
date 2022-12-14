package client

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func GetBanks() (Banks, error) {
	resp, err := http.Get(os.Getenv("BANKING_URI"))
	if err != nil {
		log.Println("[getBanks] err ", err.Error())
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[getBanks] err ", err.Error())
		return nil, err
	}
	banks := Banks{}
	return banks, json.Unmarshal(body, &banks)
}

type Banks []Bank

type Bank struct {
	OrganisationID              string      `json:"OrganisationId"`
	Status                      string      `json:"Status"`
	OrganisationName            string      `json:"OrganisationName"`
	CreatedOn                   time.Time   `json:"CreatedOn"`
	LegalEntityName             string      `json:"LegalEntityName"`
	CountryOfRegistration       string      `json:"CountryOfRegistration"`
	CompanyRegister             string      `json:"CompanyRegister"`
	Tag                         interface{} `json:"Tag"`
	Size                        interface{} `json:"Size"`
	RegistrationNumber          string      `json:"RegistrationNumber"`
	RegistrationID              string      `json:"RegistrationId"`
	RegisteredName              string      `json:"RegisteredName"`
	AddressLine1                string      `json:"AddressLine1"`
	AddressLine2                string      `json:"AddressLine2"`
	City                        string      `json:"City"`
	Postcode                    string      `json:"Postcode"`
	Country                     string      `json:"Country"`
	ParentOrganisationReference string      `json:"ParentOrganisationReference"`
	AuthorisationServers        []struct {
		AuthorisationServerID               string      `json:"AuthorisationServerId"`
		AutoRegistrationSupported           bool        `json:"AutoRegistrationSupported"`
		AutoRegistrationNotificationWebhook interface{} `json:"AutoRegistrationNotificationWebhook"`
		SupportsCiba                        bool        `json:"SupportsCiba"`
		SupportsDCR                         bool        `json:"SupportsDCR"`
		APIResources                        []struct {
			APIResourceID         string `json:"ApiResourceId"`
			APIVersion            string `json:"ApiVersion"`
			APIDiscoveryEndpoints []struct {
				APIDiscoveryID string `json:"ApiDiscoveryId"`
				APIEndpoint    string `json:"ApiEndpoint"`
			} `json:"ApiDiscoveryEndpoints"`
			FamilyComplete              bool        `json:"FamilyComplete"`
			APICertificationURI         interface{} `json:"ApiCertificationUri"`
			CertificationStatus         interface{} `json:"CertificationStatus"`
			CertificationStartDate      interface{} `json:"CertificationStartDate"`
			CertificationExpirationDate interface{} `json:"CertificationExpirationDate"`
			APIFamilyType               string      `json:"ApiFamilyType"`
		} `json:"ApiResources"`
		AuthorisationServerCertifications []interface{} `json:"AuthorisationServerCertifications"`
		CustomerFriendlyDescription       string        `json:"CustomerFriendlyDescription"`
		CustomerFriendlyLogoURI           string        `json:"CustomerFriendlyLogoUri"`
		CustomerFriendlyName              string        `json:"CustomerFriendlyName"`
		DeveloperPortalURI                string        `json:"DeveloperPortalUri"`
		TermsOfServiceURI                 string        `json:"TermsOfServiceUri"`
		NotificationWebhookAddedDate      interface{}   `json:"NotificationWebhookAddedDate"`
		OpenIDDiscoveryDocument           string        `json:"OpenIDDiscoveryDocument"`
		Issuer                            interface{}   `json:"Issuer"`
		PayloadSigningCertLocationURI     string        `json:"PayloadSigningCertLocationUri"`
		ParentAuthorisationServerID       interface{}   `json:"ParentAuthorisationServerId"`
		DeprecatedDate                    interface{}   `json:"DeprecatedDate"`
		RetirementDate                    interface{}   `json:"RetirementDate"`
		SupersededByAuthorisationServerID interface{}   `json:"SupersededByAuthorisationServerId"`
	} `json:"AuthorisationServers"`
	OrgDomainClaims []struct {
		AuthorisationDomainName string `json:"AuthorisationDomainName"`
		AuthorityName           string `json:"AuthorityName"`
		RegistrationID          string `json:"RegistrationId"`
		Status                  string `json:"Status"`
	} `json:"OrgDomainClaims"`
	OrgDomainRoleClaims []struct {
		Status              string        `json:"Status"`
		AuthorisationDomain string        `json:"AuthorisationDomain"`
		Role                string        `json:"Role"`
		Authorisations      []interface{} `json:"Authorisations"`
		RegistrationID      string        `json:"RegistrationId"`
	} `json:"OrgDomainRoleClaims"`
}
