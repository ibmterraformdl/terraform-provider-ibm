// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// PostNetworkAclsReader is a Reader for the PostNetworkAcls structure.
type PostNetworkAclsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNetworkAclsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostNetworkAclsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostNetworkAclsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostNetworkAclsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewPostNetworkAclsCreated creates a PostNetworkAclsCreated with default headers values
func NewPostNetworkAclsCreated() *PostNetworkAclsCreated {
	return &PostNetworkAclsCreated{}
}

/*PostNetworkAclsCreated handles this case with default header values.

dummy
*/
type PostNetworkAclsCreated struct {
	Payload *models.NetworkACL
}

func (o *PostNetworkAclsCreated) Error() string {
	return fmt.Sprintf("[POST /network_acls][%d] postNetworkAclsCreated  %+v", 201, o.Payload)
}

func (o *PostNetworkAclsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkACL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNetworkAclsBadRequest creates a PostNetworkAclsBadRequest with default headers values
func NewPostNetworkAclsBadRequest() *PostNetworkAclsBadRequest {
	return &PostNetworkAclsBadRequest{}
}

/*PostNetworkAclsBadRequest handles this case with default header values.

error
*/
type PostNetworkAclsBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PostNetworkAclsBadRequest) Error() string {
	return fmt.Sprintf("[POST /network_acls][%d] postNetworkAclsBadRequest  %+v", 400, o.Payload)
}

func (o *PostNetworkAclsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNetworkAclsDefault creates a PostNetworkAclsDefault with default headers values
func NewPostNetworkAclsDefault(code int) *PostNetworkAclsDefault {
	return &PostNetworkAclsDefault{
		_statusCode: code,
	}
}

/*PostNetworkAclsDefault handles this case with default header values.

unexpectederror
*/
type PostNetworkAclsDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the post network acls default response
func (o *PostNetworkAclsDefault) Code() int {
	return o._statusCode
}

func (o *PostNetworkAclsDefault) Error() string {
	return fmt.Sprintf("[POST /network_acls][%d] PostNetworkAcls default  %+v", o._statusCode, o.Payload)
}

func (o *PostNetworkAclsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}