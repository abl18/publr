load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "agent.go",
        "jaeger.go",
    ],
    importmap = "github.com/prksu/publr/vendor/contrib.go.opencensus.io/exporter/jaeger",
    importpath = "contrib.go.opencensus.io/exporter/jaeger",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/contrib.go.opencensus.io/exporter/jaeger/internal/gen-go/jaeger:go_default_library",
        "//vendor/github.com/apache/thrift/lib/go/thrift:go_default_library",
        "//vendor/go.opencensus.io/trace:go_default_library",
        "//vendor/google.golang.org/api/support/bundler:go_default_library",
    ],
)
