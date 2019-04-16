package types

import "errors"

// go does not (yet) support regular error types over gob

type RpcError string

var RpcErrorNotImplemented RpcError = "not implemented"
var RpcErrorNumTimeValuePairsMisMatch RpcError = "mismatch between number of time&value pairs"
var RpcErrorBackendStrategyNotFound RpcError = "no backend strategy found"

func (err RpcError) String() string {
	return string(err)
}

func (err RpcError) Error() error {
	return errors.New(err.String())
}