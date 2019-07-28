load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "GoUnusedProtection__.go",
        "agent.go",
        "jaeger.go",
        "jaeger-consts.go",
    ],
    importmap = "github.com/prksu/publr/vendor/contrib.go.opencensus.io/exporter/jaeger/internal/gen-go/jaeger",
    importpath = "contrib.go.opencensus.io/exporter/jaeger/internal/gen-go/jaeger",
    visibility = ["//vendor/contrib.go.opencensus.io/exporter/jaeger:__subpackages__"],
    deps = ["//vendor/github.com/apache/thrift/lib/go/thrift:go_default_library"],
)
