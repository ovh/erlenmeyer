# Influx QL

InfluxDB has its own Query DSL, that mimics SQL without being plain ANSI SQL.

```text
 SELECT <field_key>[,<field_key>,<tag_key>] FROM <measurement_name>[,<measurement_name>]
```

## Authentification

To query data to a Warp 10 instance, you will need a Warp 10 **READ TOKEN**. Use Basic Auth directly inside the URL to pass it properly, like this :

```cURL
http://user:[READ_TOKEN]@127.0.0.1:8080/influxdb
```

## Data Exploration

The [InfluxQL data exploration statements](https://docs.influxdata.com/influxdb/v1.7/query_language/data_exploration/){.external} requests supported on the metrics platform are:

| Statement        | Supported                    |
| ---------------- | ---------------------------- |
| SELECT           | yes |
| WHERE            | yes |
| GROUP BY         | yes |
| INTO             | no |

We support also all possibility to configure a query result using: `ORDER BY time DESC`, `LIMIT`, `OFFSET` or `TIME` clauses.

You can also use all `InfluxQL tips on the query syntax`of the [InfluxQL data exploration page](https://docs.influxdata.com/influxdb/v1.7/query_language/data_exploration/){.external} on the Metrics Platform.

### Mathematical operators

The following existing [arithmetic operators](https://docs.influxdata.com/influxdb/v1.7/query_language/math_operators/){.external} in InfluxQL can be used on the Metrics platform:

| Operator | Name | Supported |
|----------|------|-----------|
| + | Addition | yes |
| - | Subtraction | yes |
| * | Multiplication | yes |
| / | Division | yes |
| % | Modulo | yes |
| & | Bitwise AND | yes |
| \| | Bitwise OR | yes |
| ^ | Bitwise Exclusive-OR | yes |

The **Modulo** (%) operator is **not yet supported** accross several metrics. It can still be applied between metrics and values.

All the **bitwise** operation: **Bitwise AND** (&), **Bitwise OR** (|) and **Bitwise Exclusive-OR** (^) are only working on numbers (float numbers will be automaticallly cast to integers).

### InfluxQL functions

The valid [influxQL functions](https://docs.influxdata.com/influxdb/v1.7/query_language/functions/) on the Metrics platform are (all parameter with a `?` are optionals):

| Function | Params | Supported |
|----------|------|-----------|
| COUNT | metrics | yes |
| DISTINCT | metrics | yes |
| INTEGRAL | metrics, (duration)?  | yes |
| MEAN | metrics | yes |
| MEDIAN | metrics | yes |
| MODE | metrics | yes |
| SPREAD | metrics | yes |
| STDDEV | metrics | yes |
| SUM | metrics | yes |
| BOTTOM | metrics, number | yes |
| FIRST | metrics | yes |
| LAST | metrics | yes |
| MAX | metrics | yes |
| MIN | metrics | yes |
| PERCENTILE | metrics, number | yes |
| SAMPLE | metrics, number | yes |
| TOP | metrics, number | yes |
| ABS | metrics | yes |
| ACOS | metrics | yes |
| ATAN | metrics | yes |
| ATAN2 | metrics, metrics | yes |
| CEIL | metrics | yes |
| COS | metrics | yes |
| CUMULATIVE_SUM | metrics | yes |
| DERIVATIVE | metrics, (duration)? | yes |
| DIFFERENCE | metrics | yes |
| ELAPSED | metrics, (duration)? | yes |
| EXP | metrics | yes |
| FLOOR | metrics | yes |
| HISTOGRAM | metrics | yes |
| LN | metrics, base | yes |
| LOG | metrics | yes |
| LOG2 | metrics | yes |
| LOG10 | metrics | yes |
| MOVING_AVERAGE | metrics, number | yes |
| NON_NEGATIVE_DERIVATIVE | metrics | yes |
| NON_NEGATIVE_DIFFERENCE | metrics | yes |
| POW | metrics, number | yes |
| ROUND | metrics | yes |
| SIN | metrics | yes |
| SQRT | metrics | yes |
| HOLT_WINTERS | metrics, duration, offset | no |
| CHANDE_MOMENTUM_OSCILLATOR | metrics, period, (hold_period)?, (warmup_type)? | no |
| EXPONENTIAL_MOVING_AVERAGE | metrics, period, (hold_period)?, (warmup_type)?  | no |
| DOUBLE_EXPONENTIAL_MOVING_AVERAGE | metrics, period, (hold_period)?, (warmup_type)?  | no |
| KAUFMANS_EFFICIENCY_RATIO | metrics, period, (hold_period)? | no |
| KAUFMANS_ADAPTIVE_MOVING_AVERAGE | metrics, period, (hold_period)? | no |
| TRIPLE_EXPONENTIAL_MOVING_AVERAGE | metrics, period, (hold_period)?, (warmup_type)? | no |
| TRIPLE_EXPONENTIAL_DERIVATIVE | metrics, period, (hold_period)?, (warmup_type)? | no |
| RELATIVE_STRENGTH_INDEX | metrics, period, (hold_period)?, (warmup_type)? | no |

### Data types and cast operations

The existing [data types and cast operations](https://docs.influxdata.com/influxdb/v1.7/query_language/data_exploration/#data-types-and-cast-operations){.external} of InfluxQL matches the one supported by the Metrics platform:

| Operator | Cast | Supported |
|----------|------|-----------|
| :: | integer | yes |
| :: | float | yes |
| :: | string | yes |

### Regular expressions

You can apply Regular expression on the Metrics Platform, however we **don't** support the native InfluxQL regular expression but [Warp 10™ native supported one](https://www.warp10.io/){.external}.

### GROUPBY clause

The existing [GROUPBY clause](https://docs.influxdata.com/influxdb/v1.7/query_language/data_exploration/#the-group-by-clause){.external} of InfluxQL is supported as if on the Metrics platform (by time or by tag fields).

The `null` parameter for filling will not provide any `null` values on the Metrics platform as NULL ticks correspond to empty values in Warp 10™ .

### WHERE clause

The existing [WHERE clause](https://docs.influxdata.com/influxdb/v1.7/query_language/data_exploration/#the-where-clause){.external} of InfluxQL is supported as if on the Metrics platform.

## Schema exploration statements

The existing [SHOW statements](https://docs.influxdata.com/influxdb/v1.7/query_language/schema_exploration/){.external} of InfluxQL supported by the Metrics platform are:

| Statement                   | Supported                    |
| --------------------------- | ---------------------------- |
| SHOW DATABASES              | yes |
| SHOW MEASUREMENTS           | yes |
| SHOW FIELD KEYS             | yes |
| SHOW RETENTION POLICIES     | yes |
| SHOW TAG KEYS               | yes |
| SHOW SERIES                 | yes |
| SHOW TAG VALUES             | yes |
| SHOW TAG VALUES CARDINALITY | yes |

As the concept of databases doesn't exists in Metrics, the `SHOW DATABASES` statement will always return only one database: `metrics`.

For the `SHOW TAG VALUES CARDINALITY` statement: no measurement split are computed and only the global tag cardinality is shown (compare to the same statement on InfluxDB). To get split tag cardinality statement, refers all wanted measurement in FROM clause.

## Database management statements

The existing [database management statements](https://docs.influxdata.com/influxdb/v1.7/query_language/database_management/){.external} of InfluxQL supported by the Metrics platform are:

| Statement        | Supported                    |
| ---------------- | ---------------------------- |
| CREATE DATABASE | yes |
| DROP DATABASE | no |
| DROP SERIES  | no |
| DELETE | no |
| DROP MEASUREMENT | no |
| DROP SHARD | no |
| CREATE RETENTION POLICY | no |
| ALTER RETENTION POLICY | no |
| DROP RETENTION POLICY | no |

As the `CREATE DATABASE` statement is used by some client, this statement was implemented in Metrics and always return. However no database exists in Metrics.

## Database continuous queries

The InfluxQL [continuous queries](https://docs.influxdata.com/influxdb/v1.7/query_language/continuous_queries/){.external} **can not** be performed **yet** on the Metrics platform.

## Query parameters

Natively an InfluxQL requests expects two parameter `q` and `db`. As in metrics application the database notion doesn't exist the `db` parameter is optional.
You can optionnaly add a `precision` with value **rfc3339** as parameter to change the time representation output (by default a timestamp to a Human readable UTC date).  

## Use InfluxQL to query data from sources that were not pushed on the Influx format

On Metrics, you can push data with several different format: for example the Prometheus. As by default when a user push `native` influxQL data, we add a "." as separator between the  `measurement` and its `field keys` in our internal representation. As for example, with Prometheus you can't have any "." in the data format. We added a new clause in where statements: the `_separator` to be able to query data from all kind of sources in InfluxQL.

The `_separator` allow the user to choose a custom selector which will splits its influx `measurement` from its `field keys`. This allow the user to use the promQL "_" as separator to split Prometheus metrics classnames.

Example:

```influxQL
SELECT mean("field") FROM "prometheus_data" WHERE  time >= now() - 6h AND _separator = "_" GROUP BY time(1h) fill(null)
```

This allow also the user to query InfluxData or any other kind of data with InfluxQL and to get the `raw` data representation.

Example:

```influxQL
SELECT mean("disk.used_percent") FROM "" WHERE  time >= now() - 6h AND _separator = "" GROUP BY time(1h) fill(null)
```

## Query using cURL

A quick example to use InfluxQL on Metrics with cURL would be:

```sh
curl --request GET \
  --url 'http://m:READ_TOKEN@127.0.0.1:8080/influxdb/query?q=SELECT%20%22used_percent%22%20FROM%20%22disk%22%20WHERE%20%20time%20%3E%3D%20now()%20-%2020m&=%20'
```

This will execute the following InfluxQL query:

```InfluxQL
SELECT "used_percent" FROM "disk" WHERE  time >= now() - 20m
```

For the one used to query Influx, you will notice that the `db` mandatory parameter of Influx is not set in this query. With Metrics the database field is optional, as Metrics does not rely on databases to store its metrics. If you need segmentation, you can use different Metrics project or isolate with an additional label.

## Go further

> [!warning]
>
> InfluxQL is a new release. Any feedback on this implementation will be greatly welcomed, you can reach us on [gitter](https://gitter.im/ovh/metrics).
>