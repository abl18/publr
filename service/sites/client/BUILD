load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["client.go"],
    importpath = "github.com/prksu/publr/service/sites/client",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/api/sites/v1alpha3:go_default_library",
        "//vendor/go.opencensus.io/plugin/ocgrpc:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/balancer/roundrobin:go_default_library",
    ],
)
