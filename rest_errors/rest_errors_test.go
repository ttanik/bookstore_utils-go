package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGiveNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("teste", errors.New("database error"))
	assert.Equal(t, "teste", err.Message)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.NotNil(t, err)
	assert.Equal(t, "internal_server_error", err.Error)
	assert.Equal(t, "database error", err.Causes[0])
	assert.Equal(t, 1, len(err.Causes))
}

func TestShouldGiveNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("teste")
	assert.Equal(t, "teste", err.Message)
	assert.Equal(t, http.StatusNotFound, err.Status)
	assert.NotNil(t, err)
	assert.Equal(t, "not_found", err.Error)
}

func TestShouldGiveBadRequestError(t *testing.T) {
	err := NewBadRequestError("teste")
	assert.Equal(t, "teste", err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.NotNil(t, err)
	assert.Equal(t, "bad_request", err.Error)
}
