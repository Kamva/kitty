package kitty

import (
	"errors"
	"net/http"
)

//--------------------------------
// Entity Adapter errors
//--------------------------------

// Error code description:
// kt = kitty  (package or project name)
// ett = replies about entity adapter section (identify some part in application)
// E = Error (type of code : error|response|...)
// 0 = error number zero (id of code in that part and type)

var (
	ErrInvalidID = NewError(http.StatusBadRequest, "kt.ett.e.0", "err_invalid_id", errors.New("id value is invalid"))
)
