load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["logging.go"],
    importpath = "github.com/prksu/publr/pkg/service/logging",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/log:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/status:go_default_library",
    ],
)
