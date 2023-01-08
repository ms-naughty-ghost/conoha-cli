package schemes

// https://www.conoha.jp/docs/identity-get_version_list.php
type GetIdentityV1Response struct {
	Versions Version `json:"versions"`
}

// https://www.conoha.jp/docs/identity-get_version_detail.php
type GetIdentityV2Response struct {
	Version Values `json:"version"`
}

// https://www.conoha.jp/docs/identity-post_tokens.php
type PostIdentityV2TokensRequest struct {
	Auth Auth `json:"auth"`
}

// https://www.conoha.jp/docs/identity-post_tokens.php
type PostIdentityV2TokensResponse struct {
	Access Access `json:"access"`
}

type Version struct {
	Values []Values `json:"values"`
}
type Values struct {
	Id         string      `json:"id"`
	Links      []Links     `json:"links"`
	MediaTypes []MediaType `json:"media-types"`
	Status     string      `json:"status"`
	Updated    string      `json:"updated"`
}
type Links struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	Type string `json:"type"`
}
type MediaType struct {
	Base string `json:"base"`
}
type Access struct {
	Token          Token            `json:"token"`
	ServiceCatalog []ServiceCatalog `json:"serviceCatalog"`
	User           User             `json:"user"`
	Metadata       Metadata         `json:"metadata"`
}
type Tenant struct {
	Id             string   `json:"id"`
	Name           string   `json:"name"`
	Enabled        bool     `json:"enabled"`
	Description    string   `json:"description"`
	DomainId       string   `json:"domain_id"`
	Tyo1ImageSize  string   `json:"tyo1_image_size"`
	EndpointsLinks []string `json:"endpoints_links"`
	Type           string   `json:"type"`
}
type Token struct {
	Id       string   `json:"id"`
	Expires  string   `json:"expires"`
	IssuedAt string   `json:"issued_at"`
	Tenant   Tenant   `json:"tenant"`
	AuditIds []string `json:"audit_ids"`
}
type Endpoint struct {
	Region    string `json:"region"`
	PublicUrl string `json:"publicURL"`
}
type ServiceCatalog struct {
	Endpoints     []Endpoint `json:"endpoints"`
	EndpointLinks []string   `json:"endpoints_links"`
	Type          string     `json:"type"`
	Name          string     `json:"name"`
}
type Role struct {
	Name string `json:"name"`
}
type User struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Username   string   `json:"username"`
	RolesLinks []string `json:"roles_links"`
	Roles      []Role   `json:"roles"`
}
type Metadata struct {
	IsAdmin int      `json:"is_admin"`
	Roles   []string `json:"roles"`
}
type Auth struct {
	PasswordCredentials PasswordCredentials `json:"passwordCredentials"`
	TenantId            string              `json:"tenantId"`
}
type PasswordCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
