module github.com/Syuparn/go-api-practice/client/controller

go 1.13

replace (
	github.com/Syuparn/go-api-practice/client/domain => ../domain
	github.com/Syuparn/go-api-practice/client/view => ../view
)

require (
	github.com/Syuparn/go-api-practice/client/domain v0.0.0
	github.com/Syuparn/go-api-practice/client/view v0.0.0
	github.com/gofrs/uuid v3.3.0+incompatible
)
