module github.com/bachue/sufy-sdk-go-test

go 1.19

require (
	github.com/aws/smithy-go v1.13.5
	github.com/bachue/sufy-sdk-go v0.0.0-00010101000000-000000000000
	github.com/bachue/sufy-sdk-go/credentials v0.0.0-00010101000000-000000000000
	github.com/bachue/sufy-sdk-go/service/object v0.0.0-00010101000000-000000000000
	github.com/davecgh/go-spew v1.1.1
	github.com/stretchr/testify v1.8.2
)

require (
	github.com/aws/aws-sdk-go-v2 v1.17.4 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.10 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.13.12 // indirect
	github.com/bachue/sufy-sdk-go/internal/configsources v1.1.28 // indirect
	github.com/bachue/sufy-sdk-go/internal/endpoints/v2 v2.4.22 // indirect
	github.com/bachue/sufy-sdk-go/service/internal/accept-encoding v1.0.5 // indirect
	github.com/bachue/sufy-sdk-go/service/internal/sufyobjectshared v1.2.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/bachue/sufy-sdk-go => ../sufy-sdk-go

replace github.com/bachue/sufy-sdk-go/credentials => ../sufy-sdk-go/credentials

replace github.com/bachue/sufy-sdk-go/service/object => ../sufy-sdk-go/service/object

replace github.com/bachue/sufy-sdk-go/service/internal/accept-encoding => ../sufy-sdk-go/service/internal/accept-encoding/

replace github.com/bachue/sufy-sdk-go/service/internal/sufyobjectshared => ../sufy-sdk-go/service/internal/sufyobjectshared/

replace github.com/bachue/sufy-sdk-go/internal/configsources => ../sufy-sdk-go/internal/configsources/

replace github.com/bachue/sufy-sdk-go/internal/endpoints/v2 => ../sufy-sdk-go/internal/endpoints/v2/
