package inmem_test

import (
	"testing"

	"github.com/VladimirBlinov/AuthService/internal/model"
	"github.com/VladimirBlinov/AuthService/internal/store/inmem"
	"github.com/stretchr/testify/assert"
)

func Test_SessionsRepoCreate(t *testing.T) {
	store := inmem.New()

	s := model.TestSession(t)
	sId, err := store.Session().Create(s)

	assert.NoError(t, err)
	assert.NotEmpty(t, sId)
}
