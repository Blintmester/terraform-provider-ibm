// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// DeleteIkePoliciesIDReader is a Reader for the DeleteIkePoliciesID structure.
type DeleteIkePoliciesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIkePoliciesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteIkePoliciesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteIkePoliciesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteIkePoliciesIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteIkePoliciesIDNoContent creates a DeleteIkePoliciesIDNoContent with default headers values
func NewDeleteIkePoliciesIDNoContent() *DeleteIkePoliciesIDNoContent {
	return &DeleteIkePoliciesIDNoContent{}
}

/*DeleteIkePoliciesIDNoContent handles this case with default header values.

The IKE policy was deleted successfully.
*/
type DeleteIkePoliciesIDNoContent struct {
}

func (o *DeleteIkePoliciesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ike_policies/{id}][%d] deleteIkePoliciesIdNoContent ", 204)
}

func (o *DeleteIkePoliciesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIkePoliciesIDNotFound creates a DeleteIkePoliciesIDNotFound with default headers values
func NewDeleteIkePoliciesIDNotFound() *DeleteIkePoliciesIDNotFound {
	return &DeleteIkePoliciesIDNotFound{}
}

/*DeleteIkePoliciesIDNotFound handles this case with default header values.

An IKE policy with the specified identifier could not be found.
*/
type DeleteIkePoliciesIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteIkePoliciesIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /ike_policies/{id}][%d] deleteIkePoliciesIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIkePoliciesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIkePoliciesIDConflict creates a DeleteIkePoliciesIDConflict with default headers values
func NewDeleteIkePoliciesIDConflict() *DeleteIkePoliciesIDConflict {
	return &DeleteIkePoliciesIDConflict{}
}

/*DeleteIkePoliciesIDConflict handles this case with default header values.

The IKE policy is in use and cannot be removed.
*/
type DeleteIkePoliciesIDConflict struct {
	Payload *models.Riaaserror
}

func (o *DeleteIkePoliciesIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /ike_policies/{id}][%d] deleteIkePoliciesIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteIkePoliciesIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
