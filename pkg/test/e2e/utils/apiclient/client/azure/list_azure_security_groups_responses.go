// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListAzureSecurityGroupsReader is a Reader for the ListAzureSecurityGroups structure.
type ListAzureSecurityGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAzureSecurityGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAzureSecurityGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAzureSecurityGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAzureSecurityGroupsOK creates a ListAzureSecurityGroupsOK with default headers values
func NewListAzureSecurityGroupsOK() *ListAzureSecurityGroupsOK {
	return &ListAzureSecurityGroupsOK{}
}

/*
ListAzureSecurityGroupsOK describes a response with status code 200, with default header values.

AzureSecurityGroupsList
*/
type ListAzureSecurityGroupsOK struct {
	Payload *models.AzureSecurityGroupsList
}

// IsSuccess returns true when this list azure security groups o k response has a 2xx status code
func (o *ListAzureSecurityGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list azure security groups o k response has a 3xx status code
func (o *ListAzureSecurityGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list azure security groups o k response has a 4xx status code
func (o *ListAzureSecurityGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list azure security groups o k response has a 5xx status code
func (o *ListAzureSecurityGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list azure security groups o k response a status code equal to that given
func (o *ListAzureSecurityGroupsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAzureSecurityGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/azure/securitygroups][%d] listAzureSecurityGroupsOK  %+v", 200, o.Payload)
}

func (o *ListAzureSecurityGroupsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/azure/securitygroups][%d] listAzureSecurityGroupsOK  %+v", 200, o.Payload)
}

func (o *ListAzureSecurityGroupsOK) GetPayload() *models.AzureSecurityGroupsList {
	return o.Payload
}

func (o *ListAzureSecurityGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureSecurityGroupsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAzureSecurityGroupsDefault creates a ListAzureSecurityGroupsDefault with default headers values
func NewListAzureSecurityGroupsDefault(code int) *ListAzureSecurityGroupsDefault {
	return &ListAzureSecurityGroupsDefault{
		_statusCode: code,
	}
}

/*
ListAzureSecurityGroupsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAzureSecurityGroupsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list azure security groups default response
func (o *ListAzureSecurityGroupsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list azure security groups default response has a 2xx status code
func (o *ListAzureSecurityGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list azure security groups default response has a 3xx status code
func (o *ListAzureSecurityGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list azure security groups default response has a 4xx status code
func (o *ListAzureSecurityGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list azure security groups default response has a 5xx status code
func (o *ListAzureSecurityGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list azure security groups default response a status code equal to that given
func (o *ListAzureSecurityGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListAzureSecurityGroupsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/azure/securitygroups][%d] listAzureSecurityGroups default  %+v", o._statusCode, o.Payload)
}

func (o *ListAzureSecurityGroupsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/azure/securitygroups][%d] listAzureSecurityGroups default  %+v", o._statusCode, o.Payload)
}

func (o *ListAzureSecurityGroupsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAzureSecurityGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
