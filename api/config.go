package api

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(cpath string) (*API, error) {
	if cpath == "" {
		ucd, err := os.UserConfigDir()
		if err != nil {
			return nil, err
		}

		cpath = ucd + string(os.PathSeparator) + "purelymail.yaml"
	}

	f, err := os.Open(cpath)
	if err != nil {
		return nil, err
	}

	cdata, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	api := &API{}
	err = yaml.Unmarshal(cdata, &api)
	if err != nil {
		return nil, err
	}

	return api, nil
}
