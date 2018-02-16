Info:
- Simplest example of using goa.design. See - https://goa.design/learn/guide/
- Added goa-middleware (https://goa.design/implement/middleware/) named MetricRecorder.  It extract EntityName and Action attributes for each request and print in log.  The same can be used to record metric (not coded here).

Run:
1. At goa-design-cellar, go build -o cellar
2. At goa-design-cellar/tool/cellar-cli, go build -o cellar-cli

Reference:
- https://goa.design/learn/guide/
