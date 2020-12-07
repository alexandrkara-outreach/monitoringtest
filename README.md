# monitoringtest
Monitoring and tracing test project

## Dependencies

To start optional datadag agen

`
DD_API_KEY="SOME_API_KEY" docker-compose up
`


## Run

go run cmd/example-monitoring/main.go

## Test


Run it as:


`
curl -v http://127.0.0.1:8000/super/endpoint \
-H "X-Honeycomb-Event-Time: 2020-11-30T15:15:15.000000Z" \
-d '{"some_parameter": 100}'
`

or 

`
ab -n 1000 -c 2 http://localhost:8000/super/endpoint
`
