package helper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Tenantid string `json:"tenantid"`
	Token    string `json:"token"`
	Expires  string `json:"expires"`
}

const configDir = "~/.conoha"
const configFile = "config"

func GetConfigPath() string {
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Println("ホームディレクトリの取得に失敗しました")
	}
	// ホームディレクトリのパスに置き換え
	d := strings.Replace(configDir, "~", homePath, 1)
	err = os.MkdirAll(d, 0500)
	if err != nil {
		log.Fatalln("フォルダの作成に失敗しました")
	}
	// ディレクトリパスとファイル名を結合
	return filepath.Join(d, configFile)
}

func CreateOutputData(username *string, password *string, tenantid *string, token *string, expires *string) []byte {
	encryptUsername, err := TextEncrypt([]byte(*username))
	if err != nil {
		log.Fatalln(err)
	}
	encryptPassword, err := TextEncrypt([]byte(*password))
	if err != nil {
		log.Fatalln(err)
	}
	encryptTenantid, err := TextEncrypt([]byte(*tenantid))
	if err != nil {
		log.Fatalln(err)
	}
	encryptToken, err := TextEncrypt([]byte(*token))
	if err != nil {
		log.Fatalln(err)
	}
	s := fmt.Sprintf("{\"username\":\"%v\",\"password\":\"%v\",\"tenantid\":\"%v\",\"token\":\"%v\",\"expires\":\"%v\"}",
		encryptUsername,
		encryptPassword,
		encryptTenantid,
		encryptToken,
		*expires,
	)
	return []byte(s)
}

// 設定ファイルの書き込み
func WriteConfig(output []byte) {
	outPath := GetConfigPath()

	// ディレクトリパスとファイル名を結合
	f, err := os.Create(outPath)
	if err != nil {
		log.Fatalln("ファイルの作成に失敗しました")
	}
	defer f.Close()
	writer := io.Writer(f)
	_, err = writer.Write(output)
	if err != nil {
		log.Fatalln("設定ファイルの書き出しに失敗しました")
	}
}

// 設定ファイルの読み込み
func ReadConfig() *User {
	config := GetConfigPath()
	f, err := os.Open(config)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	bu := bufio.NewScanner(f)
	var d []byte
	// ファイル終端まで読み込み
	for bu.Scan() {
		d = append(d, bu.Bytes()...)
		if err == io.EOF {
			break // 終端
		}
	}
	var tmp User
	if err := json.Unmarshal(d, &tmp); err != nil {
		panic(err)
	}
	tmp.Username, err = TextDecrypt(tmp.Username)
	if err != nil {
		panic(err)
	}
	tmp.Password, err = TextDecrypt(tmp.Password)
	if err != nil {
		panic(err)
	}
	tmp.Tenantid, err = TextDecrypt(tmp.Tenantid)
	if err != nil {
		panic(err)
	}
	tmp.Token, err = TextDecrypt(tmp.Token)
	if err != nil {
		panic(err)
	}
	return &tmp
}
