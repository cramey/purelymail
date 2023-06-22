package api

import (
	"encoding/json"
	"fmt"
)

type ListDomainResponse struct {
	GenericResponse
	Result struct {
		Domains []Domain `json:"domains"`
	} `json:"result"`
}

type Domain struct {
	Name    string `json:"name"`
	Reset   bool   `json:"allowAccountReset"`
	SubAddr bool   `json:"symbolicSubaddressing"`
	Shared  bool   `json:"isShared"`
	DNS     struct {
		MX    bool `json:"passesMx"`
		SPF   bool `json:"passesSpf"`
		DKIM  bool `json:"passesDkim"`
		DMARC bool `json:"passesDmarc"`
	} `json:"dnsSummary"`
}

func (dom Domain) Summary() string {
	return fmt.Sprintf(
		"Shared=%s Reset=%s SubAddr=%s MX=%s SPF=%s DKIM=%s DMARC=%s",
		yNo(dom.Shared), yNo(dom.Reset), yNo(dom.SubAddr),
		yNo(dom.DNS.MX), yNo(dom.DNS.SPF),
		yNo(dom.DNS.DKIM), yNo(dom.DNS.DMARC),
	)
}

func (api *API) ListDomains(shared bool) (*[]Domain, error) {
	ep := api.endpoint() + "listDomains"
	resp, err := api.exec(ep, "POST", MIME_JSON, map[string]bool{
		"includeShared": false,
	})
	if err != nil {
		return nil, err
	}

	var ldr ListDomainResponse
	err = json.Unmarshal(resp, &ldr)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error: %s", err)
	}

	if ldr.Type != "success" {
		return nil, fmt.Errorf("listRoutingRules: %s %s", ldr.Code, ldr.Message)
	}

	return &ldr.Result.Domains, nil
}
