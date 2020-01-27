# OpenTSDB

## Authentification

To query data to a Warp 10 instance, you will need a Warp 10 **READ TOKEN**. Use Basic Auth directly inside the URL to pass it properly, like this :

```cURL
http://user:[READ_TOKEN]@127.0.0.1:8080/opentsdb
```

## Query using curl

Now let's retrieve the previously pushed data.

The full documentation is available at [http://opentsdb.net/docs/build/html/api_http/query/index.html](http://opentsdb.net/docs/build/html/api_http/query/index.html){.external}

Let's write a `query.json` file which contains the following code:

```json
{
 "start": 1346846000000,
 "end": 1346847300005,
 "queries": [
  {
   "metric": "sys.cpu.nice",
   "aggregator": "min",
   "downsample": "4m-avg",
   "tags": {
    "host": "*",
    "dc": "*"
   }
  }
 ]
}
```

This will get all the saved points and compute the query before returning the result. The curl command to execute this query is:

```shell-session
curl --data-binary @query.json 'http://user:[READ_TOKEN]@127.0.0.1:8080/opentsdb/api/query'
```

You should expects a result similar to:

```json
[{"metric":"sys.cpu.nice","tags":{"dce":"lga", "host": "web02"},"aggregateTags":["host","de"],
"dps": {"1346846340" :9, "1346846580" :9, "1346846820" :8, "1346847060" :8.5, "1346847300" :8.5}},{"m
etric":"sys.cpu.nice","tags":{"dce":"lga","host":"web01"},"aggregateTags":["host","dc"], "dps
":{"13468463460" :18, "1346846580" :18, "1346846820" :19, "1346847060" :19.5, "1346847300" :19.5}}]
```

## Supported queries attributes

The Metrics platform offers almost a full support for OpenTSDB 2.3 queries.

### OpenTSDB requests

The OpenTSDB requests attributes supported on the metrics platform are:

| Attribute         | Type            | Supported                    |
| ----------------- | --------------- | ---------------------------- |
| start             | Integer, String | yes |
| end               | Integer, String | yes |
| queries           | Array           | yes |
| noAnnotations     | Boolean         | no |
| globalAnnotations | Boolean         | no |
| msResolution      | Boolean         | yes |
| showTSUIDs        | Boolean         | no |
| showSummary (2.2) | Boolean         | no |
| showStats (2.2)   | Boolean         | no |
| showQuery (2.2)   | Boolean         | no |
| delete            | Boolean         | yes |
| timezone (2.3)    | String          | no |
| useCalendar (2.3) | Boolean         | no |

We **do not support** annotations (as in Metrics annotations can be stored in a series). `showTSUIDs` isn't implemented as our series are stored using an Hash of their classnames and tags.

The allowed strings date format are defined at [http://opentsdb.net/docs/build/html/user_guide/query/dates.html](http://opentsdb.net/docs/build/html/user_guide/query/dates.html){.external}.

### OpenTSDB sub-queries

The OpenTSDB sub-queries attributes supported on the metrics platform are:

| Attribute          | Type    | Supported                    |
| ------------------ | ------- | ---------------------------- |
| aggregator         | String  | yes |
| metric             | String  | yes |
| rate               | Boolean | yes |
| rateOptions        | Map     | yes |
| downsample         | String  | yes |
| tags               | Map     | yes |
| filters (2.2)      | List    | yes |
| explicitTags (2.3) | Boolean | yes |
| percentiles (2.4)  | Boolean | no |

Settings **explicitTags** will result only on the series that have all theirs labels key in tags map and/or in filters list.

### OpenTSDB rate-options

The OpenTSDB rate-options attributes supported on the metrics platform are:

| Attribute  | Type    | Supported                    |
| ---------- | ------- | ---------------------------- |
| counter    | Boolean | yes |
| counterMax | Integer | yes |
| resetValue | Integer | yes |
| dropResets | Boolean | yes |

### OpenTSDB Filters

The OpenTSDB Filters attributes supported on the metrics platform are:

| Attribute | Type    | Supported                    |
| --------- | ------- | ---------------------------- |
| type      | String  | yes |
| tagk      | String  | yes |
| filter    | String  | yes |
| groupBy   | Boolean | yes |

### OpenTSDB Aggregators

The OpenTSDB aggregators supported on the metrics platform are:

| Attribute | Interpolation             | Grouping/Downsampling | Supported                    |
| --------- | ------------------------- | --------------------- | ---------------------------- |
| avg       | Linear Interpolation      | Both                  | yes |
| count     | Not counted when missing  | Both                  | yes |
| dev       | Linear Interpolation      | Both                  | yes |
| ep50r3    | Linear Interpolation      | None                  | no |
| ep50r7    | Linear Interpolation      | None                  | no |
| ep75r3    | Linear Interpolation      | None                  | no |
| ep75r7    | Linear Interpolation      | None                  | no |
| ep90r3    | Linear Interpolation      | None                  | no |
| ep90r7    | Linear Interpolation      | None                  | no |
| ep95r3    | Linear Interpolation      | None                  | no |
| ep95r7    | Linear Interpolation      | None                  | no |
| ep99r3    | Linear Interpolation      | None                  | no |
| ep99r7    | Linear Interpolation      | None                  | no |
| ep999r3   | Linear Interpolation      | None                  | no |
| ep999r7   | Linear Interpolation      | None                  | no |
| first     | None                      | Downsampling          | yes |
| last      | None                      | Downsampling          | yes |
| mimmin    | Not compared when missing | Both                  | yes |
| mimmax    | Not compared when missing | Both                  | yes |
| min       | Linear Interpolation      | Both                  | yes |
| max       | Linear Interpolation      | Both                  | yes |
| none      | Not counted when missing  | Grouping              | yes |
| p50       | Linear Interpolation      | Both                  | yes |
| p75       | Linear Interpolation      | Both                  | yes |
| p90       | Linear Interpolation      | Both                  | yes |
| p95       | Linear Interpolation      | Both                  | yes |
| p99       | Linear Interpolation      | Both                  | yes |
| p999      | Linear Interpolation      | Both                  | yes |
| sum       | Linear Interpolation      | Both                  | yes |
| zimsum    | Zero when missing         | Both                  | yes |

### OpenTSDB Downsampling fill policies

The OpenTSDB downsampling fill policies supported on the metrics platform are:

| Policy | Supported                    |
| ------ | ---------------------------- |
| None   | yes |
| NaN    | yes |
| Null   | yes |
| Zero   | yes |

## Go further

> [!warning]
>
> Any feedback on this implementation will be greatly welcomed, you can reach us on [gitter](https://gitter.im/ovh/metrics).
>
