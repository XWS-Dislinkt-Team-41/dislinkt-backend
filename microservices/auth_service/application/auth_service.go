package application

import (
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/domain"
	jwt "github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	store domain.AuthStore
	permissionStore domain.PermissionStore
}

func NewAuthService(store domain.AuthStore, permissionStore domain.PermissionStore) *AuthService {
	return &AuthService{
		store: store,
		permissionStore: permissionStore,
	}
}

func (service *AuthService) RBAC(username string,method domain.Method,url string) (bool, error) {
	user, err := service.store.GetByUsername(username)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, status.Error(codes.NotFound, "User was not found")
	}

	permissions, err := service.permissionStore.GetByRole(user.Role)
	if err != nil {
		return false, err
	}
	if permissions == nil {
		return false, status.Error(codes.NotFound, "User doesn't have permissions")
	}
	
	if(!contains(permissions, &domain.Permission{getObjectId("0"),user.Role,method,url})){
		return false, status.Error(codes.NotFound, "User can't access this endpoint")
	}
	
	return true, nil
}

func contains(s []*domain.Permission, e *domain.Permission) bool {
    for _, a := range s {
        if a.Role == e.Role && a.Method == e.Method && a.Url == e.Url {
            return true
        }
    }
    return false
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

func (service *AuthService) Register(user *domain.UserCredential) (*domain.UserCredential, error) {
	user, err := service.store.Register(user)
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

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}