load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["serviceconfig.go"],
    importmap = "github.com/prksu/publr/vendor/google.golang.org/grpc/serviceconfig",
    importpath = "google.golang.org/grpc/serviceconfig",
    visibility = ["//visibility:public"],
    deps = ["//vendor/google.golang.org/grpc/internal:go_default_library"],
)
