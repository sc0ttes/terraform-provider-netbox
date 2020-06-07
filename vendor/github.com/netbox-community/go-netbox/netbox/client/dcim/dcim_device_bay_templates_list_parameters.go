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

package dcim

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

// NewDcimDeviceBayTemplatesListParams creates a new DcimDeviceBayTemplatesListParams object
// with the default values initialized.
func NewDcimDeviceBayTemplatesListParams() *DcimDeviceBayTemplatesListParams {
	var ()
	return &DcimDeviceBayTemplatesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBayTemplatesListParamsWithTimeout creates a new DcimDeviceBayTemplatesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceBayTemplatesListParamsWithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesListParams {
	var ()
	return &DcimDeviceBayTemplatesListParams{

		timeout: timeout,
	}
}

// NewDcimDeviceBayTemplatesListParamsWithContext creates a new DcimDeviceBayTemplatesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceBayTemplatesListParamsWithContext(ctx context.Context) *DcimDeviceBayTemplatesListParams {
	var ()
	return &DcimDeviceBayTemplatesListParams{

		Context: ctx,
	}
}

// NewDcimDeviceBayTemplatesListParamsWithHTTPClient creates a new DcimDeviceBayTemplatesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceBayTemplatesListParamsWithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesListParams {
	var ()
	return &DcimDeviceBayTemplatesListParams{
		HTTPClient: client,
	}
}

/*DcimDeviceBayTemplatesListParams contains all the parameters to send to the API endpoint
for the dcim device bay templates list operation typically these are written to a http.Request
*/
type DcimDeviceBayTemplatesListParams struct {

	/*DevicetypeID*/
	DevicetypeID *string
	/*DevicetypeIDn*/
	DevicetypeIDn *string
	/*ID*/
	ID *string
	/*IDGt*/
	IDGt *string
	/*IDGte*/
	IDGte *string
	/*IDLt*/
	IDLt *string
	/*IDLte*/
	IDLte *string
	/*IDn*/
	IDn *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*NameIc*/
	NameIc *string
	/*NameIe*/
	NameIe *string
	/*NameIew*/
	NameIew *string
	/*NameIsw*/
	NameIsw *string
	/*Namen*/
	Namen *string
	/*NameNic*/
	NameNic *string
	/*NameNie*/
	NameNie *string
	/*NameNiew*/
	NameNiew *string
	/*NameNisw*/
	NameNisw *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithContext(ctx context.Context) *DcimDeviceBayTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevicetypeID adds the devicetypeID to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithDevicetypeID(devicetypeID *string) *DcimDeviceBayTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetDevicetypeID(devicetypeID *string) {
	o.DevicetypeID = devicetypeID
}

// WithDevicetypeIDn adds the devicetypeIDn to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithDevicetypeIDn(devicetypeIDn *string) *DcimDeviceBayTemplatesListParams {
	o.SetDevicetypeIDn(devicetypeIDn)
	return o
}

// SetDevicetypeIDn adds the devicetypeIdN to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetDevicetypeIDn(devicetypeIDn *string) {
	o.DevicetypeIDn = devicetypeIDn
}

// WithID adds the id to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithID(id *string) *DcimDeviceBayTemplatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithIDGt(iDGt *string) *DcimDeviceBayTemplatesListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithIDGte(iDGte *string) *DcimDeviceBayTemplatesListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithIDLt(iDLt *string) *DcimDeviceBayTemplatesListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithIDLte(iDLte *string) *DcimDeviceBayTemplatesListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithIDn(iDn *string) *DcimDeviceBayTemplatesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLimit adds the limit to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithLimit(limit *int64) *DcimDeviceBayTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithName(name *string) *DcimDeviceBayTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameIc(nameIc *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameIe(nameIe *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameIew(nameIew *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameIsw(nameIsw *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNamen(namen *string) *DcimDeviceBayTemplatesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameNic(nameNic *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameNie(nameNie *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameNiew(nameNiew *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameNisw(nameNisw *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithOffset(offset *int64) *DcimDeviceBayTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithQ(q *string) *DcimDeviceBayTemplatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBayTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DevicetypeID != nil {

		// query param devicetype_id
		var qrDevicetypeID string
		if o.DevicetypeID != nil {
			qrDevicetypeID = *o.DevicetypeID
		}
		qDevicetypeID := qrDevicetypeID
		if qDevicetypeID != "" {
			if err := r.SetQueryParam("devicetype_id", qDevicetypeID); err != nil {
				return err
			}
		}

	}

	if o.DevicetypeIDn != nil {

		// query param devicetype_id__n
		var qrDevicetypeIDn string
		if o.DevicetypeIDn != nil {
			qrDevicetypeIDn = *o.DevicetypeIDn
		}
		qDevicetypeIDn := qrDevicetypeIDn
		if qDevicetypeIDn != "" {
			if err := r.SetQueryParam("devicetype_id__n", qDevicetypeIDn); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string
		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {
			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}

	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string
		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {
			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}

	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string
		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {
			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}

	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string
		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {
			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}

	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string
		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {
			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string
		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {
			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}

	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string
		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {
			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}

	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string
		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {
			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}

	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string
		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {
			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}

	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string
		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {
			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}

	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string
		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {
			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}

	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string
		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {
			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}

	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string
		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {
			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}

	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string
		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {
			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
