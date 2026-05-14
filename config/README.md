# Configuration files

The configuration files are name accordingly the environment where they'll we be used:

* `config.development.yaml`
* `config.production.yaml`
* `config.test.yaml`

Configuration file named `config.local.yaml` are merged with above files.

Example of configuration file:

```yaml
application:
  name: Materials Local
  log:
    level: Warning
```

Where:

* `name`: the applicaton/evironment name
* `log.level`: the log level, corresponding to log/slog Logger and OpenTelemetry
    * `Debug`
    * `Info`
    * `Warn` or `Warning`
    * `Error`