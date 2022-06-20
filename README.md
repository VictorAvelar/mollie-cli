# mollie-cli

### :warning: This is a work in progress

mollie-cli provides a developer friendly way to interact with [Mollie's REST API](https://docs.mollie.com/reference/v2).

It also works as a good example of how to use the [mollie-api-go](https://github.com/VictorAvelar/mollie-api-go) sdk.

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

# Roadmap

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
