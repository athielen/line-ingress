# line-ingress
Privacy first customer event collection aggregator. Phase 1 of privacy first digital customer tracking.

## Breakdown

Endpoints accept a POST to capture features or a GET request for a 1x1 .gif.

Handlers will take these requests, marshal them to a json representation of desired information.

Handlers will pass these json []byte along to backend manager to convert them to a line struct and place them in the applicable channels

## TODO

Marshal Json improved: collect ip address, get country of origin from country.go, nil or zero value for missing features

file_backend: write to file, add config to use this backend only in local/dev mode

redis_backend: get redis client, write to redis instance

handlers: clean up handlers with more official naming conventions, more error handling and correct response codes