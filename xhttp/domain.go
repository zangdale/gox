package xhttp

import (
	"golang.org/x/net/publicsuffix"
)

func MainDomain(domain string) string {
	mainDomain, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return domain
	}
	return mainDomain
}

func Publicsuffix(domain string) string {
	eTLD, _ := publicsuffix.PublicSuffix(domain)
	return eTLD
}
