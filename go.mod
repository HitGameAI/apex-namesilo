module github.com/kamalyes/go-namesilo

go 1.25.0

require (
	github.com/kamalyes/go-logger v0.9.2
	github.com/kamalyes/go-toolbox v0.19.3
	github.com/stretchr/testify v1.12.1
)

// 替换为HitGameAI的版本
replace (
	github.com/kamalyes/go-argus v0.5.2 => github.com/HitGameAI/go-argus v0.5.2
	github.com/kamalyes/go-logger v0.9.2 => github.com/HitGameAI/go-logger v0.9.2
	github.com/kamalyes/go-toolbox v0.19.3 => github.com/HitGameAI/go-toolbox v0.19.3
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/kamalyes/go-argus v0.5.2 // indirect
	go.opentelemetry.io/otel v1.44.0 // indirect
	go.opentelemetry.io/otel/trace v1.44.0 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/sys v0.37.0 // indirect
	google.golang.org/grpc v1.77.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
