# Graphite protocol

The graphite protocol is based on this [documentation](http://graphite-api.readthedocs.io/en/latest/api.html).

Available paths are:

* `/graphite/metrics`
* `/graphite/metrics/find`
* `/graphite/metrics/index.json`
* `/graphite/metrics/expand`
* `/graphite/render`

## Authentification

To query data to a Warp 10 instance, you will need a Warp 10 **READ TOKEN**. Use Basic Auth directly inside the URL to pass it properly, like this :

```cURL
http://user:[READ_TOKEN]@127.0.0.1:8080/graphite
```

## Get all available Geo Times Series

This section is dedicated to the following path `/graphite/metrics/index.json`.

The documentation of this path is available [here](http://graphite-api.readthedocs.io/en/latest/api.html#metrics-index-json).

Following parameters of the the route:

| Name  | Description                             | Type     | Default | Status |
| ----- | --------------------------------------- | -------- | ------- | ------ |
| jsonp | Wraps the response in a JSONP callback. | `string` | none    | ok     |

## Find Geo Time Series

This section is dedicated to following paths `/graphite/metrics` and `/graphite/metrics/find` which are the same fonctionnality.

The documentation of those paths are available [here](http://graphite-api.readthedocs.io/en/latest/api.html#metrics-find).

Following parameters of the the route:

| Name      | Description                                                 | Type      | Default    | Status                                  |
| --------- | ----------------------------------------------------------- | --------- | ---------- | --------------------------------------- |
| query     | The query to search for.                                    | `string`  | none       | ok                                      |
| format    | The output format to use. Can be `completer` or `treejson`. | `string`  | `treejson` | the `completer` format is not supported |
| wildcards | Whether to add a wildcard result at the end or no           | `integer` | `0`        | ok                                      |
| from      | Epoch timestamp from which to consider metrics.             | `date`    | none       | worded date is not supported            |
| until     | Epoch timestamp from which to consider metrics.             | `date`    | none       | worded date is not supported            |
| jsonp     | Wraps the response in a JSONP callback.                     | `string`  | none       | ok                                      |

## Expand Geo Times Series

This section is dedicated to the following path `/graphite/metrics/expand`.

The documentation of this path is available [here](http://graphite-api.readthedocs.io/en/latest/api.html#metrics-index-json).

Following parameters of the the route:

| Name        | Description                                                      | Type     | Default | Status |
| ----------- | ---------------------------------------------------------------- | -------- | ------- | ------ |
| query       | The metrics query. Can be specified multiple times.              | `string` | none    | ok     |
| groupByExpr | Whether to return a flat list of results or group them by query. | `int`    | 0       | ok     |
| leavesOnly  | Whether to only return leaves or both branches and leaves.       | `int`    | 0       | ok     |
| jsonp       | Wraps the response in a JSONP callback.                          | `string` | none    | ok     |

## Query Geo Times Series

This section is dedicated to the following path `/graphite/render`.

The documentation of this path is available [here](http://graphite-api.readthedocs.io/en/latest/api.html#the-render-api-render).

Following parameters of the the route:

| Name   | Description                                                                     | Type     | Default | Status |
| ------ | ------------------------------------------------------------------------------- | -------- | ------- | ------ |
| target | The query to search for.                                                        | `string` | none    | ok     |
| format | The output format to use. Can be `rickshaw`, `dygraph`, `json`, `csv` or `raw`. | `string` | `json`  | ok     |
| from   | Epoch timestamp from which to consider metrics.                                 | `date`   | none    | alpha  |
| until  | Epoch timestamp from which to consider metrics.                                 | `date`   | none    | alpha  |
| jsonp  | Wraps the response in a JSONP callback.                                         | `string` | none    | ok     |

### Functions

The documentation of graphite's functions is available [here](http://graphite-api.readthedocs.io/en/latest/functions.html).

## Go further

> [!warning]
>
> Any feedback on this implementation will be greatly welcomed, you can reach us on [gitter](https://gitter.im/ovh/metrics).
>