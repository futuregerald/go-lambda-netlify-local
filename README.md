# Netlify lambda Go function local testing example

This repo is an example of deploying Go lambda functions using a connector library that lets you use Go standard lib compliant http handlers in lambdas. The biggest benefit of doing this is that you can run the handlers locally by having a conditional that checks for an environment variable to determine whether to run `http.ListenAndServe()` for local dev vs `gateway.ListenAndServe()` for lambdas.

The other benefits are that you can re-use existing handlers, middleware, routers, etc., and do complete tests of your handler without running it in a docker image.

## Examples
This repo contains three examples that highlight three common patterns

### testfunc1
The first example is basic function that just uses the standard lib to create a handler and doesn't doa nything special

### testfunc2
This example uses chi router and logger middleware. This shows how you can easily use third-party routers to automatically get some of their convenience features, like filtering requests by methods, param parsing, built-in middleware, etc.

### testfunc3context
This uses a standard lib compliant handler but shows how to access the lambda context in case you need any of the information that Netlify injects into your function, like Netlify Identity user information, or site URL.