module github.com/Syuparn/go-api-practice/client

go 1.13

// 直下のモジュール参照のためのエイリアス（つけないと、既にpushしたものしか参照できない）
replace (
	github.com/Syuparn/go-api-practice/client/api => ./api
	github.com/Syuparn/go-api-practice/client/controller => ./controller
	github.com/Syuparn/go-api-practice/client/domain => ./domain
	github.com/Syuparn/go-api-practice/client/view => ./view
)

require (
	github.com/Syuparn/go-api-practice/client/api v0.0.0
	github.com/Syuparn/go-api-practice/client/controller v0.0.0
)
