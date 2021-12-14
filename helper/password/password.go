package password

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) string {
	bytesHashed, _ := bcrypt.GenerateFromPassword([]byte(pass), 10)
	passwordHashed := string(bytesHashed)
	return passwordHashed
}

func CheckSamePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
