curl -X POST "http://fss.com:3100/loki/api/v1/push" \
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

curl -G "http://fss.com:3100/loki/api/v1/query" \
-H "X-Scope-OrgID: tenant1" \
  --data-urlencode "query={service_name=\"service-monitor\"}" \
  --data-urlencode "limit=10" | jq



curl -G -s  "http://fss.com:3100/loki/api/v1/query" \
  -H "X-Scope-OrgID: tenant1" \
  --data-urlencode \
  'query=sum(rate({job="varlogs"}[10m])) by (level)' | jq


curl http://fss.com:3100/loki/api/v1/tail \
-H "X-Scope-OrgID: tenant1"


curl -G -s  "http://localhost:3100/loki/api/v1/labels" \
-H "X-Scope-OrgID: tenant1" \
| jq
# {
#   "status": "success",
#   "data": [
#     "foo",
#     "bar",
#     "baz"
#   ]
# }