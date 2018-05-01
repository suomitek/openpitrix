// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewDeleteClusterNodesParams creates a new DeleteClusterNodesParams object
// with the default values initialized.
func NewDeleteClusterNodesParams() *DeleteClusterNodesParams {
	var ()
	return &DeleteClusterNodesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterNodesParamsWithTimeout creates a new DeleteClusterNodesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClusterNodesParamsWithTimeout(timeout time.Duration) *DeleteClusterNodesParams {
	var ()
	return &DeleteClusterNodesParams{

		timeout: timeout,
	}
}

// NewDeleteClusterNodesParamsWithContext creates a new DeleteClusterNodesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClusterNodesParamsWithContext(ctx context.Context) *DeleteClusterNodesParams {
	var ()
	return &DeleteClusterNodesParams{

		Context: ctx,
	}
}

// NewDeleteClusterNodesParamsWithHTTPClient creates a new DeleteClusterNodesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClusterNodesParamsWithHTTPClient(client *http.Client) *DeleteClusterNodesParams {
	var ()
	return &DeleteClusterNodesParams{
		HTTPClient: client,
	}
}

/*DeleteClusterNodesParams contains all the parameters to send to the API endpoint
for the delete cluster nodes operation typically these are written to a http.Request
*/
type DeleteClusterNodesParams struct {

	/*Body*/
	Body *models.OpenpitrixDeleteClusterNodesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cluster nodes params
func (o *DeleteClusterNodesParams) WithTimeout(timeout time.Duration) *DeleteClusterNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster nodes params
func (o *DeleteClusterNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster nodes params
func (o *DeleteClusterNodesParams) WithContext(ctx context.Context) *DeleteClusterNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster nodes params
func (o *DeleteClusterNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster nodes params
func (o *DeleteClusterNodesParams) WithHTTPClient(client *http.Client) *DeleteClusterNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster nodes params
func (o *DeleteClusterNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete cluster nodes params
func (o *DeleteClusterNodesParams) WithBody(body *models.OpenpitrixDeleteClusterNodesRequest) *DeleteClusterNodesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete cluster nodes params
func (o *DeleteClusterNodesParams) SetBody(body *models.OpenpitrixDeleteClusterNodesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}