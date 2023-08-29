package helpers

import gonanoid "github.com/matoous/go-nanoid/v2"

func GenarateToken() string {
	token, _ := gonanoid.Generate("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789", 32)
	return token
}
