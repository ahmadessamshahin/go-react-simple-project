package test

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEmptyTable(t *testing.T) {
	ClearTable()
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	req, _ := http.NewRequest("GET", "/customers", nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, []byte("abcd"), body)
}
