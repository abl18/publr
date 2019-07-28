load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["client.go"],
    importpath = "github.com/prksu/publr/pkg/service/client/sites",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/api/sites/v1alpha2:go_default_library",
        "//pkg/service:go_default_library",
        "//vendor/go.opencensus.io/plugin/ocgrpc:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/balancer/roundrobin:go_default_library",
        "//vendor/google.golang.org/grpc/credentials:go_default_library",
    ],
)
