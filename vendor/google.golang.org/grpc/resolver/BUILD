load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["resolver.go"],
    importmap = "github.com/prksu/publr/vendor/google.golang.org/grpc/resolver",
    importpath = "google.golang.org/grpc/resolver",
    visibility = ["//visibility:public"],
    deps = ["//vendor/google.golang.org/grpc/serviceconfig:go_default_library"],
)
