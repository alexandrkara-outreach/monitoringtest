# monitoringtest
Monitoring and tracing test project

Run it as:

curl -v http://127.0.0.1:8000/super/endpoint \
-H "X-Honeycomb-Event-Time: 2020-11-30T15:15:15.000000Z" \
    -d '{"some_parameter": 100}'