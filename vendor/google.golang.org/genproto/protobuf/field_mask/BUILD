load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["field_mask.pb.go"],
    importmap = "github.com/prksu/publr/vendor/google.golang.org/genproto/protobuf/field_mask",
    importpath = "google.golang.org/genproto/protobuf/field_mask",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/golang/protobuf/proto:go_default_library"],
)
