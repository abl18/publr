load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["client.go"],
    importpath = "github.com/prksu/publr/pkg/client/sites",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/api/sites/v1alpha1:go_default_library",
        "//pkg/server/sites:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
    ],
)
