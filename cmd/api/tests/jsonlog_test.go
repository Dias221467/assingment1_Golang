package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/Dias221467/assingment1_Golang/internal/jsonlog"
)

func TestJSONLogger_PrintInfo(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := ioutil.TempFile("", "test_jsonlog_*.log")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up

	// Create a new JSONLogger instance
	logger := jsonlog.New(tmpfile, jsonlog.LevelInfo)

	// Log an INFO message
	logger.PrintInfo("This is an info message", nil)

	// Read the content of the temporary file
	content, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Check if the log message was written correctly
	if !strings.Contains(string(content), `"level":"INFO"`) {
		t.Error("Expected INFO log level, got:", string(content))
	}
}

// You can add more tests to cover other functions in the jsonlog package
