# PromQL

PromQL is a **Query Language** for Prometheus. It offers basic query capabilities, like OpenTSDB, plus a way to use operators between two series. All query documentation is available [here](https://prometheus.io/docs/prometheus/latest/querying/basics/){.external}

For example to retrieve your data, you can simply execute the following HTTP request using cURL.

```shell-session
curl 'http://u:READ_TOKEN@user:[READ_TOKEN]@127.0.0.1:8080/prometheus/api/v1/query_range?query=SERIES_NAME\{LABEL0_KEY="LABEL0_VALUE",LABEL1_KEY="LABEL1_VALUE"\}&start=1533127072.115&end=1533127472.115&step=2m'
```

where:

* READ_TOKEN is your read token available on your OVH Metrics account
* SERIES_NAME is the series name to retrieve. It can be for example **os.cpu**
* LABEL0_KEY and LABEL1_KEY are specific labels that the series must have with
* LABEL0_KEY values is LABEL0_VALUE and LABEL1_KEY value is LABEL1_VALUE.

## PromQL valid operators

### Binary operators

A few binary arithmetic operators exists in PromQL and can be used on the Metrics platform:

| Operator | Name | Supported |
|----------|------|-----------|
| + | addition | yes |
| - | subtraction | yes |
| * | multiplication | yes |
| / | division | yes |
| % | modulo | yes |
| ^ | power/exponentiation | yes |
| == | equal | yes |
| != | not-equal | yes |
| > | greater-than | yes |
| < | less-than | yes |
| >= | greater-or-equal | yes |
| <= | less-or-equal | yes |
| and | intersection | yes |
| or | union | yes |
| unless | complement | yes |

**Modulo** (%) and **exponentiation** (^) are **not yet supported** accross several metrics. They both can still be applied between metrics and scalar value.

For all operators, the same precedence applies than in promQL.

### Vector matching

PromQL has it's own method to manage how to match left and right entry of an operator. The valid vector matching on the Metrics platform are:

| Operator | Params | Supported |
|----------|------|-----------|
| ignoring | &lt;label list&gt; | yes |
| on | &lt;label list&gt; | yes |
| group_left | &lt;label list&gt; | yes |
| group_right | &lt;label list&gt; | yes |

Only one keyword between **on** and **ignoring** can be applied to an operator. The same applies to **group_left** or **group_right**

### Aggregation operators

The valid aggregation operators on the Metrics platform are:

| Aggregation pperator | Details | Supported |
|----------|------|-----------|
| sum | calculate sum over dimensions | yes |
| min | select minimum over dimensions | yes |
| max | select maximum over dimensions | yes |
| stddev | calculate the average over dimensions | yes |
| stdvar | calculate population standard deviation over dimensions | yes |
| count | count number of elements in the vector | yes |
| count_values | count number of elements with the same value | yes |
| bottomk | smallest k elements by sample value | yes |
| topk | largest k elements by sample value | yes |
| quantile | calculate φ-quantile (0 ≤ φ ≤ 1) over dimensions | yes |

The aggregation operator **count_values** expects a label string key for each new metrics created (per different values).

### Aggregation operators clauses

PromQL has it's own method to manage how to match left and right entry of an operator. The valid vector matching on the Metrics platform are:

| Operator | Params | Supported |
|----------|------|-----------|
| without | &lt;label list&gt; | yes |
| by | &lt;label list&gt; | yes |

Only one keyword between **without** and **by** can be applied to an aggregation operator.

### Supported PromQL functions

The valid promQL functions on the Metrics platform are:

| Function | Params | Supported |
|----------|------|-----------|
| abs | instant-vector | yes |
| absent | instant-vector | yes |
| ceil | instant-vector | yes |
| changes | range-vector | yes |
| clamp_max | instant-vector, scalar | yes |
| clamp_min | instant-vector, scalar | yes |
| count_scalar | instant-vector | yes |
| day_of_month | instant-vector | yes |
| day_of_week | instant-vector | yes |
| days_in_month | instant-vector | yes |
| delta | range-vector | yes |
| deriv | range-vector | no |
| drop_common_labels | instant-vector | yes |
| exp | instant-vector | yes |
| floor | instant-vector | yes |
| histogram_quantile | float, instant-vector | yes |
| holt_winters | range-vector, scalar, scalar| yes |
| hour | instant-vector | yes |
| idelta | range-vector | yes |
| increase | range-vector | yes |
| irate | range-vector | yes |
| label_join | instant-vector, string,  string, string, string, ... | yes |
| label_replace | instant-vector, string, string, string, string) | yes |
| ln | instant-vector | yes |
| log2 | instant-vector | yes |
| log10 | instant-vector | yes |
| minute | instant-vector | yes |
| month | instant-vector | yes |
| predict_linear | range-vector, scalar | yes |
| rate | range-vector | yes |
| resets | range-vector | yes |
| round | instant-vector, (optional) scalar | yes |
| scalar | instant-vector | yes |
| sort | instant-vector | yes |
| sort_desc | instant-vector | yes |
| sqrt | instant-vector | yes |
| time | | yes |
| timestamp | instant-vector | yes |
| vector | scalar | yes |
| year | instant-vector | yes |
| avg_over_time | range-vector | yes |
| min_over_time | range-vector | yes |
| max_over_time | range-vector | yes |
| sum_over_time | range-vector | yes |
| count_over_time | range-vector | yes |
| quantile_over_time | range-vector | yes |
| stddev_over_time | range-vector | yes |
| stdvar_over_time | range-vector | yes |

For the `timestamp` function, unary minus are not supported.

### PromQL examples queries

Here, you will find two valid queries examples.

For example this request will add 2 os.cpu series from host 1 and host 2 (we just replace the + character per it's URL encoded value %2B).

```shell-session
curl 'http://u:READ_TOKEN@127.0.0.1:8080/prometheus/api/v1/query_range?query=os.cpu\{host="1",cpu="1"\}%2Bos.cpu\{host="2",cpu="2"\}&start=1533127072.115&end=1533127472.115&step=2m'
```

Our second example here will compute the rate of the os.cpu metric for the host 1 and all cpu.

```shell-session
curl 'http://u:READ_TOKEN@127.0.0.1:8080/prometheus/api/v1/query_range?query=sum(rate(os.cpu\{host="1"\}\[1m\]))&start=1533127072.115&end=1533127472.115&step=2m'
```

To select a Time-series stored in Metrics with invalid Prometheus character as "-" you can also use the PromQL `{__name__="http-requests-total"}` syntax as Time series matcher expression.  Matchers other than = (!=, =~, !~) may also be used.

## Go further

> [!warning]
>
> Any feedback on this implementation will be greatly welcomed, you can reach us on [gitter](https://gitter.im/ovh/metrics).
>