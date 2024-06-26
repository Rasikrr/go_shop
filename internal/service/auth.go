package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go_shop/internal/models"
	"go_shop/internal/repo"
	"go_shop/internal/requests"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"unicode"
)

const (
	MIN_PASSWORD_LEN = 8
	MEDIA_DIR        = "media/profile/"
)

var (
	ErrPasswordsDoNotMatch     = errors.New("passwords do not match")
	ErrDuplicateEmailException = errors.New("user with this email is exists")
	ErrUserDoNotExists         = errors.New("user with this email does npt exists")
)

type AuthService interface {
	GetUser(email, password string) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	CreateUser(name, email, pass1, pass2 string) (string, error)
	UpdateUser(req *requests.EditProfile, photo *multipart.FileHeader, email string) error
	saveProfilePhoto(*multipart.FileHeader) (string, error)
	deleteProfilePhoto(path string) error
	generateHashPassword(string) (string, error)
	validatePassword(string) error
	GetUserById(string) (*models.User, error)
}

type AuthServiceImpl struct {
	repo repo.AuthRepo
}

func NewAuthService(repo repo.AuthRepo) *AuthServiceImpl {
	return &AuthServiceImpl{
		repo: repo,
	}
}

func (s *AuthServiceImpl) GetUser(email, password string) (*models.User, error) {
	user, err := s.repo.GetUser(email)
	if err != nil {
		log.Printf("failed to get user with email: %s | error: %v", email, err)
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		log.Printf("failed to check password %v", err)
		return nil, err
	}
	return user, nil
}

func (s *AuthServiceImpl) CreateUser(name, email, pass1, pass2 string) (string, error) {
	if pass1 != pass2 {
		return "", ErrPasswordsDoNotMatch
	}
	if err := s.validatePassword(pass1); err != nil {
		return "", err
	}
	hash, err := s.generateHashPassword(pass1)
	if err != nil {
		return "", err
	}
	user := &models.User{
		FirstName:    name,
		Email:        email,
		PasswordHash: hash,
	}
	id, err := s.repo.CreateUser(user)
	if err != nil {
		if we, ok := err.(mongo.WriteException); ok {
			for _, writeError := range we.WriteErrors {
				if writeError.Code == 11000 {
					log.Printf("Duplicate key error: %v", writeError)
					return "", ErrDuplicateEmailException
				}
			}
		}
		log.Printf("failed to create user. Email: %s. Error: %v", user.Email, err)
		return "", err
	}
	return id.Hex(), nil
}

func (s *AuthServiceImpl) generateHashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to generate hash for password %v", err)
		return "", err
	}
	return string(hash), nil
}

func (s *AuthServiceImpl) validatePassword(pass string) error {
	if len(pass) < MIN_PASSWORD_LEN {
		return fmt.Errorf("password minimum length is %d", MIN_PASSWORD_LEN)
	}
	var (
		upper      bool
		digit      bool
		specSymbol bool
	)
	for _, el := range pass {
		switch {
		case unicode.IsUpper(el):
			upper = true
		case unicode.IsDigit(el):
			digit = true
		case unicode.IsSymbol(el) || unicode.IsPunct(el):
			specSymbol = true
		}
		if upper && digit && specSymbol {

			return nil

		}
	}
	return fmt.Errorf("password must contain atleast one upper, one specSymbol and one digit")
}

func (s *AuthServiceImpl) GetUserById(id string) (*models.User, error) {
	user, err := s.repo.GetUserById(id)
	if err != nil {
		log.Printf("failed to get user by id %s | %v", id, err)
		return nil, err
	}
	return user, nil
}

func (s *AuthServiceImpl) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.repo.GetUser(email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, err
		}
		return nil, err
	}
	return user, nil
}

func (s *AuthServiceImpl) UpdateUser(req *requests.EditProfile, photo *multipart.FileHeader, email string) error {
	if photo != nil {
		user, err := s.GetUserByEmail(email)
		if err != nil {
			log.Printf("failed to get user | %v", err)
			return err
		}
		if user.PhotoPath != "" {
			if err := s.deleteProfilePhoto(MEDIA_DIR + user.PhotoPath); err != nil {
				log.Printf("failed to delete previous profile photo | %v", err)
				return err
			}
		}
		fileName, err := s.saveProfilePhoto(photo)
		if err != nil {
			return err
		}
		req.PhotoPath = fileName
	}
	data, err := bson.Marshal(req)
	if err != nil {
		log.Printf("failed to prepare data for update | %v", err)
		return err
	}

	var updateBSON bson.D
	if err := bson.Unmarshal(data, &updateBSON); err != nil {
		log.Printf("failed to prepare data for update | %v", err)
		return err
	}

	if err := s.repo.UpdateUser(updateBSON, email); err != nil {
		log.Printf("failed to update user | email: %s", email)
		return err
	}
	return nil
}

func (s *AuthServiceImpl) saveProfilePhoto(fileFromForm *multipart.FileHeader) (string, error) {
	fileName := uuid.New().String() + ".jpg"
	filePath := MEDIA_DIR + fileName
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("failed to create file | %v", err)
		return "", err
	}

	defer file.Close()
	f, err := fileFromForm.Open()
	if err != nil {
		log.Printf("failed to open file from form | %v", err)
		return "", err
	}

	defer f.Close()

	if _, err := io.Copy(file, f); err != nil {
		log.Printf("failed to write file | %v", err)
		return "", err
	}
	return fileName, nil
}

func (s *AuthServiceImpl) deleteProfilePhoto(path string) error {
	err := os.Remove(path)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			log.Println("file does not exists")
		}
		return err
	}
	return nil
}
