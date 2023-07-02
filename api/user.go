package api

import (
	"encoding/json"
	"fmt"
)

type UserConfig struct {
	User                     string `json:"userName"`
	Domain                   string `json:"domainName"`
	Password                 string `json:"password"`
	Reset                    bool   `json:"enablePasswordReset"`
	RecoveryEmail            string `json:"recoveryEmail,omitempty"`
	RecoveryEmailDescription string `json:"recoveryEmailDescription,omitempty"`
	RecoveryPhone            string `json:"recoveryPhone,omitempty"`
	RecoveryPhoneDescription string `json:"recoveryPhoneDescription,omitempty"`
	Indexing                 bool   `json:"enableSearchIndexing"`
	Welcome                  bool   `json:"sendWelcomeEmail"`
}

func (api *API) CreateUser(ucfg UserConfig) error {
	ep := api.endpoint() + "createUser"
	resp, err := api.exec(ep, "POST", MIME_JSON, ucfg)
	if err != nil {
		return err
	}

	var gr GenericResponse
	if err := json.Unmarshal(resp, &gr); err != nil {
		return fmt.Errorf("json unmarshal error: %s", err)
	}

	if gr.Type != "success" {
		return fmt.Errorf("createUser: %s %s", gr.Code, gr.Message)
	}

	return nil
}

func (api *API) DeleteUser(email string) error {
	ep := api.endpoint() + "deleteUser"
	resp, err := api.exec(ep, "POST", MIME_JSON, map[string]string{
		"userName": email,
	})
	if err != nil {
		return err
	}

	var gr GenericResponse
	if err := json.Unmarshal(resp, &gr); err != nil {
		return fmt.Errorf("json unmarshal error: %s", err)
	}

	if gr.Type != "success" {
		return fmt.Errorf("deleteUser: %s %s", gr.Code, gr.Message)
	}

	return nil
}
