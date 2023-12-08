package dreamsoft

import (
	"kubesphere.io/kubesphere/pkg/apiserver/authentication/identityprovider"
	"kubesphere.io/kubesphere/pkg/server/options"
	"net/http"
)

func init() {
	identityprovider.RegisterOAuthProvider(&dreamsoftProviderFactory{})
}

type dreamsoft struct {
	RedirectURL string `json:"redirectURL" yaml:"redirectURL"`
	ServerURL   string `json:"serverURL" yaml:"serverURL"`
}

type dreamsoftProviderFactory struct {
}

type dreamsoftIdentity struct {
	User string `json:"user"`
}

func (c dreamsoftIdentity) GetUserID() string {
	return c.User
}

func (c dreamsoftIdentity) GetUsername() string {
	return c.User
}

func (c dreamsoftIdentity) GetEmail() string {
	return ""
}

func (f dreamsoftProviderFactory) Type() string {
	return "DreamSoftProviderFactory"
}

func (f dreamsoftProviderFactory) Create(opts options.DynamicOptions) (identityprovider.OAuthProvider, error) {
	var dreamsoft dreamsoft
	return &dreamsoft, nil
}

func (c dreamsoft) IdentityExchangeCallback(req *http.Request) (identityprovider.Identity, error) {
	return &dreamsoftIdentity{User: "admin"}, nil
}
