package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jeno7u/studybud/types"
	"github.com/gin-gonic/gin"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{db: map[string]*types.User{}}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func (t *testing.T) {
		userStore.db = make(map[string]*types.User)

		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName: "user",
			Email: "invalid",
			Password: "123",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d, %v", http.StatusBadRequest, rr.Code, rr.Body)
		}
	})

	t.Run("should correctly register the user", func (t *testing.T) {
		userStore.db = make(map[string]*types.User)

		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName: "user",
			Email: "valid@mail.com",
			Password: "123",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d, %v", http.StatusCreated, rr.Code, rr.Body)
		}
	})

	t.Run("should correctly register the user", func (t *testing.T) {
		userStore.db = make(map[string]*types.User)
		
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName: "user",
			Email: "valid@mail.com",
			Password: "123",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d, %v", http.StatusCreated, rr.Code, rr.Body)
		}
	})
}

type mockUserStore struct {
	db  map[string]*types.User
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
    if value, ok := m.db[email]; ok {
        return value, nil
    }
    return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserById(id int,) (*types.User, error) {
	return nil, nil
} 

func (m *mockUserStore) CreateUser(u types.User) error {
	if m.db == nil {
		m.db = make(map[string]*types.User)
	}
	// store a copy
	userCopy := u
	m.db[u.Email] = &userCopy
	return nil
}