package company

import "errors"

var ErrUnexpected = errors.New("unexpected internal error")
var ErrCompanyNotFound = errors.New("company not found")
