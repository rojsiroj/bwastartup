package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, error := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if error != nil {
		return user, error
	}
	user.PasswordHash = string(passwordHash)
	newUser, error := s.repository.Save(user)
	if error != nil {
		return newUser, error
	}

	return newUser, nil
}
