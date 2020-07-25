files-content-exporter
======================

Export file(s) content as Prometheus metric

Usage example
------------- 

files-content-exporter.yml:

```yaml
entities:
    - file: /sys/devices/virtual/thermal/thermal_zone0/temp
    name: cpu_temp_celsius
    labels:
      thermal_zone: 0
- file: /sys/devices/virtual/thermal/thermal_zone1/temp
  name: cpu_temp_celsius
  labels:
    thermal_zone: 1


  - file: /sys/power/axp_pmu/pmu/temp
    name: pmu_temp_celsius
path_as_label_enabled: true
```


