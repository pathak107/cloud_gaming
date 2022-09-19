package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pathak107/cloudesk/pkg/api/apierrors"
	"github.com/pathak107/cloudesk/pkg/api/dto"
	"github.com/pathak107/cloudesk/pkg/api/entity"
	"github.com/pathak107/cloudesk/pkg/api/helpers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

// jwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	jwt.StandardClaims
	Name     string `json:"name"`
	UserType string `json:"user_type"`
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func NewDefaultAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *AuthService) Login(userLogin *dto.UserLogin) (string, error) {
	var user *entity.User
	res := s.db.Where(&entity.User{Email: userLogin.Email}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return "", apierrors.NewBadRequestError(res.Error, "user_login", "no user with that email")
		}
		return "", apierrors.NewServerError(res.Error, "user_login")
	}
	if s.CheckPasswordHash(helpers.ToString(userLogin.Password), helpers.ToString(user.Password)) {
		return s.generateJWT(helpers.ToString(user.Name), helpers.ToString(user.Email), user.ID, string(user.UserType))
	} else {
		return "", apierrors.NewBadRequestError(nil, "user_login", "email or password invalid")
	}
}

func (s *AuthService) generateJWT(name string, email string, user_id uint, user_type string) (string, error) {
	// Set custom and standard claims
	claims := &JwtCustomClaims{
		jwt.StandardClaims{
			Issuer:   "pathak107",
			IssuedAt: time.Now().Unix(),
		},
		name,
		user_type,
		user_id,
		email,
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token using the secret signing key
	t, err := token.SignedString([]byte(s.getSecretKey()))
	if err != nil {
		return "", apierrors.NewServerError(err, "jwt_generation")
	}
	return t, nil

}

func (s *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(s.getSecretKey()), nil
	})
}

func (s *AuthService) getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret1428!@#"
	}
	return secret
}

func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return string(bytes), apierrors.NewServerError(err, "password_hash_generation")
	}
	return string(bytes), nil
}
