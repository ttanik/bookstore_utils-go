package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGiveNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("teste", errors.New("database error"))
	assert.Equal(t, "teste", err.GetMessage())
	assert.Equal(t, http.StatusInternalServerError, err.GetStatus())
	assert.NotNil(t, err)
	assert.Equal(t, "message: teste - status: 500 - error: internal_server_error - causes: [[database error]]", err.GetError())
	assert.Equal(t, "database error", err.GetCauses()[0])
	assert.Equal(t, 1, len(err.GetCauses()))
}

func TestShouldGiveNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("teste")
	assert.Equal(t, "teste", err.GetMessage())
	assert.Equal(t, http.StatusNotFound, err.GetStatus())
	assert.NotNil(t, err)
	assert.Equal(t, "message: teste - status: 404 - error: not_found - causes: [[]]", err.GetError())
}

func TestShouldGiveBadRequestError(t *testing.T) {
	err := NewBadRequestError("teste")
	assert.Equal(t, "teste", err.GetMessage())
	assert.Equal(t, http.StatusBadRequest, err.GetStatus())
	assert.NotNil(t, err)
	assert.Equal(t, "message: teste - status: 400 - error: bad_request - causes: [[]]", err.GetError())
}

func TestGetters(t *testing.T) {
	err := NewBadRequestError("bad request error")
	causes := err.GetCauses()
	message := err.GetMessage()
	status := err.GetStatus()
	error := err.GetError()
	assert.Equal(t, []interface{}([]interface{}(nil)), causes)
	assert.Equal(t, "bad request error", message)
	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, "message: bad request error - status: 400 - error: bad_request - causes: [[]]", error)
}
