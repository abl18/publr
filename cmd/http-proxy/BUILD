load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/prksu/publr/cmd/http-proxy",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/api/posts/v1alpha3:go_default_library",
        "//pkg/api/sites/v1alpha3:go_default_library",
        "//pkg/api/users/v1alpha3:go_default_library",
        "//pkg/service:go_default_library",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/runtime:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/balancer/roundrobin:go_default_library",
        "//vendor/google.golang.org/grpc/credentials:go_default_library",
    ],
)

go_binary(
    name = "http-proxy",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_image(
    name = "image",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
