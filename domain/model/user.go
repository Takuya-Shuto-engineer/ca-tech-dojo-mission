package model

// モデル
type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// ファクトリ
func CreateUser(name string) (User, error) {
	uuid, err := UUIDGenerator()
	return User{uuid, name}, err
}
