package authserver

import (
	"context"
	"testing"

	"github.com/VladimirBlinov/AuthService/internal/model"
	"github.com/VladimirBlinov/AuthService/internal/store/inmem"
	"github.com/VladimirBlinov/AuthService/pkg/authservice"
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

	testCases := []struct {
		name    string
		sid     *authservice.SessionID
		IsValid bool
	}{
		{
			name:    "checked",
			sid:     sid,
			IsValid: true,
		},
		{
			name: "not found",
			sid: &authservice.SessionID{
				ID: "new",
			},
			IsValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sChecked, err := sm.Check(ctx, tc.sid)
			if tc.IsValid {
				assert.NoError(t, err)
				assert.Equal(t, s.UserID, sChecked.UserID)
			} else {
				assert.Error(t, err)
				assert.Nil(t, sChecked)
			}
		})
	}
}

func Test_ServerDelete(t *testing.T) {
	store := inmem.New()
	logger := logrus.New()

	sm := NewSessionManager(authservice.UnimplementedAuthServiceServer{}, store, logger)
	s := model.TestInputSession(t)
	ctx := context.Background()
	sid, _ := sm.Create(ctx, s)

	testCases := []struct {
		name      string
		sid       *authservice.SessionID
		IsDeleted bool
	}{
		{
			name:      "deleted",
			sid:       sid,
			IsDeleted: true,
		},
		{
			name: "not found",
			sid: &authservice.SessionID{
				ID: "new",
			},
			IsDeleted: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			deleted, err := sm.Delete(ctx, tc.sid)
			if tc.IsDeleted {
				assert.NoError(t, err)
				assert.True(t, deleted.Dummy)

				sChecked, err := sm.Check(ctx, sid)
				assert.Error(t, err)
				assert.Nil(t, sChecked)
			} else {
				assert.Error(t, err)
				assert.False(t, deleted.Dummy)
			}
		})
	}

}
