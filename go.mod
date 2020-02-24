module github.com/lemon-cloud-service/lemon-cloud-gateway

go 1.13

require (
	github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components v0.0.0-00010101000000-000000000000
	github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils v0.0.0-00010101000000-000000000000
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.1.1
	github.com/micro/go-plugins/micro/cors v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/micro v1.18.0
	github.com/micro/micro/v2 v2.1.0
)

replace github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils => ../lemon-cloud-common/lemon-cloud-common-utils

replace github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components => ../lemon-cloud-common/lemon-cloud-common-components

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
