package auth

import "golang.org/x/crypto/bcrypt"

func Hash(param string) (string,error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(param), bcrypt.DefaultCost)
	if err != nil{
		return "", err
	}
	return string(hash), nil
}