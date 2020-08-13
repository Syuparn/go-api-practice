module github.com/Syuparn/go-api-practice/client/api

go 1.13

replace github.com/Syuparn/go-api-practice/client/domain => ../domain

require (
	github.com/Syuparn/go-api-practice/client/domain v0.0.0
	github.com/gofrs/uuid v3.3.0+incompatible
	gopkg.in/h2non/gock.v1 v1.0.15
)
