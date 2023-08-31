curl -X POST "http://localhost:3100/loki/api/v1/push" \
-H "Content-Type: application/json" \
-H "X-Scope-OrgID: tenant1" \
-d '{
  "streams": [
    {
      "labels": {
        "service_name": "service-monitor",
        "level": "info"
      },
      "entries": [
        {
          "ts": "2023-08-30T20:37:56Z",
          "line": "This is an info log message."
        }
      ]
    }
  ]
}'