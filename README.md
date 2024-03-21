# Shout

![Shout Splash Logo](https://static.caffeineforcode.com/shout-logo-splash-dark.png)

## Description

Shout is a versatile command-line utility designed to enhance CI/CD pipelines with efficient logging and alerting capabilities. 
It allows developers and DevOps professionals to log information, warnings, and errors directly from their pipelines and send customizable alerts to specified webhook URLs. 
With support for structured logging and color-coded outputs, Shout makes monitoring and debugging CI/CD processes more intuitive and accessible.


## Features

- **Logging**: Easily log messages, warnings, and errors with color-coded outputs to distinguish between log levels.
- **Webhook Alerts**: Send customizable alerts to any webhook URL, allowing integration with various monitoring and messaging platforms.

## Installation

To install Shout, clone the repository and build the application using Go:

```bash
git clone https://github.com/bacchusjackson/shout
cd shout
go build -o shout .
```

Ensure you have Go installed on your system to compile the application.

## Usage

Shout is designed with simplicity in mind, providing a straightforward CLI interface for all its features. Here are some examples of how to use Shout:

![demo](https://static.caffeineforcode.com/shout-demo.gif)

### Logging Information, Warnings, and Errors

Log an information message:

```bash
./shout log 'Deployment started'
```

Log a warning:

```bash
./shout warn 'Deployment taking longer than expected'
```

Log an error:

```bash
./shout error 'Deployment failed'
```

### Sending Alerts

Send an alert to a webhook URL:

```bash
./shout alert [WEBHOOK URL] key1=value1 key2=value2
```

### Displaying Version Information

To display the version and build information of Shout:

```bash
./shout version
```

## Contributing

We welcome contributions! If you'd like to contribute, please fork the repository and use a feature branch. Pull requests are warmly welcome.

## License

Apache 2.0 Open Source
