# Test Server: Simulating Fast/Slow Paths and GC pressure

## Overview

- /fast: A quick response, ideal for throughput testing.
- /slow: Simulates latency and contention.
- /gc: Simulate GC heavy workflow.
- net/http/pprof: Exposes runtime profiling on localhost:6060.

To start server run command:

```bash
go run server.go
```

## Simulating Load: Tools That Reflect Reality


| Tool   | Focus                         | Scriptable          | Metrics Depth                 | Ideal Use Case                                          |
|--------|-------------------------------|---------------------|-------------------------------|---------------------------------------------------------|
| vegeta | Constant rate load generation | No (but composable) | High (histogram, percentiles) | Tracking latency percentiles over time; CI benchmarking |
| wrk    | Max throughput stress tests   | Yes (Lua)           | Medium                        | Measuring raw server capacity and concurrency limits    |
| k6     | Scenario-based simulation     | Yes (JavaScript)    | High (VU metrics, dashboards) | Simulating real-world user workflows and pacing         |
