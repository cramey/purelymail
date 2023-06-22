package api

import (
	"encoding/json"
	"fmt"
)

type ListRoutingRuleResponse struct {
	GenericResponse
	Result struct {
		Rules []RoutingRule `json:"rules"`
	} `json:"result"`
}

type RoutingRule struct {
	ID        int64    `json:"id,omitempty"`
	Domain    string   `json:"domainName"`
	Prefix    bool     `json:"prefix"`
	MatchUser string   `json:"matchUser"`
	Addresses []string `json:"targetAddresses"`
}

func (api *API) ListRoutingRules() (*[]RoutingRule, error) {
	ep := api.endpoint() + "listRoutingRules"
	resp, err := api.exec(ep, "POST", MIME_JSON, map[string]string{})
	if err != nil {
		return nil, err
	}

	var rrr ListRoutingRuleResponse
	err = json.Unmarshal(resp, &rrr)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error: %s", err)
	}

	if rrr.Type != "success" {
		return nil, fmt.Errorf("listRoutingRules: %s %s", rrr.Code, rrr.Message)
	}

	return &rrr.Result.Rules, nil
}

func (api *API) CreateRoutingRule(domain, user string, prefix bool, addresses []string) error {
	rr := RoutingRule{
		Domain: domain, MatchUser: user, Prefix: prefix, Addresses: addresses,
	}
	ep := api.endpoint() + "createRoutingRule"
	resp, err := api.exec(ep, "POST", MIME_JSON, rr)
	if err != nil {
		return err
	}

	var gr GenericResponse
	err = json.Unmarshal(resp, &gr)
	if err != nil {
		return fmt.Errorf("json unmarshal error: %s", err)
	}

	if gr.Type != "success" {
		return fmt.Errorf("createRoutingRule: %s %s", gr.Code, gr.Message)
	}
	return nil
}

func (api *API) DeleteRoutingRule(id int64) error {
	req := map[string]int64{"routingRuleId": id}
	ep := api.endpoint() + "deleteRoutingRule"
	resp, err := api.exec(ep, "POST", MIME_JSON, req)
	if err != nil {
		return err
	}

	var gr GenericResponse
	err = json.Unmarshal(resp, &gr)
	if err != nil {
		return fmt.Errorf("json unmarshal error: %s", err)
	}

	if gr.Type != "success" {
		return fmt.Errorf("deleteRoutingRule: %s %s", gr.Code, gr.Message)
	}
	return nil
}
