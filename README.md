# TodoリストAPI

## 概要
Go、Gin Webフレームワーク、Gorm ORMを使用して構築されたTodoリストAPIです。ユーザーがTodoアイテムを管理し、ユーザー認証を処理できます。

## 機能
- ユーザー登録とログイン。
- TodoアイテムのCRUD操作。

## セットアップ
1. Goがマシンにインストールされていることを確認してください。[Goの公式ページ](https://golang.org/dl/)からダウンロードできます。
2. このリポジトリをクローンします。

## 依存関係のインストール
このプロジェクトは以下のライブラリに依存しています。プロジェクトディレクトリで以下のコマンドを実行して、必要なライブラリをインストールしてください。

```bash
# Gin Webフレームワーク
go get -u github.com/gin-gonic/gin

# Gorm ORM
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

# パスワード暗号化用のbcrypt
go get -u golang.org/x/crypto/bcrypt
```
