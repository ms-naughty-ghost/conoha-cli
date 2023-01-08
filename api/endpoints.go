package api

import (
	"net/http"
	"net/url"
)

const (
	V1Identity = "https://identity.tyo2.conoha.io"
	V2Identity = "https://identity.tyo2.conoha.io/v2.0"
	Account    = "https://account.tyo2.conoha.io/v1"
	Compute    = "https://compute.tyo2.conoha.io/v2"
	Volume     = "https://block-storage.tyo2.conoha.io/v2"
	Database   = "https://database-hosting.tyo2.conoha.io/v1"
	Image      = "https://image-service.tyo2.conoha.io"
	Dns        = "https://dns-service.tyo2.conoha.io"
	Objects    = "https://object-storage.tyo2.conoha.io/v1"
	Network    = "https://networking.tyo2.conoha.io"
)
const (
	ArgAccount  = "account"
	ArgCompute  = "compute"
	ArgVolume   = "volume"
	ArgDatabase = "database"
	ArgImage    = "image"
	ArgDns      = "dns"
	ArgObjects  = "objects"
	ArgIdentity = "identity"
	ArgNetwork  = "network"
)

const (
	List   = "get_version_list"
	Detail = "get_version_detail"
	Tokens = "post_tokens"
)

const (
	DomainList = "get_domain_list"
	RecordList = "get_record_list"
)

func (ctx *ApiContext) SetUri(endpoint string, tenantid string, option string) error {
	var err error = nil
	ctx.Endpoints = endpoint
	ctx.Operation = option
	switch endpoint {
	case ArgAccount:
		ctx.uri, err = url.JoinPath(Account, tenantid)
	case ArgCompute:
		ctx.uri, err = url.JoinPath(Compute, tenantid)
	case ArgVolume:
		ctx.uri, err = url.JoinPath(Volume, tenantid)
	case ArgObjects:
		ctx.uri, err = url.JoinPath(Objects, tenantid)
	case ArgDatabase:
		ctx.uri = Database
	case ArgImage:
		ctx.uri = Image
	case ArgDns:
		switch option {
		case List:
			ctx.method = http.MethodGet
			ctx.uri = Dns
		case DomainList:
			ctx.method = http.MethodGet
			ctx.uri, err = url.JoinPath(Dns, "v1/domains")
		case RecordList:
			ctx.method = http.MethodGet
			ctx.uri = Dns
		}
	case ArgIdentity:
		switch option {
		case List:
			ctx.method = http.MethodGet
			ctx.uri = V1Identity
		case Detail:
			ctx.method = http.MethodGet
			ctx.uri = V2Identity
		case Tokens:
			ctx.method = http.MethodPost
			ctx.uri, err = url.JoinPath(V2Identity, "tokens")
		}
	case ArgNetwork:
		ctx.uri = Network
	default:
		ctx.uri = ""
	}
	if err != nil {
		return err
	}
	return nil
}
