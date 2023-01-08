package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/ms-naughty-ghost/conoha-cli/api"
	"github.com/ms-naughty-ghost/conoha-cli/helper"
)

func main() {
	endpointsdesc := fmt.Sprintf("操作対象エンドポイントを指定してください\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n",
		api.ArgAccount,
		api.ArgCompute,
		api.ArgVolume,
		api.ArgDatabase,
		api.ArgImage,
		api.ArgDns,
		api.ArgObjects,
		api.ArgIdentity,
		api.ArgNetwork,
	)
	// 基本的なオプション
	var (
		username = flag.String("u", "", "ユーザ名をオプションの後に入力してください")
		password = flag.String("p", "", "ユーザパスワードをオプションの後に入力してください")
		tenantId = flag.String("t", "", "テナントIDをオプションの後に入力してください")
		endpoint = flag.String("e", "endpoint", endpointsdesc)
		option   = flag.String("o", "option", "操作内容をしてしてください")
	)
	// リクエストごとのオプション
	var (
		domainName = flag.String("domain_name", "", "ドメイン名")
		domainId   = flag.String("domain_id", "", "レコードのユニークID")
	)

	// 引数の解析
	flag.Parse()
	if *username != "" && *password != "" && *tenantId != "" {
		// ユーザーデータの保存　~/.conoha/config
		var token = ""
		var expires = ""
		helper.WriteConfig(helper.CreateOutputData(username, password, tenantId, &token, &expires))
	}
	// 既に設定ファイルが保存されていないか確認
	user := helper.ReadConfig()

	if (user.Username == "" && user.Password == "") && user.Token == "" {
		log.Println("接続先情報が指定されていません")
	}

	expires, err := time.Parse(time.RFC3339, user.Expires)
	if err != nil {
		log.Println(err)
	}
	if time.Now().After(expires) {
		// log.Println("アクセストークンの期限が切れています")
		refresh := api.New(api.ArgIdentity, user.Tenantid, api.Tokens)
		refresh.Excute(user, map[string]interface{}{})

	}
	if *endpoint == "endpoint" {
		return
	}
	c := api.New(*endpoint, user.Tenantid, *option)
	err = c.Excute(user, map[string]interface{}{
		api.DomainName: domainName,
		api.DomainId:   domainId,
	})
	if err != nil {
		log.Fatalln(err)
	}
}
