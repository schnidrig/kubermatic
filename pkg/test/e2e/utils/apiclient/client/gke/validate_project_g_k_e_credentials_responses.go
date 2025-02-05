// Code generated by go-swagger; DO NOT EDIT.

package gke

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ValidateProjectGKECredentialsReader is a Reader for the ValidateProjectGKECredentials structure.
type ValidateProjectGKECredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateProjectGKECredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateProjectGKECredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewValidateProjectGKECredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewValidateProjectGKECredentialsOK creates a ValidateProjectGKECredentialsOK with default headers values
func NewValidateProjectGKECredentialsOK() *ValidateProjectGKECredentialsOK {
	return &ValidateProjectGKECredentialsOK{}
}

/*
ValidateProjectGKECredentialsOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type ValidateProjectGKECredentialsOK struct {
}

// IsSuccess returns true when this validate project g k e credentials o k response has a 2xx status code
func (o *ValidateProjectGKECredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate project g k e credentials o k response has a 3xx status code
func (o *ValidateProjectGKECredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate project g k e credentials o k response has a 4xx status code
func (o *ValidateProjectGKECredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate project g k e credentials o k response has a 5xx status code
func (o *ValidateProjectGKECredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate project g k e credentials o k response a status code equal to that given
func (o *ValidateProjectGKECredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ValidateProjectGKECredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/validatecredentials][%d] validateProjectGKECredentialsOK ", 200)
}

func (o *ValidateProjectGKECredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/validatecredentials][%d] validateProjectGKECredentialsOK ", 200)
}

func (o *ValidateProjectGKECredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateProjectGKECredentialsDefault creates a ValidateProjectGKECredentialsDefault with default headers values
func NewValidateProjectGKECredentialsDefault(code int) *ValidateProjectGKECredentialsDefault {
	return &ValidateProjectGKECredentialsDefault{
		_statusCode: code,
	}
}

/*
ValidateProjectGKECredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ValidateProjectGKECredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the validate project g k e credentials default response
func (o *ValidateProjectGKECredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this validate project g k e credentials default response has a 2xx status code
func (o *ValidateProjectGKECredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this validate project g k e credentials default response has a 3xx status code
func (o *ValidateProjectGKECredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this validate project g k e credentials default response has a 4xx status code
func (o *ValidateProjectGKECredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this validate project g k e credentials default response has a 5xx status code
func (o *ValidateProjectGKECredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this validate project g k e credentials default response a status code equal to that given
func (o *ValidateProjectGKECredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ValidateProjectGKECredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/validatecredentials][%d] validateProjectGKECredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ValidateProjectGKECredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/validatecredentials][%d] validateProjectGKECredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ValidateProjectGKECredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ValidateProjectGKECredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
