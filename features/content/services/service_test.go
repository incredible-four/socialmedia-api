package services

import (
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
}