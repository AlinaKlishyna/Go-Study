package auth

import (
	"errors"
	"time"

	"github.com/AlinaKlishyna/Go-Study.git/jwt/core"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const signingKey = "SUPER_SECRET_KEY"
const tokenTTL = 12 * time.Hour

type Authorization struct {
	domain *core.Domain
}

type TokenClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}

func NewAuthorization(domain *core.Domain) *Authorization {
	return &Authorization{domain: domain}
}

// ---------------- PASSWORD UTILS ---------------- //

func hashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPassword(pass, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)) == nil
}

// ---------------- JWT ---------------- //

func (a *Authorization) Register(username, password string) (*core.User, error) {
	hashed, _ := hashPassword(password)
	user := &core.User{
		Username: username,
		Password: hashed,
	}
	return a.domain.CreateUser(user)
}

func (a *Authorization) GenerateToken(username, password string) (string, error) {
	user, err := a.domain.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !checkPassword(password, user.Password) {
		return "", errors.New("wrong password")
	}

	claims := TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}
