### Domain CNAME Lookup

This is a simple command line program written in Go that looks up and prints the CNAME record for a given domain name.
Usage

The domain name to lookup is specified using the -d flag:

```go run main.go -d example.com```

This will print :

```CNAME Record for example.com is: www.example.net```

If there is no CNAME record configured for the given domain, the program will output an error.

## Examples

Look up CNAME for google.com:


```go run main.go -d google.com```

Prints:

```CNAME Record for google.com is: www.google.com```

Try a domain without a CNAME:


```go run main.go -d example.org```

Prints:


```lookup example.org: no such host```

## Requirements:

    Go 1.11 or higher
