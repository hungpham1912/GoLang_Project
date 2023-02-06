package service

import (
	userRepositories "CRUD/src/repositories"
	"CRUD/src/system/module/controller/auth/dto"
	CreateUserDto "CRUD/src/system/module/controller/users/dto"
	"fmt"

	userModel "CRUD/src/system/module/model/users"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
)

func Register(registerUser dto.RegisterDto) userModel.ResponseRegister {

	newPassword := HashPassword(registerUser.PassWord)
	var userDto CreateUserDto.UserDto
	userDto.PassWord = newPassword
	userDto.Email = registerUser.Email
	userDto.Name = registerUser.Name

	fmt.Println("newPassword: ", newPassword)

	var result userModel.ResponseRegister

	ds := userRepositories.FindByEmail(registerUser.Email)
	if len(ds) > 0 {
		result.Status = false
		result.Message = "Email already exist"
		result.Data = userModel.User{}
		result.AccessToken = ""
		return result
	}

	user := userRepositories.Save(userDto)

	result.Status = true
	result.Message = "OK"
	result.Data = user
	result.AccessToken = "asdadasd"

	return result

}

func HashPassword(passWord string) string {
	key := os.Getenv("HASH_KEY")
	hashPassword := encrypt([]byte(passWord), key)
	return hex.EncodeToString(hashPassword)
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func encryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}

func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}
