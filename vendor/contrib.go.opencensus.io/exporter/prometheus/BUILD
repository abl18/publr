load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "prometheus.go",
        "sanitize.go",
    ],
    importmap = "github.com/prksu/publr/vendor/contrib.go.opencensus.io/exporter/prometheus",
    importpath = "contrib.go.opencensus.io/exporter/prometheus",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
        "//vendor/github.com/prometheus/client_golang/prometheus/promhttp:go_default_library",
        "//vendor/go.opencensus.io/metric/metricdata:go_default_library",
        "//vendor/go.opencensus.io/metric/metricexport:go_default_library",
        "//vendor/go.opencensus.io/stats/view:go_default_library",
    ],
)
