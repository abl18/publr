load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["metric.go"],
    importmap = "github.com/prksu/publr/vendor/github.com/prometheus/client_golang/prometheus/internal",
    importpath = "github.com/prometheus/client_golang/prometheus/internal",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/prometheus/client_model/go:go_default_library"],
)
