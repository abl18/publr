load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["wrappers.pb.go"],
    importmap = "github.com/prksu/publr/vendor/github.com/golang/protobuf/ptypes/wrappers",
    importpath = "github.com/golang/protobuf/ptypes/wrappers",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/golang/protobuf/proto:go_default_library"],
)
