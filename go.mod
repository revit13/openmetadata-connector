module fybrik.io/openmetadata-connector

go 1.17

require (
	fybrik.io/fybrik v1.1.0
	fybrik.io/openmetadata-connector/datacatalog-go v0.0.0
	fybrik.io/openmetadata-connector/datacatalog-go-client v0.0.0
	fybrik.io/openmetadata-connector/datacatalog-go-models v0.0.0
	github.com/hashicorp/go-retryablehttp v0.7.1
	github.com/rs/zerolog v1.26.0
	github.com/spf13/cobra v1.5.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/fatih/structs v1.1.0
	github.com/onsi/ginkgo/v2 v2.1.4
	github.com/onsi/gomega v1.19.0
)

replace fybrik.io/openmetadata-connector/datacatalog-go => ./auto-generated/api

replace fybrik.io/openmetadata-connector/datacatalog-go-models => ./auto-generated/models

replace fybrik.io/openmetadata-connector/datacatalog-go-client => ./auto-generated/client
