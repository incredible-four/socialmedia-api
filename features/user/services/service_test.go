package services

import (
	"errors"
	"incrediblefour/features/user"
	"incrediblefour/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("Success register", func(t *testing.T) {
		inputData := user.Core{Avatar: "", Banner: "", Name: "habib", Email: "habib@habib.com", Username: "Bekasi", Bio: "", Password: "habib123"}
		repo.On("Register", mock.Anything).Return(nil).Once()

		srv := New(repo)
		err := srv.Register(inputData)

		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Server problem", func(t *testing.T) {
		inputData := user.Core{Avatar: "", Banner: "", Name: "habib", Email: "habib@habib.com", Username: "Bekasi", Bio: "", Password: "habib123"}
		repo.On("Register", mock.Anything).Return(errors.New("There is a problem with the server")).Once()

		srv := New(repo)
		err := srv.Register(inputData)
		
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("Data already exist", func(t *testing.T) {
		inputData := user.Core{Avatar: "", Banner: "", Name: "habib", Email: "habib@habib.com", Username: "Bekasi", Bio: "", Password: "habib123"}
		repo.On("Register", mock.Anything).Return(errors.New("duplicated")).Once()

		srv := New(repo)
		err := srv.Register(inputData)
		
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "exist")
		repo.AssertExpectations(t)
	})

	t.Run("Password error", func(t *testing.T) {
		inputData := user.Core{Avatar: "", Banner: "", Name: "habib", Email: "habib@habib.com", Username: "Bekasi", Bio: "", Password: "habib123"}
		repo.On("Register", mock.Anything).Return(errors.New("Unable to process password")).Once()

		srv := New(repo)
		err := srv.Register(inputData)
		
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("Query error", func(t *testing.T) {
		inputData := user.Core{Avatar: "", Banner: "", Name: "habib", Email: "habib@habib.com", Username: "Bekasi", Bio: "", Password: "habib123"}
		repo.On("Register", mock.Anything).Return(errors.New("query")).Once()

		srv := New(repo)
		err := srv.Register(inputData)
		
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	
}