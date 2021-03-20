// +build e2e

package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

const (
	BASE_URL = "http://localhost:8080"
)

func TestPostComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/", "author": "123456", "body": "hello world"}`).
		Post(BASE_URL + "/api/comment")

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestGetAllComments(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comment")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}

func TestGetComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		Get(BASE_URL + "/api/comment/1")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}

func TestPutComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/put", "author": "999", "body": "put world"}`).
		Put(BASE_URL + "/api/comment/1")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}

func TestDeleteComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		Delete(BASE_URL + "/api/comment/1")
	if err != nil {
		t.Fail()
	}

	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
