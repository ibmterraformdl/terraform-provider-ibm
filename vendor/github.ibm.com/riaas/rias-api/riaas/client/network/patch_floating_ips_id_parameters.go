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

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// NewPatchFloatingIpsIDParams creates a new PatchFloatingIpsIDParams object
// with the default values initialized.
func NewPatchFloatingIpsIDParams() *PatchFloatingIpsIDParams {
	var ()
	return &PatchFloatingIpsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchFloatingIpsIDParamsWithTimeout creates a new PatchFloatingIpsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchFloatingIpsIDParamsWithTimeout(timeout time.Duration) *PatchFloatingIpsIDParams {
	var ()
	return &PatchFloatingIpsIDParams{

		timeout: timeout,
	}
}

// NewPatchFloatingIpsIDParamsWithContext creates a new PatchFloatingIpsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchFloatingIpsIDParamsWithContext(ctx context.Context) *PatchFloatingIpsIDParams {
	var ()
	return &PatchFloatingIpsIDParams{

		Context: ctx,
	}
}

// NewPatchFloatingIpsIDParamsWithHTTPClient creates a new PatchFloatingIpsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchFloatingIpsIDParamsWithHTTPClient(client *http.Client) *PatchFloatingIpsIDParams {
	var ()
	return &PatchFloatingIpsIDParams{
		HTTPClient: client,
	}
}

/*PatchFloatingIpsIDParams contains all the parameters to send to the API endpoint
for the patch floating ips ID operation typically these are written to a http.Request
*/
type PatchFloatingIpsIDParams struct {

	/*Body*/
	Body *models.PatchFloatingIpsIDParamsBody
	/*ID
	  The floating IP address

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch floating ips ID params
func (o *PatchFloatingIpsIDParams) WithTimeout(timeout time.Duration) *PatchFloatingIpsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch floating ips ID params
func (o *PatchFloatingIpsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch floating ips ID params
func (o *PatchFloatingIpsIDParams) WithContext(ctx context.Context) *PatchFloatingIpsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch floating ips ID params
func (o *PatchFloatingIpsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch floating ips ID params
func (o *PatchFloatingIpsIDParams) WithHTTPClient(client *http.Client) *PatchFloatingIpsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch floating ips ID params
func (o *PatchFloatingIpsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch floating ips ID params
func (o *PatchFloatingIpsIDParams) WithBody(body *models.PatchFloatingIpsIDParamsBody) *PatchFloatingIpsIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch floating ips ID params
func (o *PatchFloatingIpsIDParams) SetBody(body *models.PatchFloatingIpsIDParamsBody) {
	o.Body = body
}

// WithID adds the id to the patch floating ips ID params
func (o *PatchFloatingIpsIDParams) WithID(id string) *PatchFloatingIpsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch floating ips ID params
func (o *PatchFloatingIpsIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the patch floating ips ID params
func (o *PatchFloatingIpsIDParams) WithVersion(version string) *PatchFloatingIpsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the patch floating ips ID params
func (o *PatchFloatingIpsIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PatchFloatingIpsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}