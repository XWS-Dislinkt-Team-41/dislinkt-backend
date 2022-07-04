package application

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/domain"
	events "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/register_user"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	store           domain.AuthStore
	permissionStore domain.PermissionStore
	orchestrator    *RegisterUserOrchestrator
}

func NewAuthService(store domain.AuthStore, orchestrator *RegisterUserOrchestrator, permissionStore domain.PermissionStore) *AuthService {
	return &AuthService{
		store:           store,
		orchestrator:    orchestrator,
		permissionStore: permissionStore,
	}
}

func (service *AuthService) RBAC(username string, method domain.Method, url string) (bool, error) {
	user, err := service.store.GetByUsername(username)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, status.Error(codes.NotFound, "User was not found")
	}
	fmt.Println(user.Role)
	permissions, err := service.permissionStore.GetByRole(user.Role)
	if err != nil {
		return false, err
	}
	if permissions == nil {
		return false, status.Error(codes.NotFound, "User doesn't have permissions")
	}

	if (!contains(permissions, &domain.Permission{getObjectId("0"), user.Role, method, url})) {
		return false, status.Error(codes.NotFound, "User can't access this endpoint")
	}

	return true, nil
}

func contains(permissions []*domain.Permission, permission *domain.Permission) bool {
	for _, permissionInDatabase := range permissions {
		urlsMatch, _ := regexp.MatchString(permissionInDatabase.Url, permission.Url)
		if permissionInDatabase.Role == permission.Role &&
			permissionInDatabase.Method == permission.Method &&
			urlsMatch {
			return true
		}
	}
	return false
}

func (service *AuthService) ConnectAgent(user *domain.UserCredential) (*domain.JWTToken, error) {
	user, err := service.store.Login(user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, status.Error(codes.NotFound, "User was not found")
	}
	token, err := service.GenerateAPIToken(user.Username)
	if err != nil {
		return nil, err
	}
	return token, nil
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
	fmt.Println(user.Role)
	var userCredential = domain.UserCredential{
		Username: user.Username,
		Password: user.Password,
		Role:     domain.Role(user.Role),
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
	claims["type"] = "service"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		err = fmt.Errorf("Something Went Wrong: %s", err.Error())
		return nil, err
	}
	return &domain.JWTToken{Token: tokenString}, nil
}

func (service *AuthService) GenerateAPIToken(username string) (*domain.JWTToken, error) {
	var mySigningKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["type"] = "API"
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

func (service *AuthService) DeleteById(id primitive.ObjectID) error {
	err := service.store.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}
