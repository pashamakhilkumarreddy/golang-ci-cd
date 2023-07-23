package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePage(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(HomePage)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != "Home Page" {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), "Home Page")
	}
}

func TestWsEndPoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/ws", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(WsEndPoint)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != "Hola Mundo!" {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), "Hola Mundo!")
	}
}
