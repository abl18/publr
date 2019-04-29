load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["client.go"],
    importpath = "github.com/prksu/publr/pkg/client/users",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/api/users/v1alpha1:go_default_library",
        "//pkg/server/users:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
    ],
)
