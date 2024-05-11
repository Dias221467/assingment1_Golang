package main

import (
	"errors"
	"os"
	"testing"

	"github.com/Dias221467/assingment1_Golang/internal/data"
	"github.com/Dias221467/assingment1_Golang/internal/jsonlog"
)

func (config *config) TestHandleDuplicateEmail(t *testing.T) {
	var cfg config
	cfg.db.dsn = "postgres://postgres:lbfc2005@localhost/d.ibragimovDB?sslmode=disable"
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer db.Close()

	userModel := data.UserModel{DB: db}

	user := &data.User{
		Name:      "Duplicate User",
		Email:     "2344@mail.com", // Existing email in the database
		Activated: false,
	}

	user.Password.Set("lbfc2005")

	err = userModel.Insert(user)
	if err == nil || !errors.Is(err, data.ErrDuplicateEmail) {
		t.Errorf("expected ErrDuplicateEmail, got: %v", err)
	}
}
