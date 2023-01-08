package schemes

// https://www.conoha.jp/docs/identity-get_version_list.php
type GetDnsV1VersionResponse struct {
	Versions Version `json:"versions"`
}

type GetDnsV1DomainListResponse struct {
	Domains []Domain `json:"domains"`
}

type Domain struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Email       string `json:"email"`
	Ttl         int    `json:"ttl"`
	Serial      int    `json:"serial"`
	Description string `json:"description"`
	Gslb        string `json:"gslb"`
}
