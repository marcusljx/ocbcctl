package vars

import "errors"

var (
	ErrInternal = errors.New("internal error, please view logs")

	ErrCallbackHostEnvNotSet  = errors.New("OCBCCTL_CALLBACK_HOST is not set")
	ErrCallbackHostInvalidURL = errors.New("OCBCCTL_CALLBACK_HOST is an invalid URL")
)
