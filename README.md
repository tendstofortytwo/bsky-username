# bsky-username

reads a map from domain-to-DID stored in `db.json` and serves DIDs depending on the host in the HTTP request, allowing you to add custom handles for a domain easily.

## usage

1. start the server

	```
	go run . -port=3000
	```
2. redirect requests (caddyfile example)

	```
	*.enby.club {
		reverse_proxy localhost:3000
	}
	```

## license

MIT license; see LICENSE.md.
