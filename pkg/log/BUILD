load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "log.go",
        "logger.go",
        "util.go",
    ],
    importpath = "github.com/prksu/publr/pkg/log",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/sirupsen/logrus:go_default_library",
        "//vendor/google.golang.org/grpc/codes:go_default_library",
    ],
)
