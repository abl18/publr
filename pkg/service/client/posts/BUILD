load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["client.go"],
    importpath = "github.com/prksu/publr/pkg/service/client/posts",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/api/posts/v1alpha2:go_default_library",
        "//pkg/service/server/posts:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
    ],
)
