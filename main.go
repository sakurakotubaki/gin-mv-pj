package main

import (
	"fmt"
	"go-mc-app/controller"
	"go-mc-app/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
    // SQLiteデータベースに接続します
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	// AutoMigrateを呼び出してtodosテーブルを作成します
	err = db.AutoMigrate(&model.Todo{})
	if err != nil {
		panic("failed to migrate database")
	}

    // TodoControllerを初期化します
	todoController := controller.NewTodoController(db)
    // Ginのルーターを初期化します
	r := gin.Default()

    // HTTP GETでsqliteのデータを取得するエンドポイントを作成します
	r.GET("/todos", func(c *gin.Context) {
        // TodoControllerのGetTodosメソッドを呼び出して、データを取得します
		todos, err := todoController.GetTodos()
		if err != nil {
            // エラーが発生した場合は、エラーメッセージを返します
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todos)
	})
 
    // HTTP POSTでsqliteのデータを作成するエンドポイントを作成します
	r.POST("/todos", func(c *gin.Context) {
		var todo model.Todo
        // GinのShouldBindJSONメソッドを使って、リクエストボディをTodo構造体にバインドします
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
        // TodoControllerのCreateTodoメソッドを呼び出して、データを作成します
		err := todoController.CreateTodo(&todo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todo)
	})

    // HTTP PUTでsqliteのデータを更新するエンドポイントを作成します
	r.PUT("/todos/:id", func(c *gin.Context) {
		var todo model.Todo
        // GinのShouldBindJSONメソッドを使って、リクエストボディをTodo構造体にバインドします
		if err := c.ShouldBindJSON(&todo); err != nil {
            // エラーが発生した場合は、エラーメッセージを返します
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
        // TodoControllerのUpdateTodoメソッドを呼び出して、データを更新します
		err := todoController.UpdateTodo(&todo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todo)
	})

    // HTTP DELETEでsqliteのデータを削除するエンドポイントを作成します
	r.DELETE("/todos/:id", func(c *gin.Context) {
		var todo model.Todo
        // GinのShouldBindJSONメソッドを使って、リクエストボディをTodo構造体にバインドします
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
        // TodoControllerのDeleteTodoメソッドを呼び出して、データを削除します
		err := todoController.DeleteTodo(&todo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	})

    // ポート8080でサーバーを起動します
	r.Run(":8080") // http://localhost:8080/todos
}
