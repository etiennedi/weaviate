/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateActionsPropertiesCreateHandlerFunc turns a function with the right signature into a weaviate actions properties create handler
type WeaviateActionsPropertiesCreateHandlerFunc func(WeaviateActionsPropertiesCreateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateActionsPropertiesCreateHandlerFunc) Handle(params WeaviateActionsPropertiesCreateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateActionsPropertiesCreateHandler interface for that can handle valid weaviate actions properties create params
type WeaviateActionsPropertiesCreateHandler interface {
	Handle(WeaviateActionsPropertiesCreateParams, interface{}) middleware.Responder
}

// NewWeaviateActionsPropertiesCreate creates a new http.Handler for the weaviate actions properties create operation
func NewWeaviateActionsPropertiesCreate(ctx *middleware.Context, handler WeaviateActionsPropertiesCreateHandler) *WeaviateActionsPropertiesCreate {
	return &WeaviateActionsPropertiesCreate{Context: ctx, Handler: handler}
}

/*WeaviateActionsPropertiesCreate swagger:route POST /actions/{actionId}/properties/{propertyName} actions weaviateActionsPropertiesCreate

Add a single reference to a class-property when cardinality is set to 'hasMany'.

Add a single reference to a class-property when cardinality is set to 'hasMany'.

*/
type WeaviateActionsPropertiesCreate struct {
	Context *middleware.Context
	Handler WeaviateActionsPropertiesCreateHandler
}

func (o *WeaviateActionsPropertiesCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateActionsPropertiesCreateParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
