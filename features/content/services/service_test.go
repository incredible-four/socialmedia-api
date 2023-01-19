package services

import (
	"errors"
	"incrediblefour/features/content"
	"incrediblefour/helper"
	"incrediblefour/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	repo := mocks.NewContentData(t)

	t.Run("Success update data", func(t *testing.T) {
		inputContent := content.Core{Caption: "Happy independence day"}
		resContent := content.Core{ID: uint(1), Caption: "Happy independence day"}
		repo.On("Update", uint(1), uint(1), inputContent).Return(resContent, nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, uint(1), inputContent)

		assert.Nil(t, err)
		assert.Equal(t, resContent.ID, res.ID)
		assert.Equal(t, inputContent.Caption, res.Caption)
		repo.AssertExpectations(t)
	})

	t.Run("Invalid JWT token", func(t *testing.T) {
		inputContent := content.Core{Caption: "Happy independence day"}
		srv := New(repo)

		_, token := helper.GenerateJWT(0)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, 1, inputContent)
		
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		inputContent := content.Core{Caption: "Happy independence day"}
		repo.On("Update", uint(2), uint(2), inputContent).Return(content.Core{}, errors.New("data not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(2)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, 2, inputContent)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Server problem", func(t *testing.T) {
		inputContent := content.Core{Caption: "Happy independence day"}
		repo.On("Update", uint(1), uint(1), inputContent).Return(content.Core{}, errors.New("There is a problem with the server")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, 1, inputContent)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewContentData(t)

	t.Run("Success delete book", func(t *testing.T) {
		repo.On("Delete", uint(1), uint(1)).Return(nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)

	})

	t.Run("Invalid JWT", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(0)
		err := srv.Delete(token, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("Delete", uint(2), uint(2)).Return(errors.New("Data not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(2)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 2)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})
}

func TestContentDetail(t *testing.T) {
	repo := mocks.NewContentData(t)

	t.Run("Success show content detail", func(t *testing.T) {
		resData := content.Core{ID: uint(1), Avatar: "ava.png", Username: "habib", Image: "bajo.png", Caption: "berangkaaat"}
		repo.On("ContentDetail", uint(1)).Return(resData, nil).Once()

		srv := New(repo)
		res, err := srv.ContentDetail(uint(1))

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("ContentDetail", uint(2)).Return(content.Core{}, errors.New("Data not found")).Once()
		srv := New(repo)

		res, err := srv.ContentDetail(uint(2))

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		assert.Equal(t, res, content.Core{})
		repo.AssertExpectations(t)
	})

	t.Run("Server problem", func(t *testing.T) {
		repo.On("ContentDetail", uint(2)).Return(nil, errors.New("There is a problem with the server")).Once()
		srv := New(repo)

		res, err := srv.ContentDetail(uint(2))

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res, content.Core{})
		repo.AssertExpectations(t)
	})
}

func TestContentList(t *testing.T) {
	repo := mocks.NewContentData(t)

	t.Run("Success show content detail", func(t *testing.T) {
		resData := []content.Core{}
		repo.On("ContentList").Return(resData, nil).Once()

		srv := New(repo)
		res, err := srv.ContentList()

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("ContentList").Return([]content.Core{}, errors.New("Data not found")).Once()
		srv := New(repo)

		res, err := srv.ContentList()

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		assert.Equal(t, res, content.Core{})
		repo.AssertExpectations(t)
	})

	t.Run("Server problem", func(t *testing.T) {
		repo.On("ContentList").Return([]content.Core{}, errors.New("There is a problem with the server")).Once()
		srv := New(repo)

		res, err := srv.ContentList()

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res, []content.Core{})
		repo.AssertExpectations(t)
	})
}

func TestGetProfile(t *testing.T) {
	repo := mocks.NewContentData(t)

	t.Run("Success show a user profile", func(t *testing.T) {
		resData := content.Core{ID: uint(1), Avatar: "ava.png", Username: "habib", Image: "bajo.png", Caption: "berangkaaat"}
		repo.On("GetProfile", "habib").Return(resData, nil).Once()
		
		srv := New(repo)
		res, err := srv.ContentDetail(uint(1))

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("GetProfile", "habib").Return(nil, errors.New("Data not found")).Once()
		srv := New(repo)

		res, err := srv.GetProfile("habib")

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		assert.Equal(t, res,  nil)
		repo.AssertExpectations(t)
	})

	t.Run("Server problem", func(t *testing.T) {
		repo.On("GetProfile", "habib").Return(nil, errors.New("Unable to process the data")).Once()
		srv := New(repo)

		res, err := srv.GetProfile("habib")

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data")
		assert.Equal(t, res, nil)
		repo.AssertExpectations(t)
	})
}


