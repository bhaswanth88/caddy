package standard

import (
	// standard Caddy HTTP app modules
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/caddyauth"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/encode"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/encode/brotli"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/encode/gzip"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/encode/zstd"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/fileserver"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/headers"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/map"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/proxyprotocol"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/push"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/requestbody"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/reverseproxy"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/reverseproxy/fastcgi"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/reverseproxy/forwardauth"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/rewrite"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/templates"
	_ "github.com/bhaswanth88/caddy/v2/modules/caddyhttp/tracing"
)
