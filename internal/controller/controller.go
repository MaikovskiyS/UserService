package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"practice/internal/models"
	"practice/internal/utils"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type service interface {
	GetFriendsById(id uint) ([]string, error)
	AddFriend(ctx context.Context, f models.Friendship) (name1, name2 string, err error)
	DeleteUser(ctx context.Context, id *uint) error
	GetUserById(ctx context.Context, id *uint) (*models.User, error)
	CreateUser(ctx context.Context, u *models.User) (uint, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
}
type Controller struct {
	log     *logrus.Logger
	service service
	chi     *chi.Mux
}

//fabric funcion
func NewController(logger *logrus.Logger, svc service) *Controller {
	c := chi.NewRouter()
	return &Controller{
		log:     logger,
		service: svc,
		chi:     c,
	}
}
func (c *Controller) CreateEndPoints() *Controller {
	c.chi.Post("/getuser", c.GetUserByIdEndPoint)
	c.chi.Post("/create", c.CreateUserEndPoint)
	c.chi.Get("/users", c.GetAllUsersEndPoint)
	c.chi.Delete("/deleteuser", c.DeleteUserEndPoint)
	c.chi.Post("/addfriend", c.AddFriendEndPoint)
	c.chi.Post("/allfriends", c.GetFriendsByIdEndPoint)
	return c
}
func (c *Controller) Start() error {
	return http.ListenAndServe("localhost:8080", c.chi)
}

//POST create user
func (c *Controller) CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	c.log.Info("create context")
	user := &models.User{}
	c.log.Info("create user struct")
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		c.log.Info(err)
	}
	c.log.Info("decode json request")

	defer r.Body.Close()
	id, err := c.service.CreateUser(ctx, user)
	if err != nil {
		c.log.Info(err)
	}
	text := "Пользователь добавлен, Ваш ID:"
	response := text + strconv.Itoa(int(id))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(response))

}

//GET get all users
func (c *Controller) GetAllUsersEndPoint(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	data, err := c.service.GetAllUsers(ctx)
	if err != nil {
		c.log.Info(err)
	}
	response, err := json.Marshal(data)
	if err != nil {
		c.log.Info(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

//POST get user by id
func (c *Controller) GetUserByIdEndPoint(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadJson(*r)
	if err != nil {
		c.log.Info(err)
	}
	ctx := context.Background()

	user, err := c.service.GetUserById(ctx, &id)
	if err != nil {
		c.log.Info(err)
	}
	response, err := json.Marshal(user)
	if err != nil {
		c.log.Info(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

//DELETE delete user
func (c *Controller) DeleteUserEndPoint(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id, err := utils.ReadJson(*r)
	if err != nil {
		c.log.Info("cant read json")
	}
	c.service.DeleteUser(ctx, &id)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Пользовалель удален"))
}

//add friend
func (c *Controller) AddFriendEndPoint(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var newFriendship models.Friendship
	err := json.NewDecoder(r.Body).Decode(&newFriendship)
	if err != nil {
		c.log.Info("cant decode request", err)
	}
	fmt.Println(newFriendship)
	name1, name2, _ := c.service.AddFriend(ctx, newFriendship)
	fmt.Println(name1, name2)
	//TODO
	response := name1 + " и " + name2 + " теперь Друзья!"
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(response))
}
func (c *Controller) GetFriendsByIdEndPoint(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	text := string(data)
	id, err := strconv.ParseUint(text, 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	arr, _ := c.service.GetFriendsById(uint(id))
	str := strings.Join(arr, " ")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(str))
}
