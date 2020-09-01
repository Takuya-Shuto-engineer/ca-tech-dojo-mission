package model

import "github.com/google/uuid"

// モデルで使うドメインには関係ない関数とか
func UUIDGenerator() (string, error) {
	// UUIDの生成
	uuidObj, err := uuid.NewRandom()
	uuid := uuidObj.String()
	return uuid, err
}
