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

## GORMの使い方
GORMでレコードを作成するときは、`db.Create`を使う。今回だとTodoControllerを使って作成する。
```go
// CreateTodoは、Todoを作成します。
func (t *TodoController) CreateTodo(todo *model.Todo) error {
    if err := t.DB.Create(todo).Error; err != nil {
        return err
    }
    return nil
}
```

レコードを取得するには、単一か全てか別れますが、今回だと全件取得なので、`db.Find`を使います。
```go
// GetTodosは、Todoを取得します。
func (t *TodoController) GetTodos() ([]model.Todo, error) {
    var todos []model.Todo
    if err := t.DB.Find(&todos).Error; err != nil {
        return nil, err
    }
    return todos, nil
}
```

レコードの更新をするときは、`db.Save`を使う。
```go
// UpdateTodoは、Todoを更新します。
func (t *TodoController) UpdateTodo(todo *model.Todo) error {
    if err := t.DB.Save(todo).Error; err != nil {
        return err
    }
    return nil
}
```

レコードの削除には、`db.Delete`を使う。削除にも種類があって、主キーを使用した削除、一括削除、論理削除などがあるらしい。

```go
// DeleteTodoは、Todoを削除します。
func (t *TodoController) DeleteTodo(todo *model.Todo) error {
    if err := t.DB.Delete(todo).Error; err != nil {
        return err
    }
    return nil
}
```