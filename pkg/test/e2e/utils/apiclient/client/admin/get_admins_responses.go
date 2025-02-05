// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// GetAdminsReader is a Reader for the GetAdmins structure.
type GetAdminsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdminsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdminsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAdminsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAdminsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAdminsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAdminsOK creates a GetAdminsOK with default headers values
func NewGetAdminsOK() *GetAdminsOK {
	return &GetAdminsOK{}
}

/*
GetAdminsOK describes a response with status code 200, with default header values.

Admin
*/
type GetAdminsOK struct {
	Payload []*models.Admin
}

// IsSuccess returns true when this get admins o k response has a 2xx status code
func (o *GetAdminsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get admins o k response has a 3xx status code
func (o *GetAdminsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admins o k response has a 4xx status code
func (o *GetAdminsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get admins o k response has a 5xx status code
func (o *GetAdminsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get admins o k response a status code equal to that given
func (o *GetAdminsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAdminsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin][%d] getAdminsOK  %+v", 200, o.Payload)
}

func (o *GetAdminsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/admin][%d] getAdminsOK  %+v", 200, o.Payload)
}

func (o *GetAdminsOK) GetPayload() []*models.Admin {
	return o.Payload
}

func (o *GetAdminsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminsUnauthorized creates a GetAdminsUnauthorized with default headers values
func NewGetAdminsUnauthorized() *GetAdminsUnauthorized {
	return &GetAdminsUnauthorized{}
}

/*
GetAdminsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetAdminsUnauthorized struct {
}

// IsSuccess returns true when this get admins unauthorized response has a 2xx status code
func (o *GetAdminsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get admins unauthorized response has a 3xx status code
func (o *GetAdminsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admins unauthorized response has a 4xx status code
func (o *GetAdminsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get admins unauthorized response has a 5xx status code
func (o *GetAdminsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get admins unauthorized response a status code equal to that given
func (o *GetAdminsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAdminsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin][%d] getAdminsUnauthorized ", 401)
}

func (o *GetAdminsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/admin][%d] getAdminsUnauthorized ", 401)
}

func (o *GetAdminsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdminsForbidden creates a GetAdminsForbidden with default headers values
func NewGetAdminsForbidden() *GetAdminsForbidden {
	return &GetAdminsForbidden{}
}

/*
GetAdminsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetAdminsForbidden struct {
}

// IsSuccess returns true when this get admins forbidden response has a 2xx status code
func (o *GetAdminsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get admins forbidden response has a 3xx status code
func (o *GetAdminsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admins forbidden response has a 4xx status code
func (o *GetAdminsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get admins forbidden response has a 5xx status code
func (o *GetAdminsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get admins forbidden response a status code equal to that given
func (o *GetAdminsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAdminsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin][%d] getAdminsForbidden ", 403)
}

func (o *GetAdminsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/admin][%d] getAdminsForbidden ", 403)
}

func (o *GetAdminsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdminsDefault creates a GetAdminsDefault with default headers values
func NewGetAdminsDefault(code int) *GetAdminsDefault {
	return &GetAdminsDefault{
		_statusCode: code,
	}
}

/*
GetAdminsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetAdminsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get admins default response
func (o *GetAdminsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get admins default response has a 2xx status code
func (o *GetAdminsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get admins default response has a 3xx status code
func (o *GetAdminsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get admins default response has a 4xx status code
func (o *GetAdminsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get admins default response has a 5xx status code
func (o *GetAdminsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get admins default response a status code equal to that given
func (o *GetAdminsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetAdminsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin][%d] getAdmins default  %+v", o._statusCode, o.Payload)
}

func (o *GetAdminsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/admin][%d] getAdmins default  %+v", o._statusCode, o.Payload)
}

func (o *GetAdminsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAdminsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
