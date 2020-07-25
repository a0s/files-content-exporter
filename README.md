files-content-exporter
======================

Export file(s) content as Prometheus metric

Usage example
------------- 

`examples/config.yml`

```yaml
path_as_label_enabled: true
entities:
  - file: /sys/devices/virtual/thermal/thermal_zone0/temp
    name: cpu_temp_celsius
    labels:
      thermal_zone: 0
  - file: /sys/devices/virtual/thermal/thermal_zone1/temp
    name: cpu_temp_celsius
    labels:
      thermal_zone: 1
  - file: /sys/devices/virtual/thermal/thermal_zone2/temp
    name: cpu_temp_celsius
    labels:
      thermal_zone: 2
  - file: /sys/devices/virtual/thermal/thermal_zone3/temp
    name: cpu_temp_celsius
    labels:
      thermal_zone: 3
  - file: /sys/power/axp_pmu/pmu/temp
    name: pmu_temp_celsius
```
