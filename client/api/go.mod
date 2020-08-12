module api

go 1.13

require (
	domain v0.0.0
	github.com/gofrs/uuid v3.3.0+incompatible
	gopkg.in/h2non/gock.v1 v1.0.15
)

// for importing local modules
replace domain => ../domain
