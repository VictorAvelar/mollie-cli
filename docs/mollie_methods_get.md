## mollie methods get

Retrieve a single method by its ID.

### Synopsis

Retrieve a single method by its ID. Note that if a method is not available on the website profile 
a status 404 Not found is returned. When the method is not enabled,a status 403 Forbidden 
is returned. You can enable payments methods via the Enable payment method endpoint in the 
Profiles API, or via your Mollie Dashboard.

```
mollie methods get [flags]
```

### Options

```
      --currency string   the currency to receiving the minimumAmount and maximumAmount in
  -h, --help              help for get
      --id string         the payment method id
      --locale string     get the payment method name in the corresponding language
```

### Options inherited from parent commands

```
  -c, --config string   specifies a custom config file to be used
  -d, --debug           enables debug logging information
  -m, --mode string     indicates the api target from test/live (default "test")
      --print-json      toggle the output type to json
  -t, --token string    the type of token to use for auth (default "MOLLIE_API_TOKEN")
  -v, --verbose         print verbose logging messages (defaults to false)
```

### SEE ALSO

* [mollie methods](mollie_methods.md)	 - All payment methods that Mollie offers and can be activated

###### Auto generated by spf13/cobra on 1-Nov-2020