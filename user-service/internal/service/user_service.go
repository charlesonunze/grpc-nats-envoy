package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/charlesonunze/grpc-nats-envoy/user-service/internal/repo"
	"github.com/golang-jwt/jwt"
	"github.com/nats-io/nats.go"
)

// UserService - interface for the user service
type UserService interface {
	LoginUser(ctx context.Context, name string) (string, error)
	GetUserBalance(ctx context.Context, tkn string) (int32, error)
}

type userService struct {
	repo repo.UserRepo
	nc   *nats.Conn
}

// New - returns an instance of the UserService
func New(repo repo.UserRepo, nc *nats.Conn) UserService {
	return &userService{
		repo: repo,
		nc:   nc,
	}
}

var (
	mySigningKey           = []byte(os.Getenv("SECRET_KEY"))
	GET_USER_BALANCE_TOPIC = "get_user_balance"
	USER_BALANCE_TOPIC     = "user_balance"
)

func printMsg(m *nats.Msg) {
	log.Printf("Received on USR-SVC [%s]: '%s'", m.Subject, string(m.Data))
}

// LoginUser - logs in the user and returns a jwt
func (s *userService) LoginUser(ctx context.Context, name string) (string, error) {
	user, err := s.repo.GetUserByName(name)
	if err != nil {
		return "", err
	}

	if user.ID == "" {
		return "", nil
	}

	token, err := generateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserBalance - returns the user balances
func (s *userService) GetUserBalance(ctx context.Context, tkn string) (int32, error) {
	var balance int32

	// verify token
	userID, err := verifyToken(tkn)
	if err != nil {
		return balance, err
	}

	s.nc.Publish(GET_USER_BALANCE_TOPIC, []byte(userID))

	// Channel Subscriber
	ch := make(chan *nats.Msg)
	_, err = s.nc.ChanSubscribe(USER_BALANCE_TOPIC, ch)
	if err != nil {
		fmt.Printf("err%v", err)
		log.Fatal(err)
	}

	for {
		select {
		case msg := <-ch:
			data := string(msg.Data)
			fmt.Println("USR - Received message from " + msg.Subject + " channel.")
			fmt.Println("USR - payload " + data)

			i, err := strconv.ParseInt(data, 10, 32)
			if err != nil {
				panic(err)
			}

			balance = int32(i)

			return balance, nil
		}
	}
}

func verifyToken(tkn string) (string, error) {
	token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(("Invalid Signing Method"))
		}

		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			return nil, fmt.Errorf(("Expired token"))
		}

		return mySigningKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return fmt.Sprintf("%v", claims["id"]), nil
	}

	return "", fmt.Errorf(("Token verification failed"))
}

func generateJWT(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", fmt.Errorf("Something Went Wrong: %s", err.Error())
	}

	return tokenString, nil
}
