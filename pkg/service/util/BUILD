load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["page_token.go"],
    importpath = "github.com/prksu/publr/pkg/service/util",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/google.golang.org/grpc/codes:go_default_library",
        "//vendor/google.golang.org/grpc/status:go_default_library",
    ],
)
