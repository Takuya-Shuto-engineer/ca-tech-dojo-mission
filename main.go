package main

import (
	"ca-mission/domain/model"
	"ca-mission/interfaces/auth"
	"log"
	"net/http"
	//"sync"
	"ca-mission/infrastructure/persistence/database"

	"github.com/ant0ine/go-json-rest/rest"
)

// メモリにUsserを保存するためのstore
// var store = map[string]*model.User{}
// var lock = sync.RWMutex{}

// DBアクセスのためのレポジトリ
var userRepository = database.NewUserRepository()

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/user/create", CreateUser),
		rest.Get("/user/get", GetUser),
		rest.Put("/user/update", UpdateUser),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func CreateUser(writer rest.ResponseWriter, request *rest.Request) {
	// リクエスト受け取り用の構造体を作成
	requestContainer := model.UserCreateRequest{}

	// requestContainerにrequestで渡されたデータを代入, エラーの有無を確認
	err := request.DecodeJsonPayload(&requestContainer)
	if err != nil {
		rest.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if requestContainer.Name == "" {
		rest.Error(writer, "user name required", 400)
		return
	}

	// UUIDとnameでuserを生成
	user, err := model.CreateUser(requestContainer.Name)
	if err != nil {
		rest.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// tokenの生成
	responseContainer := model.UserCreateResponse{}
	responseContainer.Token, err = auth.CreateToken(user)
	if err != nil {
		rest.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// メモリに保存
	// lock.Lock()
	// store[user.Id] = &user
	// lock.Unlock()

	// DBに保存
	err = userRepository.InsertDB(user)
	if err != nil {
		rest.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンス
	writer.WriteHeader(http.StatusOK)
	writer.WriteJson(responseContainer)

}

func GetUser(writer rest.ResponseWriter, request *rest.Request) {
	// tokenをリクエストヘッダーから取得
	tokenString := request.Header.Get("X-Token")
	if tokenString == "" {
		rest.Error(writer, "x-token required", 400)
		return
	}

	// tokenの認証
	token, err := auth.VerifyToken(tokenString)
	if err != nil {
		rest.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// ペイロードの読み出しとidの入手
	claims := auth.ReadClaims(token)
	userId := claims["id"].(string) // interface -> string
	println(userId)

	// メモリからuserの特定
	// user := store[userId]
	// if user == nil {
	//	rest.Error(writer, "user not found", 400)
	//	return
	//}

	// DBからuserの特定
	user, err := userRepository.GetByUserIdDB(userId)
	if err != nil {
		rest.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		rest.Error(writer, "user not found", 400)
		return
	}

	// レスポンスJson用の構造体にnameを入れておく
	responseContainer := model.UserGetResponse{}
	responseContainer.Name = user.Name

	// レスポンス
	writer.WriteHeader(http.StatusOK)
	writer.WriteJson(responseContainer)
}

func UpdateUser(writer rest.ResponseWriter, request *rest.Request) {
	// tokenをリクエストヘッダーから取得
	tokenString := request.Header.Get("X-Token")
	if tokenString == "" {
		rest.Error(writer, "x-token required", 400)
		return
	}
	// tokenの認証
	token, err := auth.VerifyToken(tokenString)
	if err != nil {
		rest.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// ペイロードの読み出しとidの入手
	claims := auth.ReadClaims(token)
	userId := claims["id"].(string) // interface -> string

	// userの特定
	//user := store[userId]
	//if user == nil {
	//	rest.Error(writer, "user not found", 400)
	//	return
	//}

	// DBからuserの特定
	user, err := userRepository.GetByUserIdDB(userId)
	if err != nil {
		rest.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		rest.Error(writer, "user not found", 400)
		return
	}

	// リクエストボディ受け取り用の構造体を作成
	requestContainer := model.UserUpdateReqest{}

	// nameにrequestで渡されたデータを代入, エラーの有無を確認
	err = request.DecodeJsonPayload(&requestContainer)
	if err != nil {
		rest.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if requestContainer.Name == "" {
		rest.Error(writer, "user name required", 400)
		return
	}

	// user.Nameを更新
	updatedUser := model.User{user.Id, requestContainer.Name}

	// メモリ上で更新
	// lock.Lock()
	// store[user.Id] = user
	// lock.Unlock()

	// DB上で更新
	err = userRepository.UpdateDB(updatedUser)
	if err != nil {
		rest.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンス
	writer.WriteHeader(http.StatusOK)

}
