load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["status.go"],
    importmap = "github.com/prksu/publr/vendor/google.golang.org/grpc/status",
    importpath = "google.golang.org/grpc/status",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes:go_default_library",
        "//vendor/google.golang.org/genproto/googleapis/rpc/status:go_default_library",
        "//vendor/google.golang.org/grpc/codes:go_default_library",
        "//vendor/google.golang.org/grpc/internal:go_default_library",
    ],
)
