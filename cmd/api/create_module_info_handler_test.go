package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Dias221467/assingment1_Golang/internal/data"
	"github.com/Dias221467/assingment1_Golang/internal/jsonlog"
	"github.com/julienschmidt/httprouter"
)

func TestCreateModuleInfoHandlerMissingField(t *testing.T) {
	// Create a new application instance

	var cfg config
	cfg.db.dsn = "postgres://postgres:lbfc2005@localhost/d.ibragimovDB?sslmode=disable"
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer db.Close()
	dbModel := &data.DBModel{
		DB: db,
	}
	app := &application{
		db: dbModel,
	}
	// Create a new router and register the createModuleInfoHandler function
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/v1/moduleinfoTEST", app.createModuleInfoHandler)

	// Create a new request with a JSON body missing the "name" field
	type payload struct {
		ModuleName     string `json:"module_name"`
		ModuleDuration int    `json:"module_duration"`
		ExamType       string `json:"exam_type"`
		Version        int    `json:"version"`
	}

	a := payload{
		ModuleName:     "This is test module",
		ModuleDuration: 6,
		ExamType:       "running",
		Version:        1,
	}

	expectedData, _ := json.Marshal(a)
	fmt.Println("Marshaled data: ", expectedData)
	req, err := http.NewRequest(http.MethodPost, "/v1/moduleinfoTEST", strings.NewReader(`{
		"module_name": "AITU2",
		"module_duration": 6,
		"exam_type": "running",
		"version": 1
	}`))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a new ResponseRecorder, which satisfies the http.ResponseWriter interface
	rr := httptest.NewRecorder()

	// Call the handler function with the request and ResponseRecorder
	router.ServeHTTP(rr, req)

	// Assert that the response has a status code of http.StatusBadRequest
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}

	// Assert that the response body contains an error message indicating the missing field
	var response map[string]string
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}
	errorMessage := response["error"]
	if errorMessage != "missing required field: name" {
		t.Errorf("Expected error message to be 'missing required field: name', got %s", errorMessage)
	}
}
