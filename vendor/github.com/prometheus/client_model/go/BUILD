load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["metrics.pb.go"],
    importmap = "github.com/prksu/publr/vendor/github.com/prometheus/client_model/go",
    importpath = "github.com/prometheus/client_model/go",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/golang/protobuf/proto:go_default_library"],
)
