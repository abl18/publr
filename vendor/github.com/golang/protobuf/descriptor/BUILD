load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["descriptor.go"],
    importmap = "github.com/prksu/publr/vendor/github.com/golang/protobuf/descriptor",
    importpath = "github.com/golang/protobuf/descriptor",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/descriptor:go_default_library",
    ],
)
