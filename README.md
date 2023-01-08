# Conoha CLI
# 事前準備
Conohaの管理コンソールからAPIユーザーを作成する。
# CLIの初期化
ユーザ情報の設定を行う。`~/.conoha/config`が作成される。更新する場合はもう一度初期化を実行する。
```
conoha -u "user" -p "password" -t "tenant"
```
- -u ユーザ名([事前準備](#事前準備)で作成したAPIユーザ)
- -p パスワード([事前準備](#事前準備)で作成したAPIユーザ)
- -t テナントID(APIユーザの作成ページ上部にあるテナント情報)
# 使用できるAPI
## Identity API v2.0
### [バージョン情報一覧](https://www.conoha.jp/docs/identity-get_version_list.php)
バージョン情報取得
#### 使用するオプション
- -e identity
- -o get_version_list
### [バージョン情報詳細取得](https://www.conoha.jp/docs/identity-get_version_detail.php)
バージョン情報詳細取得
#### 使用するオプション
- -e identity
- -o get_version_detail
### [トークン発行](https://www.conoha.jp/docs/identity-post_tokens.php)
有効なトークン情報を取得する
#### 使用するオプション
- -e identity
- -o post_tokens