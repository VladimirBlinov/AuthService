package authserver

import (
	"testing"

	"github.com/VladimirBlinov/AuthService/internal/store/inmem"
)

func Test_SessionCreate(t *testing.T) {
	store := inmem.New()
	sm := SessionManager.New()
}
