// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// GetInstanceProfilesReader is a Reader for the GetInstanceProfiles structure.
type GetInstanceProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstanceProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstanceProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetInstanceProfilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetInstanceProfilesOK creates a GetInstanceProfilesOK with default headers values
func NewGetInstanceProfilesOK() *GetInstanceProfilesOK {
	return &GetInstanceProfilesOK{}
}

/*GetInstanceProfilesOK handles this case with default header values.

instance profiles retrieved successfully
*/
type GetInstanceProfilesOK struct {
	Payload *models.GetInstanceProfilesOKBody
}

func (o *GetInstanceProfilesOK) Error() string {
	return fmt.Sprintf("[GET /instance/profiles][%d] getInstanceProfilesOK  %+v", 200, o.Payload)
}

func (o *GetInstanceProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetInstanceProfilesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceProfilesDefault creates a GetInstanceProfilesDefault with default headers values
func NewGetInstanceProfilesDefault(code int) *GetInstanceProfilesDefault {
	return &GetInstanceProfilesDefault{
		_statusCode: code,
	}
}

/*GetInstanceProfilesDefault handles this case with default header values.

unexpectederror
*/
type GetInstanceProfilesDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get instance profiles default response
func (o *GetInstanceProfilesDefault) Code() int {
	return o._statusCode
}

func (o *GetInstanceProfilesDefault) Error() string {
	return fmt.Sprintf("[GET /instance/profiles][%d] GetInstanceProfiles default  %+v", o._statusCode, o.Payload)
}

func (o *GetInstanceProfilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
