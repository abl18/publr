load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "logging.go",
        "service.go",
    ],
    importpath = "github.com/prksu/publr/pkg/service",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/log:go_default_library",
        "//vendor/contrib.go.opencensus.io/exporter/jaeger:go_default_library",
        "//vendor/contrib.go.opencensus.io/exporter/prometheus:go_default_library",
        "//vendor/go.opencensus.io/plugin/ocgrpc:go_default_library",
        "//vendor/go.opencensus.io/stats/view:go_default_library",
        "//vendor/go.opencensus.io/trace:go_default_library",
        "//vendor/go.opencensus.io/zpages:go_default_library",
        "//vendor/golang.org/x/net/http2:go_default_library",
        "//vendor/golang.org/x/net/http2/h2c:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/credentials:go_default_library",
        "//vendor/google.golang.org/grpc/status:go_default_library",
    ],
)
