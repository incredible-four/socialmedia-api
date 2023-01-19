package services

import (
	"errors"
	"incrediblefour/features/user"
	"incrediblefour/helper"
	"incrediblefour/mocks"
	"mime/multipart"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("Success register", func(t *testing.T) {
		inputData := user.Core{Name: "habib", Email: "habib@habib.com", Username: "Bekasi", Password: "habib123"}
		repo.On("Register", mock.Anything).Return(nil).Once()

		srv := New(repo)
		err := srv.Register(inputData)

		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Server problem", func(t *testing.T) {
		inputData := user.Core{Name: "habib", Email: "habib@habib.com", Username: "Bekasi", Password: "habib123"}
		repo.On("Register", mock.Anything).Return(errors.New("There is a problem with the server")).Once()

		srv := New(repo)
		err := srv.Register(inputData)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("Data already exist", func(t *testing.T) {
		inputData := user.Core{Name: "habib", Email: "habib@habib.com", Username: "Bekasi", Password: "habib123"}
		repo.On("Register", mock.Anything).Return(errors.New("duplicated")).Once()

		srv := New(repo)
		err := srv.Register(inputData)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "exist")
		repo.AssertExpectations(t)
	})

	t.Run("Password error", func(t *testing.T) {
		inputData := user.Core{Name: "habib", Email: "habib@habib.com", Username: "Bekasi", Password: "habib123"}
		repo.On("Register", mock.Anything).Return(errors.New("Unable to process password")).Once()

		srv := New(repo)
		err := srv.Register(inputData)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("Query error", func(t *testing.T) {
		inputData := user.Core{Name: "habib", Email: "habib@habib.com", Username: "Bekasi", Password: "habib123"}
		repo.On("Register", mock.Anything).Return(errors.New("query")).Once()

		srv := New(repo)
		err := srv.Register(inputData)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("Login success", func(t *testing.T) {
		inputUsername := "habib"
		hashed, _ := helper.GeneratePassword("habib123")
		resData := user.Core{ID: uint(1), Name: "Muhammad Habibullah", Email: "habib@habib.com", Username: "habib", Password: hashed}

		repo.On("Login", inputUsername).Return(resData, nil).Once()

		srv := New(repo)
		token, res, err := srv.Login(inputUsername, "habib123")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Server problem", func(t *testing.T) {
		inputUsername := "habib"
		repo.On("Login", inputUsername).Return(user.Core{}, errors.New("There is a problem with the server")).Once()

		srv := New(repo)
		token, res, err := srv.Login(inputUsername, "habib123")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Empty(t, token)
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		inputUsername := "habibun"
		repo.On("Login", inputUsername).Return(user.Core{}, errors.New("Data not found")).Once()

		srv := New(repo)
		token, res, err := srv.Login(inputUsername, "habibun123")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "Data not found")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Email or Password doesn't match", func(t *testing.T) {
		inputUsername := "habib"
		hashed, _ := helper.GeneratePassword("habib1234")
		resData := user.Core{ID: uint(1), Name: "Muhammad Habibullah", Email: "habib@habib.com", Username: "habib", Password: hashed}
		repo.On("Login", inputUsername).Return(resData, nil).Once()

		srv := New(repo)
		token, res, err := srv.Login(inputUsername, "habib123")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "match")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}

func TestProfile(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("Success show profile", func(t *testing.T) {
		resData := user.Core{ID: uint(1), Name: "Muhammad Habibullah", Email: "habib@habib.com", Username: "habib", Avatar: "ava.png", Banner: "Banner.png", Bio: "Hi stalker!"}

		repo.On("Profile", uint(1)).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Profile(pToken)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Profile", uint(1)).Return(user.Core{}, errors.New("data not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Server problem", func(t *testing.T) {
		repo.On("Profile", mock.Anything).Return(user.Core{}, errors.New("There is a problem with the server")).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}

// func getFileHeader(file *os.File) (*multipart.FileHeader, *multipart.FileHeader, error) {
// 	// get file size
// 	fileStat, err := file.Stat()
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	// create *multipart.FileHeader
// 	return &multipart.FileHeader{
// 		Filename: fileStat.Name(),
// 		Size:     fileStat.Size(),
// 	}, nil, nil
// }

func TestUpdate(t *testing.T) {
	repo := mocks.NewUserData(t)
	hashed, _ := helper.GeneratePassword("alif123")
	var a, b multipart.FileHeader
	t.Run("success update profile", func(t *testing.T) {
		inputData := user.Core{ID: uint(1), Avatar: "https://res.cloudinary.com/dg5psgujz/image/upload/v1674109694/incredible4/wkoznm8ujw5l66vyqg2i.jpg", Banner: "https://res.cloudinary.com/dg5psgujz/image/upload/v1674109696/incredible4/hk5vhzorzxagspio0hue.png", Bio: "BE 14", Name: "Alif Muhamad Hafidz", Email: "alif@alif.com", Username: "alif", Password: hashed}
		resData := user.Core{ID: uint(1), Avatar: "https://res.cloudinary.com/dg5psgujz/image/upload/v1674109694/incredible4/wkoznm8ujw5l66vyqg2i.jpg", Banner: "https://res.cloudinary.com/dg5psgujz/image/upload/v1674109696/incredible4/hk5vhzorzxagspio0hue.png", Bio: "BE 14", Name: "Alif Muhamad Hafidz", Email: "alif@alif.com", Username: "alif"}
		repo.On("Update", uint(1), inputData).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		inputData.Password = hashed
		res, err := srv.Update(a, b, pToken, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Name, res.Name)
		assert.Equal(t, resData.Email, res.Email)
		assert.Equal(t, resData.Bio, res.Bio)
		repo.AssertExpectations(t)
	})
}

func TestDeactivate(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("success deactivate", func(t *testing.T) {
		repo.On("Deactivate", uint(1)).Return(nil).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Deactivate(pToken)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("Deactivate", uint(2)).Return(errors.New("data not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(2)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Deactivate(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})

	t.Run("Invalid JWT", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Deactivate(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})

	t.Run("Server problem", func(t *testing.T) {
		repo.On("Deactivate", uint(2)).Return(errors.New("There is a problem with the server")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(2)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Deactivate(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

}
