load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["struct.pb.go"],
    importmap = "github.com/prksu/publr/vendor/github.com/golang/protobuf/ptypes/struct",
    importpath = "github.com/golang/protobuf/ptypes/struct",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/golang/protobuf/proto:go_default_library"],
)
