package secrets

import "golang.org/x/crypto/bcrypt"

const hashCost int = 10

// HashPassword generates the password hash using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	return string(bytes), err
}

// MatchPassword checks if the password matches the hashed password
func MatchPassword(rawPassword, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(rawPassword))
	return err == nil
}
