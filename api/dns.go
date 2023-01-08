package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ms-naughty-ghost/conoha-cli/api/schemes"
	"github.com/ms-naughty-ghost/conoha-cli/helper"
)

const (
	DomainName = "domain_name"
	DomainId   = "domain_id"
)

func (c *ApiContext) ExcuteDns(user *helper.User, req map[string]interface{}) error {
	switch c.Operation {
	case List:
		var err error
		c.Request, err = http.NewRequest(c.method, c.uri, nil)
		if err != nil {
			return err
		}
		c.Request.Header.Add("Content-Type", "application/json")
		r, err := c.Run()
		if err != nil {
			return err
		}
		defer r.Body.Close()

		body, _ := io.ReadAll(r.Body)
		if r.Status != "200" {
			return fmt.Errorf(string(body))
		}
		var tmp schemes.GetDnsV1VersionResponse
		if err := json.Unmarshal(body, &tmp); err != nil {
			panic(err)
		}
		log.Println("バージョン情報取得", tmp)
	case DomainList:

		if *req[DomainName].(*string) == "" {
			log.Fatalln("必要なパラメータがしてされていません")
		}
		var err error
		c.Request, err = http.NewRequest(c.method, c.uri, nil)
		if err != nil {
			return err
		}
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("X-Auth-Token", user.Token)
		c.SetQueryOptions(map[string]string{
			"name": *req[DomainName].(*string),
		})
		r, err := c.Run()
		if err != nil {
			return err
		}

		defer r.Body.Close()

		body, _ := io.ReadAll(r.Body)
		if r.Status != "200" {
			log.Println(c.Request)
			log.Println(r.Header)
			return fmt.Errorf(string(body))
		}
		var tmp schemes.GetDnsV1DomainListResponse
		if err := json.Unmarshal(body, &tmp); err != nil {
			panic(err)
		}
		log.Println("ドメイン一覧", tmp)
	case RecordList:

	}
	return nil
}
