load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["jsonpb.go"],
    importmap = "github.com/prksu/publr/vendor/github.com/golang/protobuf/jsonpb",
    importpath = "github.com/golang/protobuf/jsonpb",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/struct:go_default_library",
    ],
)
