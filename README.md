# Erlenmeyer: Time Series query translator

Erlenmeyer is a Go Proxy used to parse multiple Open Source TimeSeries DataBase query (OpenTSDB, PromQL, Prometheus-remote_read, InfluxQL and Graphite) . Then they are translated into WarpScript to produce native [Warp 10](https://warp10.io/) queries.

![Erlenmeyer: Time Series query translator](./assets/logo.png)


## Building erlenmeyer

You can test and run locally erlenemeyer following those [building steps](./doc/BUILDING.md).

## Configuration

You can retrieve a `config.sample.yml` file, that can be re-used to configure erlenmeyer. 

```sh
cp config.sample.yml /Path/to/erlenmeyer.yaml
```

In this config file, you will retrieve the Warp10 backend endpoint to set `warp_endpoint`. Erlenmeyer will use this endpoint to resolve WarpScript generated queries. 

## Run

Run the dev compiled version:

```sh
./build/erlenmeyer --config /Path/to/erlenmeyer.yaml
```

Erlenmeyer supports some flags as `--listen` to specify erlenmeyer listen address and `--config` to specify the config file to use. 

More information about the supported flags are provided by executing: 

```sh
./build/erlenmeyer -h
```

## Supported protocols

| Name       | State   | Documentation |
| ---------- | ------- | ------------- |
| Warp10     | Native  | [doc](https://warp10.io/doc/reference) |
| PromQL     | Near full | [doc](./doc/promql.md) |
| Prometheus - Remote-read | Near full | [doc](./doc/remote_read.md) |
| OpenTSDB   | Near full | [doc](./doc/openTSDB.md) |
| Graphite   | Partial | [doc](./doc/graphite.md) |
| InfluxQL   | New | [doc](./doc/influxql.md) |

## Motivation

At [OVHcloud](https://github.com/ovh), a lot of Metrics users were used to a previous TSDB experience with which they felt confortable and efficient. The goal of Erlenmeyer was to leverage their existing habits, while converging on the Metrics Time Series platform. Hence we decided to welcome customers adding more and more protocols from several Open Source Time Series DB. Under the hood, the Warp10 platform offers the powerful WarpScript query language which was a great help on this process, by just implementing all the query layers as a transpilation step. The best part of it, is now you can have: the same backend, the same data and some Grafana Dashboard written in PromQL when others are using OpenTSDB!

## PromQL configuration

Warp10 metrics names can include some characters not available with Prometheus. Erlenmeyer can filter result metrics meta to replace for example `.` by `_`. 
To activate it, you will need to set some of the following configuration keys:

```yaml
# Enable filter on remote read enpoint
prometheus.remote_read.meta.replace.enabled: true
# Replace all series meta "." by "_"
prometheus.remote_read.meta.replace.map: 
  ".": "_"

# Enable filter on query enpoint only on series classnames
prometheus.query.classname.replace.enabled: true
# Replace all series classname "." by "_"
prometheus.query.classname.replace.map: 
  ".": "_"

# Enable filter on query enpoint only on series labels
prometheus.query.labels.replace.enabled: true
# Replace all series classname "." by "_"
prometheus.query.labels.replace.map: 
  ".": "_"
```

## Status

Erlenmeyer is used in production.

## Contributing

Instructions on how to contribute to Erlenmeyer are available on the [Contributing](./CONTRIBUTING.md) page.

## Metrics

Erlenmeyer exposes metrics about his usage:

| name                               | labels                  | type    | description                           |
| ---------------------------------- | ----------------------- | ------- | ------------------------------------- |
| erlenmeyer_exec_request            | app, token_id, protocol | counter | Warp execution count                  |
| erlenmeyer_exec_fetched_datapoints | app, token_id, protocol | counter | Number of datapoints fetched          |
| erlenmeyer_exec_ops                | app, token_id, protocol | counter | Number of WarpScript operations       |
| erlenmeyer_exec_error_request      | app, token_id, protocol | counter | Warp 10 error by user application     |
| erlenmeyer_http_request            |                         | counter | Number of http request handled        |
| erlenmeyer_http_error_request      |                         | counter | Number of http request in error       |
| erlenmeyer_http_status_code        | status                  | counter | Counter per requests status code      |
| erlenmeyer_http_response_time      | path                    | counter | Requests response time in nanoseconds |
| erlenmeyer_protocol_request        | protocol                | counter | Requests by protocol                  |
| erlenmeyer_protocol_error_request  | protocol, status        | counter | Requests error by protocol and status |
| erlenmeyer_graphite_function       | function                | counter | Function used by user of graphite     |
| erlenmeyer_influxdb_request        | function                | counter | Number of requests handled            |
| erlenmeyer_influxdb_errors         | function                | counter | Number of requests in errors          |
| erlenmeyer_influxdb_warning        | function                | counter | Number of errored client requests     |
| erlenmeyer_opentsdb_request        | function                | counter | Number of requests handled            |
| erlenmeyer_opentsdb_errors         | function                | counter | Number of requests in errors          |
| erlenmeyer_opentsdb_warning        | function                | counter | Number of errored client requests     |
| erlenmeyer_promql_request          | function                | counter | Number of requests handled            |
| erlenmeyer_promql_request          | function                | counter | Number of requests handled            |

You can enable a basic auth for the `/metrics` endpoint by adding the following keys in the config file:

```yaml
metrics.basicauth.enabled: true
metrics.basicauth.user: test
metrics.basicauth.password: $2b$12$hNf2lSsxfm0.i4a.1kVpSOVyBCfIB51VRjgBUyv6kdnyTlgWj81Ay 
```

The metrics basic auth password needs to be hashed. To do it you can follow the same step as in [prometheus basic auth doc](https://prometheus.io/docs/guides/basic-auth/): 

Create your own local `gen-pass.py` pyhton file: 

```python
import getpass
import bcrypt

password = getpass.getpass("password: ")
hashed_password = bcrypt.hashpw(password.encode("utf-8"), bcrypt.gensalt())
print(hashed_password.decode())
```

Save it and run it, this should prompt you the hashed password to use for the configuration file.

## Licence

Erlenmeyer is released under a [3-BSD clause license](./LICENSE).

## Get in touch

- Gitter: [metrics](https://gitter.im/ovh/metrics)
