load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "export.go",
        "reader.go",
    ],
    importmap = "github.com/prksu/publr/vendor/go.opencensus.io/metric/metricexport",
    importpath = "go.opencensus.io/metric/metricexport",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opencensus.io/metric/metricdata:go_default_library",
        "//vendor/go.opencensus.io/metric/metricproducer:go_default_library",
        "//vendor/go.opencensus.io/trace:go_default_library",
    ],
)
