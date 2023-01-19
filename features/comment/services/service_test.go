package services

import (
	"incrediblefour/features/comment"
	"incrediblefour/helper"
	"incrediblefour/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewCommentData(t)

	t.Run("Success add comment", func(t *testing.T) {
		inputData := comment.Core{Text: "bang giveaway bang"}
		repo.On("Add", mock.Anything).Return(nil).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		srv := New(repo)
		res, err := srv.Add(pToken, inputData, uint(1))

		assert.Nil(t, err)
		assert.Equal(t, inputData.Text, res.Text)
		assert.Equal(t, inputData.ID, res.ID)
		repo.AssertExpectations(t)
	})
}