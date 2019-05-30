// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetVpcsVpcIDAddressPrefixesIDParams creates a new GetVpcsVpcIDAddressPrefixesIDParams object
// with the default values initialized.
func NewGetVpcsVpcIDAddressPrefixesIDParams() *GetVpcsVpcIDAddressPrefixesIDParams {
	var ()
	return &GetVpcsVpcIDAddressPrefixesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVpcsVpcIDAddressPrefixesIDParamsWithTimeout creates a new GetVpcsVpcIDAddressPrefixesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVpcsVpcIDAddressPrefixesIDParamsWithTimeout(timeout time.Duration) *GetVpcsVpcIDAddressPrefixesIDParams {
	var ()
	return &GetVpcsVpcIDAddressPrefixesIDParams{

		timeout: timeout,
	}
}

// NewGetVpcsVpcIDAddressPrefixesIDParamsWithContext creates a new GetVpcsVpcIDAddressPrefixesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVpcsVpcIDAddressPrefixesIDParamsWithContext(ctx context.Context) *GetVpcsVpcIDAddressPrefixesIDParams {
	var ()
	return &GetVpcsVpcIDAddressPrefixesIDParams{

		Context: ctx,
	}
}

// NewGetVpcsVpcIDAddressPrefixesIDParamsWithHTTPClient creates a new GetVpcsVpcIDAddressPrefixesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVpcsVpcIDAddressPrefixesIDParamsWithHTTPClient(client *http.Client) *GetVpcsVpcIDAddressPrefixesIDParams {
	var ()
	return &GetVpcsVpcIDAddressPrefixesIDParams{
		HTTPClient: client,
	}
}

/*GetVpcsVpcIDAddressPrefixesIDParams contains all the parameters to send to the API endpoint
for the get vpcs vpc ID address prefixes ID operation typically these are written to a http.Request
*/
type GetVpcsVpcIDAddressPrefixesIDParams struct {

	/*ID
	  The prefix identifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*VpcID
	  The VPC identifier

	*/
	VpcID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vpcs vpc ID address prefixes ID params
func (o *GetVpcsVpcIDAddressPrefixesIDParams) WithTimeout(timeout time.Duration) *GetVpcsVpcIDAddressPrefixesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vpcs vpc ID address prefixes ID params
func (o *GetVpcsVpcIDAddressPrefixesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vpcs vpc ID address prefixes ID params
func (o *GetVpcsVpcIDAddressPrefixesIDParams) WithContext(ctx context.Context) *GetVpcsVpcIDAddressPrefixesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vpcs vpc ID address prefixes ID params
func (o *GetVpcsVpcIDAddressPrefixesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vpcs vpc ID address prefixes ID params
func (o *GetVpcsVpcIDAddressPrefixesIDParams) WithHTTPClient(client *http.Client) *GetVpcsVpcIDAddressPrefixesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vpcs vpc ID address prefixes ID params
func (o *GetVpcsVpcIDAddressPrefixesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get vpcs vpc ID address prefixes ID params
func (o *GetVpcsVpcIDAddressPrefixesIDParams) WithID(id string) *GetVpcsVpcIDAddressPrefixesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vpcs vpc ID address prefixes ID params
func (o *GetVpcsVpcIDAddressPrefixesIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get vpcs vpc ID address prefixes ID params
func (o *GetVpcsVpcIDAddressPrefixesIDParams) WithVersion(version string) *GetVpcsVpcIDAddressPrefixesIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get vpcs vpc ID address prefixes ID params
func (o *GetVpcsVpcIDAddressPrefixesIDParams) SetVersion(version string) {
	o.Version = version
}

// WithVpcID adds the vpcID to the get vpcs vpc ID address prefixes ID params
func (o *GetVpcsVpcIDAddressPrefixesIDParams) WithVpcID(vpcID string) *GetVpcsVpcIDAddressPrefixesIDParams {
	o.SetVpcID(vpcID)
	return o
}

// SetVpcID adds the vpcId to the get vpcs vpc ID address prefixes ID params
func (o *GetVpcsVpcIDAddressPrefixesIDParams) SetVpcID(vpcID string) {
	o.VpcID = vpcID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVpcsVpcIDAddressPrefixesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	// path param vpc_id
	if err := r.SetPathParam("vpc_id", o.VpcID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
