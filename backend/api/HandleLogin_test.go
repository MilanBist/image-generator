package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/image-generator/internal/models"
)

func TestHandleLogin_CheckingWithInterfaceFakes(t *testing.T){
	store := &FakeUserStore{
		ID: 7,
	}

	token := &FakeTokenService{
		RefreshToken: "refresh-123-123",
		AccessToken: "access-123-1234",
	}

	handler := &LoginHandler{
		Store: store,
		Token: token,
	}

	body := `{
		"email": "testuser123@gmail.com",
		"password": "Milbis123@#"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(body))
	res := httptest.NewRecorder()
	handler.HandleLogin(res, req)

	if res.Code != http.StatusAccepted{
		t.Fatalf("expected status %d, got %d", http.StatusAccepted, res.Code)
	}

	var response models.Response
	json.NewDecoder(res.Body).Decode(&response)
	fmt.Println("The response obtained is: ", response)

	fmt.Println(response)

	if !response.Success{
		t.Fatalf("Expected success = true, got success = false message = %s", response.Message)
	}

	data, ok := response.Data.(map[string]any)
	fmt.Println(data)
	if !ok{
		t.Fatalf("expected Data to be an object, got %#v", response.Data)
	}

	if data["access"] != "access-123-1234" {
		t.Fatalf("expected refresh token to be returned, got %#v", data["access"])
	}

	if data["refresh"] != "refresh-123-123"{	
		t.Fatalf("expected refresh token to be returned, got %#v", data["refresh"])	
	}

}
