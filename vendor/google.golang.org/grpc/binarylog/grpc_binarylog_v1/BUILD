load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["binarylog.pb.go"],
    importmap = "github.com/prksu/publr/vendor/google.golang.org/grpc/binarylog/grpc_binarylog_v1",
    importpath = "google.golang.org/grpc/binarylog/grpc_binarylog_v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/duration:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/timestamp:go_default_library",
    ],
)
