package main

import (
	p "backend/src"
	dbTest "backend/src/test"
	"os"
	"testing"
)

var app = new(p.App)

func TestMain(m *testing.M) {

	code := m.Run()
	dbTest.ClearTable()
	os.Exit(code)
}
