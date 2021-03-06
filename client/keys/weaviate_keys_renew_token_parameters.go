// Code generated by go-swagger; DO NOT EDIT.

package keys

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
)

// NewWeaviateKeysRenewTokenParams creates a new WeaviateKeysRenewTokenParams object
// with the default values initialized.
func NewWeaviateKeysRenewTokenParams() *WeaviateKeysRenewTokenParams {
	var ()
	return &WeaviateKeysRenewTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateKeysRenewTokenParamsWithTimeout creates a new WeaviateKeysRenewTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateKeysRenewTokenParamsWithTimeout(timeout time.Duration) *WeaviateKeysRenewTokenParams {
	var ()
	return &WeaviateKeysRenewTokenParams{

		timeout: timeout,
	}
}

// NewWeaviateKeysRenewTokenParamsWithContext creates a new WeaviateKeysRenewTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateKeysRenewTokenParamsWithContext(ctx context.Context) *WeaviateKeysRenewTokenParams {
	var ()
	return &WeaviateKeysRenewTokenParams{

		Context: ctx,
	}
}

// NewWeaviateKeysRenewTokenParamsWithHTTPClient creates a new WeaviateKeysRenewTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateKeysRenewTokenParamsWithHTTPClient(client *http.Client) *WeaviateKeysRenewTokenParams {
	var ()
	return &WeaviateKeysRenewTokenParams{
		HTTPClient: client,
	}
}

/*WeaviateKeysRenewTokenParams contains all the parameters to send to the API endpoint
for the weaviate keys renew token operation typically these are written to a http.Request
*/
type WeaviateKeysRenewTokenParams struct {

	/*KeyID
	  Unique ID of the key.

	*/
	KeyID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate keys renew token params
func (o *WeaviateKeysRenewTokenParams) WithTimeout(timeout time.Duration) *WeaviateKeysRenewTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate keys renew token params
func (o *WeaviateKeysRenewTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate keys renew token params
func (o *WeaviateKeysRenewTokenParams) WithContext(ctx context.Context) *WeaviateKeysRenewTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate keys renew token params
func (o *WeaviateKeysRenewTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate keys renew token params
func (o *WeaviateKeysRenewTokenParams) WithHTTPClient(client *http.Client) *WeaviateKeysRenewTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate keys renew token params
func (o *WeaviateKeysRenewTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyID adds the keyID to the weaviate keys renew token params
func (o *WeaviateKeysRenewTokenParams) WithKeyID(keyID strfmt.UUID) *WeaviateKeysRenewTokenParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the weaviate keys renew token params
func (o *WeaviateKeysRenewTokenParams) SetKeyID(keyID strfmt.UUID) {
	o.KeyID = keyID
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateKeysRenewTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param keyId
	if err := r.SetPathParam("keyId", o.KeyID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
