package services

import (
	"errors"
	"incrediblefour/features/comment"
	"incrediblefour/helper"
	"incrediblefour/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewCommentData(t)

	t.Run("Success add comment", func(t *testing.T) {
		inputData := comment.Core{Text: "bang giveaway bang"}
		resData := comment.Core{ID: uint(1), Text: "bang giveaway bang"}
		repo.On("Add", uint(1), inputData, uint(1)).Return(resData, nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, inputData, uint(1))
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Text, res.Text)
		repo.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		inputData := comment.Core{Text: "bang giveaway bang"}
		repo.On("Add", uint(1), inputData, uint(1)).Return(comment.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, inputData, uint(1))
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("user tidak ditemukan", func(t *testing.T) {
		inputData := comment.Core{Text: "bang giveaway bang"}
		repo.On("Add", uint(1), inputData, uint(1)).Return(comment.Core{}, errors.New("not found")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, inputData, uint(1))
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		inputData := comment.Core{Text: "bang giveaway bang"}
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		res, err := srv.Add(token, inputData, uint(1))
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "not found")
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewCommentData(t)

	t.Run("suskes hapus comment", func(t *testing.T) {
		repo.On("Delete", uint(1), uint(1)).Return(nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)

	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(0)
		err := srv.Delete(token, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		repo.On("Delete", uint(2), uint(2)).Return(errors.New("data not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(2)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 2)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})
}
