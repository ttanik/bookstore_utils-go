package rest_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RestErr interface {
	GetMessage() string
	GetStatus() int
	GetError() string
	GetCauses() []interface{}
}
type restErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restErr{Message: message, Status: status, Error: err, Causes: causes}
}
func NewError(msg string) error {
	return errors.New(msg)
}
func NewBadRequestError(message string) RestErr {
	return restErr{Message: message, Status: http.StatusBadRequest, Error: "bad_request"}
}

func NewNotFoundError(message string) RestErr {
	return restErr{Message: message, Status: http.StatusNotFound, Error: "not_found"}
}

func NewInternalServerError(message string, err error) RestErr {
	result := restErr{Message: message, Status: http.StatusInternalServerError, Error: "internal_server_error"}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}
func (e restErr) GetError() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [%v]", e.Message, e.Status, e.Error, e.Causes)
}
func (e restErr) GetMessage() string {
	return e.Message
}
func (e restErr) GetStatus() int {
	return e.Status
}

func (e restErr) GetCauses() []interface{} {
	return e.Causes
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}
