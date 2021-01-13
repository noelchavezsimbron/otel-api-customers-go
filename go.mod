module api-customers

go 1.14

require (
	github.com/gorilla/mux v1.8.0
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.9.0
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	go.mongodb.org/mongo-driver v1.4.4
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.15.1
	go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo v0.15.1
	go.opentelemetry.io/otel v0.15.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.15.0
	go.opentelemetry.io/otel/sdk v0.15.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)
