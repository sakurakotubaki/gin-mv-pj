# MCでGo言語を使ってみる
MCとは、Mがモデルで、Cがコントローラーです。

レイヤーはこのように分けています。`main.go`でプログラムを実行します。

```
.
├── README.md
├── controller
│   └── todo_controller.go
├── go.mod
├── go.sum
├── main.go
├── model
│   └── todo_model.go
└── test.db
```

1. Go言語を環境構築する
```bash
go mod init go-mc-app
```

2. Ginを追加する
```bash
go get -u github.com/gin-gonic/gin
```

3. gormを追加する
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```