package main

import (
	"os"
	p "project/src"
	dbTest "project/src/test"
	"testing"
)

var app = new(p.App)

func TestMain(m *testing.M) {

	code := m.Run()
	dbTest.ClearTable()
	os.Exit(code)
}

func TestCustomerApis(t *testing.T) {
	dbTest.TestEmptyTable(t)
}
