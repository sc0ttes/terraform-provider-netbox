// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewIpamVlansReadParams creates a new IpamVlansReadParams object
// with the default values initialized.
func NewIpamVlansReadParams() *IpamVlansReadParams {
	var ()
	return &IpamVlansReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlansReadParamsWithTimeout creates a new IpamVlansReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamVlansReadParamsWithTimeout(timeout time.Duration) *IpamVlansReadParams {
	var ()
	return &IpamVlansReadParams{

		timeout: timeout,
	}
}

// NewIpamVlansReadParamsWithContext creates a new IpamVlansReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamVlansReadParamsWithContext(ctx context.Context) *IpamVlansReadParams {
	var ()
	return &IpamVlansReadParams{

		Context: ctx,
	}
}

// NewIpamVlansReadParamsWithHTTPClient creates a new IpamVlansReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamVlansReadParamsWithHTTPClient(client *http.Client) *IpamVlansReadParams {
	var ()
	return &IpamVlansReadParams{
		HTTPClient: client,
	}
}

/*IpamVlansReadParams contains all the parameters to send to the API endpoint
for the ipam vlans read operation typically these are written to a http.Request
*/
type IpamVlansReadParams struct {

	/*ID
	  A unique integer value identifying this VLAN.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam vlans read params
func (o *IpamVlansReadParams) WithTimeout(timeout time.Duration) *IpamVlansReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlans read params
func (o *IpamVlansReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlans read params
func (o *IpamVlansReadParams) WithContext(ctx context.Context) *IpamVlansReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlans read params
func (o *IpamVlansReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlans read params
func (o *IpamVlansReadParams) WithHTTPClient(client *http.Client) *IpamVlansReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlans read params
func (o *IpamVlansReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam vlans read params
func (o *IpamVlansReadParams) WithID(id int64) *IpamVlansReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlans read params
func (o *IpamVlansReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlansReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
