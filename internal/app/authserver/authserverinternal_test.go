package authserver

import (
	"context"
	"testing"

	"github.com/VladimirBlinov/AuthService/internal/authservice"
	"github.com/VladimirBlinov/AuthService/internal/model"
	"github.com/VladimirBlinov/AuthService/internal/store/inmem"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_ServerCreate(t *testing.T) {
	store := inmem.New()
	logger := logrus.New()
	sm := NewSessionManager(authservice.UnimplementedAuthServiceServer{}, store, logger)

	s := model.TestInputSession(t)
	ctx := context.Background()
	sid, err := sm.Create(ctx, s)
	assert.NoError(t, err)
	assert.NotEmpty(t, sid)
}

func Test_ServerCheсk(t *testing.T) {
	store := inmem.New()
	logger := logrus.New()

	sm := NewSessionManager(authservice.UnimplementedAuthServiceServer{}, store, logger)

	s := model.TestInputSession(t)
	ctx := context.Background()
	sid, _ := sm.Create(ctx, s)

	sChecked, err := sm.Check(ctx, sid)

	assert.NoError(t, err)
	assert.Equal(t, s.UserID, sChecked.UserID)
}
