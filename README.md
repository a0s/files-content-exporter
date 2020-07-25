files-content-exporter
======================

Export file(s) content as Prometheus metric

Settings
--------

* `FILES_CONTENT_EXPORTER_CONFIG_FILE_PATH` | default: _/config.yml_ - path to configuration yaml

* `FILES_CONTENT_EXPORTER_PORT` | default: _9457_ - port to bind

* `FILES_CONTENT_EXPORTER_HOST` | default: _127.0.0.1_  - host to bind


Usage
-----

Example of configuration to fetch heat metrics for Cubieboard2 (see `examples/config.yml`)

```yaml
path_as_label_enabled: true                                 # include path to file as `path` label
entities:                                                   # list of entities (one file - one metric)
  - file: /sys/devices/virtual/thermal/thermal_zone0/temp   # path to file with metric, *required*
    name: cpu_temp_celsius                                  # metric's name in export, *required*
    labels:                                                 # list of labels, optional
      thermal_zone: 0
    help: CPU thermal 0                                     # description of metric
  - file: /sys/devices/virtual/thermal/thermal_zone1/temp
    name: cpu_temp_celsius
    labels:
      thermal_zone: 1
    help: CPU thermal 1
  - file: /sys/devices/virtual/thermal/thermal_zone2/temp
    name: cpu_temp_celsius
    labels:
      thermal_zone: 2
    help: CPU thermal 2
  - file: /sys/devices/virtual/thermal/thermal_zone3/temp
    name: cpu_temp_celsius
    labels:
      thermal_zone: 3
    help: CPU thermal 3
  - file: /sys/power/axp_pmu/pmu/temp
    name: pmu_temp_celsius
```

Run as docker container

```sh
docker run \
  -v /sys:/sys \
  -v `pwd`/example/config.yml:/config.yml \
  -p 9457:9457 \
  -e FILES_CONTENT_EXPORTER_HOST=0.0.0.0 \
  a00s/files-content-exporter
```


