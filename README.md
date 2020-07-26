files-content-exporter
======================

Export file(s) content as Prometheus metric

Settings
--------

* `FILES_CONTENT_EXPORTER_CONFIG_FILE_PATH` | default: _/config.yml_ - path to config.yml

* `FILES_CONTENT_EXPORTER_PORT` | default: _9457_ - port to bind

* `FILES_CONTENT_EXPORTER_HOST` | default: _127.0.0.1_  - host to bind

* `FILES_CONTENT_EXPORTER_LOG_LEVEL` | default: _INFO_ - log level, one of `DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL`

config.yml
----------

This file describes metrics (entities) which will be export through `/metrics` endpoint. Example of config.yml (`examples/config.yml`) is valid for Cubieboard2. But you are free to use any files as a source of metrics.

> Warn! For entites with same names you *must* use equal help text and labels names. Otherwise you will get _"panic: a previously registered descriptor with the same fully-qualified name as Desc{...} has different label names or a different help string"_

```yaml
path_as_label_enabled: true                                 # [optional] use path to file with metric as `path` label
entities:                                                   # list of entities (one file - one metric)
  - file: /sys/devices/virtual/thermal/thermal_zone0/temp   # [required] path to file with metric
    name: cpu_temp_celsius                                  # [required] metric's name in export
    labels:                                                 # [optional] list of labels, optional
      thermal_zone: 0
    help: CPU thermal 0                                     # [optional] description of metric
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

Run `files-content-exporter` as Docker container
---------------------------------------------

```bash
docker run \
  -v /sys:/sys \
  -v `pwd`/examples/config.yml:/config.yml \
  -p 9457:9457 \
  -e FILES_CONTENT_EXPORTER_HOST=0.0.0.0 \
  a00s/files-content-exporter
```

Example response
----------------

```
# HELP cpu_temp_celsius CPU thermal
# TYPE cpu_temp_celsius gauge
cpu_temp_celsius{path="/sys/devices/virtual/thermal/thermal_zone0/temp",thermal_zone="0"} 45200
cpu_temp_celsius{path="/sys/devices/virtual/thermal/thermal_zone1/temp",thermal_zone="1"} 0
cpu_temp_celsius{path="/sys/devices/virtual/thermal/thermal_zone2/temp",thermal_zone="2"} 0
cpu_temp_celsius{path="/sys/devices/virtual/thermal/thermal_zone3/temp",thermal_zone="3"} 0
# HELP pmu_temp_celsius
# TYPE pmu_temp_celsius gauge
pmu_temp_celsius{path="/sys/power/axp_pmu/pmu/temp"} 44500
```

