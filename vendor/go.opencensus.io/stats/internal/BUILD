load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["record.go"],
    importmap = "github.com/prksu/publr/vendor/go.opencensus.io/stats/internal",
    importpath = "go.opencensus.io/stats/internal",
    visibility = ["//visibility:public"],
    deps = ["//vendor/go.opencensus.io/tag:go_default_library"],
)
