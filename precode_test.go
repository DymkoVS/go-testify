package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	//	totalCount := 4
	// req := ... // здесь нужно создать запрос к сервису
	req, err := http.NewRequest("GET", "/?count=5&city=moscow", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Check status code is OK
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body
	expectedResponse := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	if responseRecorder.Body.String() != expectedResponse {
		t.Errorf("handler returned unexpected body: got %v want %v", responseRecorder.Body.String(), expectedResponse)
	}

	// здесь нужно добавить необходимые проверки
}
