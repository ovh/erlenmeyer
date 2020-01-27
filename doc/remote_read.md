# Prometheus Remote-read

Edit `prometheus.yml` configuration file, this file contains the global instance configuration.

You can find this file by greping the process which use it.

```sh
ps aux | grep prometheus | grep -v 'grep'
```

The process arg `--config.file` contains the configuration file path.

To setup prometheus remote configuration add the following lines (replace captitals strings by your tokens):

```yaml
remote_read:
  - url: http://127.0.0.1:8080/prometheus/remote_read
    basic_auth:
      username: ''
      password: 'READ_TOKEN'
```

Don't forget to restart your Prometheus instance to apply modifications.

## Going further

> [!warning]
>
> Any feedback on this implementation will be greatly welcomed, you can reach us on [gitter](https://gitter.im/ovh/metrics).
>