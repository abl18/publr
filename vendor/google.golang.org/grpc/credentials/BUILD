load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "credentials.go",
        "tls13.go",
    ],
    importmap = "github.com/prksu/publr/vendor/google.golang.org/grpc/credentials",
    importpath = "google.golang.org/grpc/credentials",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/google.golang.org/grpc/credentials/internal:go_default_library",
    ],
)
