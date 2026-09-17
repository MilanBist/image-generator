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



func TestHandleRegister_CheckingInterfaceFunctions(t *testing.T){
	store := &FakeRegistrationImplementation{
		Id: 1,
		Exists: false,
	}

	token := &FakeTokenService{
		RefreshToken: "refresh-123-123",
		AccessToken: "access-123-123",
	}

	handler := &RegisterHandler{
		Store: store,
		Token: token,
	}


	body := `{
		"userName": "Milan",
		"email": "bistmilan46@gmail.com",
		"password": "Milbis123@#"
	}`


	req := httptest.NewRequest(http.MethodPost, "/api/register", strings.NewReader(body))
	res := httptest.NewRecorder()
	handler.HandleRegister(res, req)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("For the registration handler: ")
	fmt.Println(res)

	// 
	if res.Code != http.StatusAccepted{
		t.Errorf("Expected status %d got %d", http.StatusAccepted, res.Code)
	}

	// send the data to register the user
	var response models.Response
	json.NewDecoder(res.Body).Decode(&response)

	
	if response.Success != true{
		t.Fatal("Expected success = true but got false. ")
	}

	data, ok := response.Data.(map[string]any)
	if !ok{
		t.Errorf("Expected data to be of map but can't get it.")
	}



	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	if data["access"] != "access-123-123" {
		t.Fatalf("expected refresh token to be returned, got %#v", data["access"])
	}

	if data["refresh"] != "refresh-123-123"{	
		t.Fatalf("expected refresh token to be returned, got %#v", data["refresh"])	
	}
}