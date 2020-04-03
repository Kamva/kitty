package gate

import (
	"errors"
	"github.com/Kamva/gutil"
	"github.com/Kamva/hexa"
)

// UserWithOwner define model that has owner.
type ResourceWithOwner interface {
	// Owner method returns id.
	OwnerID() hexa.ID
}

// UserOwnsResourcePolicy policy returns true if the user own provided resource.
func UserOwnsResourcePolicy(c hexa.Context, u hexa.User, r interface{}) (bool, error) {
	if gutil.IsNil(r) {
		return false, nil
	}

	if m, ok := r.(ResourceWithOwner); ok {
		return u.Identifier().Equal(m.OwnerID()), nil
	}
	return false, errors.New("provided resource is invalid")
}

// DefaultPolicy is default policy for gates.
var DefaultPolicy = UserOwnsResourcePolicy

// Assertion
var _ hexa.GatePolicy = UserOwnsResourcePolicy
