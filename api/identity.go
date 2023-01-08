package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ms-naughty-ghost/conoha-cli/api/schemes"
	"github.com/ms-naughty-ghost/conoha-cli/helper"
)

func (c *ApiContext) ExcuteIdentity(user *helper.User) error {
	switch c.Operation {
	case List:
		var err error
		c.Request, err = http.NewRequest(c.method, c.uri, nil)
		if err != nil {
			return err
		}
		r, err := c.Run()
		if err != nil {
			return err
		}
		defer r.Body.Close()

		body, _ := io.ReadAll(r.Body)
		if r.Status != "200" {
			return fmt.Errorf(string(body))
		}
		var tmp schemes.GetIdentityV1Response
		if err := json.Unmarshal(body, &tmp); err != nil {
			panic(err)
		}
		log.Println("バージョン情報取得", tmp)
	case Detail:
		var err error
		c.Request, err = http.NewRequest(c.method, c.uri, nil)
		if err != nil {
			return err
		}
		r, err := c.Run()
		if err != nil {
			return err
		}
		defer r.Body.Close()

		body, _ := io.ReadAll(r.Body)
		if r.Status != "200" {
			return fmt.Errorf(string(body))
		}
		var tmp schemes.GetIdentityV2Response
		if err := json.Unmarshal(body, &tmp); err != nil {
			panic(err)
		}
		log.Println("バージョン情報詳細取得", tmp)
	case Tokens:
		reqBody := schemes.PostIdentityV2TokensRequest{
			Auth: schemes.Auth{
				PasswordCredentials: schemes.PasswordCredentials{
					Username: user.Username,
					Password: user.Password,
				},
				TenantId: user.Tenantid,
			},
		}
		JsonString, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		c.Request, err = http.NewRequest(c.method, c.uri, bytes.NewBuffer(JsonString))
		if err != nil {
			return err
		}
		r, err := c.Run()
		if err != nil {
			return err
		}
		defer r.Body.Close()

		body, _ := io.ReadAll(r.Body)
		if r.Status != "200" {
			return fmt.Errorf(string(body))
		}
		var tmp schemes.PostIdentityV2TokensResponse
		if err := json.Unmarshal(body, &tmp); err != nil {
			panic(err)
		}
		c.Username = tmp.Access.User.Username
		c.Password = user.Password
		c.Tenant = user.Tenantid
		c.Token = tmp.Access.Token.Id
		helper.WriteConfig(helper.CreateOutputData(&c.Username, &c.Password, &c.Tenant, &c.Token, &tmp.Access.Token.Expires))
		log.Println("アクセストークン", tmp.Access.Token.Id)
	default:
		return fmt.Errorf("操作内容が指定されていません")
	}
	return nil
}
