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

// NewGetSubnetsSubnetIDPublicGatewayParams creates a new GetSubnetsSubnetIDPublicGatewayParams object
// with the default values initialized.
func NewGetSubnetsSubnetIDPublicGatewayParams() *GetSubnetsSubnetIDPublicGatewayParams {
	var ()
	return &GetSubnetsSubnetIDPublicGatewayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubnetsSubnetIDPublicGatewayParamsWithTimeout creates a new GetSubnetsSubnetIDPublicGatewayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSubnetsSubnetIDPublicGatewayParamsWithTimeout(timeout time.Duration) *GetSubnetsSubnetIDPublicGatewayParams {
	var ()
	return &GetSubnetsSubnetIDPublicGatewayParams{

		timeout: timeout,
	}
}

// NewGetSubnetsSubnetIDPublicGatewayParamsWithContext creates a new GetSubnetsSubnetIDPublicGatewayParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSubnetsSubnetIDPublicGatewayParamsWithContext(ctx context.Context) *GetSubnetsSubnetIDPublicGatewayParams {
	var ()
	return &GetSubnetsSubnetIDPublicGatewayParams{

		Context: ctx,
	}
}

// NewGetSubnetsSubnetIDPublicGatewayParamsWithHTTPClient creates a new GetSubnetsSubnetIDPublicGatewayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSubnetsSubnetIDPublicGatewayParamsWithHTTPClient(client *http.Client) *GetSubnetsSubnetIDPublicGatewayParams {
	var ()
	return &GetSubnetsSubnetIDPublicGatewayParams{
		HTTPClient: client,
	}
}

/*GetSubnetsSubnetIDPublicGatewayParams contains all the parameters to send to the API endpoint
for the get subnets subnet ID public gateway operation typically these are written to a http.Request
*/
type GetSubnetsSubnetIDPublicGatewayParams struct {

	/*SubnetID
	  The subnet identifier

	*/
	SubnetID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get subnets subnet ID public gateway params
func (o *GetSubnetsSubnetIDPublicGatewayParams) WithTimeout(timeout time.Duration) *GetSubnetsSubnetIDPublicGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get subnets subnet ID public gateway params
func (o *GetSubnetsSubnetIDPublicGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get subnets subnet ID public gateway params
func (o *GetSubnetsSubnetIDPublicGatewayParams) WithContext(ctx context.Context) *GetSubnetsSubnetIDPublicGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get subnets subnet ID public gateway params
func (o *GetSubnetsSubnetIDPublicGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get subnets subnet ID public gateway params
func (o *GetSubnetsSubnetIDPublicGatewayParams) WithHTTPClient(client *http.Client) *GetSubnetsSubnetIDPublicGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get subnets subnet ID public gateway params
func (o *GetSubnetsSubnetIDPublicGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubnetID adds the subnetID to the get subnets subnet ID public gateway params
func (o *GetSubnetsSubnetIDPublicGatewayParams) WithSubnetID(subnetID string) *GetSubnetsSubnetIDPublicGatewayParams {
	o.SetSubnetID(subnetID)
	return o
}

// SetSubnetID adds the subnetId to the get subnets subnet ID public gateway params
func (o *GetSubnetsSubnetIDPublicGatewayParams) SetSubnetID(subnetID string) {
	o.SubnetID = subnetID
}

// WithVersion adds the version to the get subnets subnet ID public gateway params
func (o *GetSubnetsSubnetIDPublicGatewayParams) WithVersion(version string) *GetSubnetsSubnetIDPublicGatewayParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get subnets subnet ID public gateway params
func (o *GetSubnetsSubnetIDPublicGatewayParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubnetsSubnetIDPublicGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param subnet_id
	if err := r.SetPathParam("subnet_id", o.SubnetID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
