package application

import (
	"fmt"
	"os"
	"time"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/domain"
	events "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/register_user"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	store        domain.AuthStore
	orchestrator *RegisterUserOrchestrator
}

func NewAuthService(store domain.AuthStore, orchestrator *RegisterUserOrchestrator) *AuthService {
	return &AuthService{
		store:        store,
		orchestrator: orchestrator,
	}
}

func (service *AuthService) Login(user *domain.UserCredential) (*domain.JWTToken, error) {
	user, err := service.store.Login(user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, status.Error(codes.NotFound, "User was not found")
	}
	token, err := service.GenerateJWT(user.Username)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (service *AuthService) Register(user *events.UserDetails) (*events.UserDetails, error) {
	var userCredential = domain.UserCredential{
		Username: user.Username,
		Password: user.Password,
	}
	registedUser, err := service.store.Register(&userCredential)
	if err != nil {
		return nil, err
	}
	user.Id = registedUser.Id.Hex()
	err = service.orchestrator.Start(*user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *AuthService) GenerateJWT(username string) (*domain.JWTToken, error) {
	var mySigningKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		err = fmt.Errorf("Something Went Wrong: %s", err.Error())
		return nil, err
	}
	return &domain.JWTToken{Token: tokenString}, nil
}

func (service *AuthService) DeleteById(id primitive.ObjectID) error {
	err := service.store.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}
