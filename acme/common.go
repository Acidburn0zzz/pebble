package acme

import "gopkg.in/square/go-jose.v2"

// acme.Resource values identify different types of ACME resources
type Resource string

const (
	ResourceNewNonce = Resource("new-nonce")
	ResourceNewReg   = Resource("new-reg")
	ResourceNewOrder = Resource("new-order")
	// TODO(@cpu): Should there be a resource for challenge or just use 'authz'?
	ResourceChallenge = Resource("challenge")
)

const (
	StatusPending = "pending"
	StatusInvalid = "invalid"
	StatusValid   = "valid"

	IdentifierDNS = "dns"

	ChallengeHTTP01 = "http-01"

	HTTP01BaseURL = ".well-known/acme-challenge/"
)

type Identifier struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// TODO(@cpu) - Rename Registration to Account, update refs
type Registration struct {
	Status    string           `json:"status"`
	Key       *jose.JSONWebKey `json:"key"`
	Contact   []string         `json:"contact"`
	ToSAgreed bool             `json:"terms-of-service-agreed"`
	Orders    string           `json:"orders"`
}

// An Order is created to request issuance for a CSR
type Order struct {
	Status         string   `json:"status"`
	Expires        string   `json:"expires"`
	CSR            string   `json:"csr"`
	NotBefore      string   `json:"notBefore"`
	NotAfter       string   `json:"notAfter"`
	Authorizations []string `json:"authorizations"`
	Certificate    string   `json:"certificate,omitempty"`
}

// An Authorization is created for each identifier in an order
type Authorization struct {
	Status     string     `json:"status"`
	Identifier Identifier `json:"identifier"`
	Challenges []string   `json:"challenges"`
	Expires    string     `json:"expires"`
}

// A Challenge is used to validate an Authorization
type Challenge struct {
	Type                     string         `json:"type"`
	URL                      string         `json:"url"`
	Token                    string         `json:"token"`
	Status                   string         `json:"status"`
	Validated                string         `json:"validated,omitempty"`
	ProvidedKeyAuthorization string         `json:"keyAuthorization,omitempty"`
	Error                    ProblemDetails `json:"error,omitempty"`
}
