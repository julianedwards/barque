/*
Package barque holds a a number of application level constants and
shared resources for the cedar application.
*/
package barque

import (
	"time"
)

const ()

// BuildRevision stores the commit in the git repository at build time
// and is specified with -ldflags at build time
var BuildRevision = ""

const (
	AuthTokenCookie  = "barque-token"
	APIUserHeader    = "Api-User"
	APIKeyHeader     = "Api-Key"
	QueueName        = "barque.service"
	TokenExpireAfter = time.Hour
)
