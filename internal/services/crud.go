package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/xarick/golang-crud/config"
	"github.com/xarick/golang-crud/internal/models"
	"github.com/xarick/golang-crud/pkg"
)

type CRUDService struct {
	cfg    *config.Application
	client *http.Client
}

func NewCRUDService(cfg *config.Application, client *http.Client) *CRUDService {
	return &CRUDService{cfg: cfg, client: client}
}

func (cs *CRUDService) CreateUser(us models.UserCrUp) (models.User, error) {
	users, err := cs.LoadUsers()
	if err != nil {
		return models.User{}, err
	}

	newID := pkg.GetUUID()

	newUser := models.User{
		ID:      newID,
		Name:    us.Name,
		Email:   us.Email,
		Address: us.Address,
	}

	users = append(users, newUser)

	err = cs.SaveUsers(users)
	if err != nil {
		return models.User{}, err
	}

	return newUser, nil
}

func (cs *CRUDService) GetUser(id string) (models.User, error) {
	users, err := cs.LoadUsers()
	if err != nil {
		return models.User{}, err
	}

	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return models.User{}, fmt.Errorf("user not found")
}

func (cs *CRUDService) GetUsers() ([]models.User, error) {
	return cs.LoadUsers()
}

func (cs *CRUDService) UpdateUser(id string, us models.UserCrUp) (models.User, error) {
	users, err := cs.LoadUsers()
	if err != nil {
		return models.User{}, err
	}

	for i, user := range users {
		if user.ID == id {
			users[i].Name = us.Name
			users[i].Email = us.Email
			users[i].Address = us.Address

			err = cs.SaveUsers(users)
			if err != nil {
				return models.User{}, err
			}

			return users[i], nil
		}
	}

	return models.User{}, fmt.Errorf("user not found")
}

func (cs *CRUDService) DeleteUser(id string) error {
	users, err := cs.LoadUsers()
	if err != nil {
		return err
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)

			err = cs.SaveUsers(users)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("user not found")
}

func (cs *CRUDService) LoadUsers() ([]models.User, error) {
	file, err := os.Open(cs.cfg.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.User{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var users []models.User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (cs *CRUDService) SaveUsers(users []models.User) error {
	file, err := os.OpenFile(cs.cfg.FileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	return encoder.Encode(users)
}
