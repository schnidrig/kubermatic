// Code generated by go-swagger; DO NOT EDIT.

package anexia

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListAnexiaVlansReader is a Reader for the ListAnexiaVlans structure.
type ListAnexiaVlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAnexiaVlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAnexiaVlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAnexiaVlansDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAnexiaVlansOK creates a ListAnexiaVlansOK with default headers values
func NewListAnexiaVlansOK() *ListAnexiaVlansOK {
	return &ListAnexiaVlansOK{}
}

/*
ListAnexiaVlansOK describes a response with status code 200, with default header values.

AnexiaVlanList
*/
type ListAnexiaVlansOK struct {
	Payload models.AnexiaVlanList
}

// IsSuccess returns true when this list anexia vlans o k response has a 2xx status code
func (o *ListAnexiaVlansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list anexia vlans o k response has a 3xx status code
func (o *ListAnexiaVlansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list anexia vlans o k response has a 4xx status code
func (o *ListAnexiaVlansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list anexia vlans o k response has a 5xx status code
func (o *ListAnexiaVlansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list anexia vlans o k response a status code equal to that given
func (o *ListAnexiaVlansOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAnexiaVlansOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/anexia/vlans][%d] listAnexiaVlansOK  %+v", 200, o.Payload)
}

func (o *ListAnexiaVlansOK) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/anexia/vlans][%d] listAnexiaVlansOK  %+v", 200, o.Payload)
}

func (o *ListAnexiaVlansOK) GetPayload() models.AnexiaVlanList {
	return o.Payload
}

func (o *ListAnexiaVlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAnexiaVlansDefault creates a ListAnexiaVlansDefault with default headers values
func NewListAnexiaVlansDefault(code int) *ListAnexiaVlansDefault {
	return &ListAnexiaVlansDefault{
		_statusCode: code,
	}
}

/*
ListAnexiaVlansDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAnexiaVlansDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list anexia vlans default response
func (o *ListAnexiaVlansDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list anexia vlans default response has a 2xx status code
func (o *ListAnexiaVlansDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list anexia vlans default response has a 3xx status code
func (o *ListAnexiaVlansDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list anexia vlans default response has a 4xx status code
func (o *ListAnexiaVlansDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list anexia vlans default response has a 5xx status code
func (o *ListAnexiaVlansDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list anexia vlans default response a status code equal to that given
func (o *ListAnexiaVlansDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListAnexiaVlansDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/anexia/vlans][%d] listAnexiaVlans default  %+v", o._statusCode, o.Payload)
}

func (o *ListAnexiaVlansDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/anexia/vlans][%d] listAnexiaVlans default  %+v", o._statusCode, o.Payload)
}

func (o *ListAnexiaVlansDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAnexiaVlansDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
