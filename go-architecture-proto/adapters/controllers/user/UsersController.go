package controllers

import (
	"fmt"
	"net/http"

	"go-architecture-proto/entities/models"
	repository "go-architecture-proto/entities/repositories/user"
	usecase "go-architecture-proto/usecases/user"
)

// GET:users
// POST:users
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	// 本番用リポジトリ層
	repository := repository.UserRepository{}
	// サービス層作成
	service := usecase.NewUserService(&repository)

	requestMethod := r.Method // これでhttpリクエストのメソッドを取得

	// GET
	if requestMethod == "GET" {
		users = usecase.ShowUsers()
		var output = ""
		for _, user := range users {
			output += fmt.Sprintf("id = %d name = %s ", user.Id, user.Name)
		}

		// ステータスコードを設定
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, output)
		return
	}

	// POST
	if requestMethod == "POST" {
		usecase.CreateUser()
		w.WriteHeader(http.StatusCreated)
		return
	}

	// それ以外のメソッドはエラー
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}
