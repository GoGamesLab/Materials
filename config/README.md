# Configuration files

## Configuration files

The configuration files are name accordingly the environment (`ENV` variable) where they'll we be used:

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

* `application.name`: the applicaton/evironment name, default: "Grind"
* `application.log.level`: the log level, corresponding to log/slog Logger and OpenTelemetry, default: "Info"
    * `Debug`
    * `Info`
    * `Warn` or `Warning`
    * `Error`


## Environment variables

All config file definitions (except `ENV`) has a corresponding environment variable.

### `ENV`

The `ENV` environment variable is read and a corresponding config file is loaded. The default value is "production".

For example, if `ENV` contains "home", a config file named "config.home.yaml" will be located and loaded if found, and if not a warning is issued and the program proceed.

If a file named `.env` exists and contains a variable `ENV`, the value in this variable will be used.

### `APP_NAME`

The `APP_NAME` environment variable corresponds to the `application.name` config file value.

### `LOG_LEVEL`

The `LOG_LEVEL` environment variable corresponds to the `application.log.level` config file value.
