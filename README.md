# mollie-cli

## :warning: This is a work in progress

mollie-cli provides a developer friendly way to interact with [Mollie's REST API](https://docs.mollie.com/reference/v2).

It also works as a good example of how to use the [mollie-api-go](https://github.com/VictorAvelar/mollie-api-go) sdk.

## CLI Documentation

Generated automatically using cobra. [Read the command docs](docs/mollie.md)

## Installation

You can install this cli using the go toolchain.

```
go install github.com/VictorAvelar/mollie-cli/cmd/mollie@latest
```

### Building from source

```
git clone git@github.com:VictorAvelar/mollie-cli.git
```

The compile the library yourself, there are a couple of make commands to support this:

If you want to compile your current branch:

```
make compile-current
```

If you want to compile master:

```
make compile-master
```

## Configuration

Copy the config file [`.mollie.yaml`](.mollie.yaml) to your home folder or your home folder config. If you are using this during development you can also create a copy in the directory where you execute mollie commands or specify a custom location using the `--config` flag.

You can also run:

```bash
wget https://raw.githubusercontent.com/VictorAvelar/mollie-cli/master/.mollie.yaml
```

## Example config

```yaml
mollie:
  core:
    # Print JSON togeth er with the standard tab formatted format.
    json: false
    # Enables verbose logging during the command execution.
    verbose: false
    # Enables printing the curl representation for the request
    # performed, useful for importing to other tools like postman,
    # insomnia, etc...
    curl: false
    # For some features it is necessary to specify the type of
    # actions being performed especially when using organization
    # tokens. Accepted values are test/live.
    mode: test
    # Debug prints the report caller on log entries.
    debug: true
    # Set a custom path to parse config, by default the CLI will
    # attempt to parse from ~, ~/.config, and the current working
    # directory (pwd).
    config:
      default: true
      custom_path: "."
  # Fields are each of the data points returned as part of a response
  # all contains all the possible printable values and printable must
  # contain the values you want to print by default, printable can be
  # overwritten by using the persistent `-f` flag.
  fields:
  # ... field map definition
```

## Roadmap

## Authentication

- [x] API token authentication
- [x] Organization token authentication
- [x] Custom env variable authentication
- [ ] Mollie connect OAuth2

## Resources

- [x] Payments
- [x] Methods
- [x] Refunds
- [x] Chargebacks
- [x] Captures
- [ ] Orders
- [ ] Shipments
- [x] Customers
- [ ] Mandates
- [ ] Subscriptions
- [ ] Connect
- [x] Permissions
- [ ] Organizations
- [x] Profiles
- [ ] Onboarding
- [ ] Settlements
- [x] Invoices
- [ ] Miscellaneous

## Utilities

- [x] Browse - Opens Mollie related resources on a web browser.
