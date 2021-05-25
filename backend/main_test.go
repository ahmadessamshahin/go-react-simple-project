package main

import (
	dbTest "backend/src/test"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	code := m.Run()
	dbTest.ClearTable()
	os.Exit(code)
}
