load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["internal.go"],
    importmap = "github.com/prksu/publr/vendor/google.golang.org/grpc/internal",
    importpath = "google.golang.org/grpc/internal",
    visibility = ["//visibility:public"],
    deps = ["//vendor/google.golang.org/grpc/connectivity:go_default_library"],
)
