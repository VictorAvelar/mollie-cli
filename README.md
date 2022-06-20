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

## Configuration

Copy the config file [`.mollie.yaml`](.mollie.yaml) to your home folder or your home folder config. If you are using this during development you can also create a copy in the directory where you execute mollie commands or specify a custom location using the `--config` flag.

### Example config

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
    payments:
      all:
        - "RESOURCE"
        - "ID"
        - "MODE"
        - "STATUS"
        - "CANCELABLE"
        - "AMOUNT"
        - "METHOD"
        - "DESCRIPTION"
        - "SEQUENCE"
        - "REMAINING"
        - "REFUNDED"
        - "CAPTURED"
        - "SETTLEMENT"
        - "APP_FEE"
        - "CREATED_AT"
        - "AUTHORIZED_AT"
        - "EXPIRES"
        - "PAID_AT"
        - "FAILED_AT"
        - "CANCELED_AT"
        - "CUSTOMER_ID"
        - "SETTLEMENT_ID"
        - "MANDATE_ID"
        - "SUBSCRIPTION_ID"
        - "ORDER_ID"
        - "REDIRECT"
        - "WEBHOOK"
        - "LOCALE"
        - "COUNTRY"
      printable:
        - "ID"
        - "MODE"
        - "STATUS"
        - "AMOUNT"
        - "METHOD"
        - "DESCRIPTION"
        - "COUNTRY"
    customers:
      all:
        - "RESOURCE"
        - "ID"
        - "MODE"
        - "NAME"
        - "EMAIL"
        - "LOCALE"
        - "METADATA"
        - "CREATED_AT"
      printable:
        - "ID"
        - "NAME"
        - "EMAIL"
        - "LOCALE"
        - "METADATA"
        - "CREATED_AT"
    methods:
      all:
        - "RESOURCE"
        - "ID"
        - "DESCRIPTION"
        - "ISSUERS"
        - "MIN_AMOUNT"
        - "MAX_AMOUNT"
        - "LOGO"
      printable:
        - "ID"
        - "DESCRIPTION"
        - "ISSUERS"
        - "MIN_AMOUNT"
        - "MAX_AMOUNT"
    permissions:
      all:
        - "RESOURCE"
        - "ID"
        - "DESCRIPTION"
        - "GRANTED"
      printable:
        - "ID"
        - "DESCRIPTION"
        - "GRANTED"
    profiles:
      all:
        - "RESOURCE"
        - "ID"
        - "MODE"
        - "NAME"
        - "WEBSITE"
        - "EMAIL"
        - "PHONE"
        - "CATEGORY_CODE"
        - "STATUS"
        - "REVIEW"
        - "CREATED_AT"
      printable:
        - "ID"
        - "NAME"
        - "WEBSITE"
        - "EMAIL"
        - "PHONE"
        - "CATEGORY_CODE"
        - "STATUS"
        - "REVIEW"
    captures:
      all:
        - "RESOURCE"
        - "ID"
        - "MODE"
        - "AMOUNT"
        - "SETTLEMENT_AMOUNT"
        - "PAYMENT_ID"
        - "SHIPMENT_ID"
        - "SETTLEMENT_ID"
        - "CREATED_AT"
      printable:
        - "ID"
        - "AMOUNT"
        - "SETTLEMENT_AMOUNT"
        - "CREATED_AT"
    chargebacks:
      all:
        - "RESOURCE"
        - "ID"
        - "AMOUNT"
        - "SETTLEMENT_AMOUNT"
        - "CREATED_AT"
        - "REVERSED_AT"
        - "PAYMENT_ID"
      printable:
        - "ID"
        - "PAYMENT_ID"
        - "AMOUNT"
        - "SETTLEMENT_AMOUNT"
        - "CREATED_AT"
        - "REVERSED_AT"
    refunds:
      all:
        - "RESOURCE"
        - "ID"
        - "AMOUNT"
        - "SETTLEMENT_ID"
        - "SETTLEMENT_AMOUNT"
        - "DESCRIPTION"
        - "METADATA"
        - "STATUS"
        - "PAYMENT_ID"
        - "ORDER_ID"
        - "CREATED_AT"
      printable:
        - "ID"
        - "AMOUNT"
        - "SETTLEMENT_AMOUNT"
        - "DESCRIPTION"
        - "STATUS"
    invoices:
      all:
        - "RESOURCE"
        - "ID"
        - "REFERENCE"
        - "VAT_NUMBER"
        - "STATUS"
        - "ISSUED_AT"
        - "PAID_AT"
        - "DUE_AT"
        - "NET_AMOUNT"
        - "VAT_AMOUNT"
        - "GROSS_AMOUNT"
      printable:
        - "ID"
        - "REFERENCE"
        - "VAT_NUMBER"
        - "STATUS"
        - "NET_AMOUNT"
        - "VAT_AMOUNT"
        - "GROSS_AMOUNT"
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
