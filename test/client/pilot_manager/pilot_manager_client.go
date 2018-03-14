// Code generated by go-swagger; DO NOT EDIT.

package pilot_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new pilot manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pilot manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetSubtaskStatus gets subtask status
*/
func (a *Client) GetSubtaskStatus(params *GetSubtaskStatusParams) (*GetSubtaskStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSubtaskStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSubtaskStatus",
		Method:             "GET",
		PathPattern:        "/v1/pilots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSubtaskStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSubtaskStatusOK), nil

}

/*
HandleSubtask handles subtask
*/
func (a *Client) HandleSubtask(params *HandleSubtaskParams) (*HandleSubtaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHandleSubtaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "HandleSubtask",
		Method:             "POST",
		PathPattern:        "/v1/pilots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HandleSubtaskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HandleSubtaskOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}