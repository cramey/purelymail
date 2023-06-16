package api

type API struct {
	Token string
}

func (api *API) endpoint() string {
	return "https://purelymail.com/api/v0/"
}
