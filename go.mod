module github.com/omniboost/go-sap-b1-service-layer

go 1.24.0

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/gorilla/schema v0.0.0-20171211162101-9fa3b6af65dc
	github.com/omniboost/go-omniboost-http-client v0.0.0-20250321143618-7cf3b125e8f0
	github.com/pkg/errors v0.9.1
	gopkg.in/guregu/null.v3 v3.5.0
)

require (
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
	golang.org/x/oauth2 v0.28.0 // indirect
)

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20191030093734-a170fe1a7240
