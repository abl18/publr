load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["stream_chunk.pb.go"],
    importmap = "github.com/prksu/publr/vendor/github.com/grpc-ecosystem/grpc-gateway/internal",
    importpath = "github.com/grpc-ecosystem/grpc-gateway/internal",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/any:go_default_library",
    ],
)
