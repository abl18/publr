load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "dbstats_go1.11.go",
        "driver.go",
        "driver_go1.10.go",
        "driver_go1.8.go",
        "driver_go1.9.go",
        "observability.go",
        "options.go",
    ],
    importmap = "github.com/prksu/publr/vendor/contrib.go.opencensus.io/integrations/ocsql",
    importpath = "contrib.go.opencensus.io/integrations/ocsql",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opencensus.io/stats:go_default_library",
        "//vendor/go.opencensus.io/stats/view:go_default_library",
        "//vendor/go.opencensus.io/tag:go_default_library",
        "//vendor/go.opencensus.io/trace:go_default_library",
    ],
)
