load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/prksu/publr/cmd/service/sites",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/api/sites/v1alpha3:go_default_library",
        "//pkg/service:go_default_library",
        "//pkg/service/server/sites:go_default_library",
    ],
)

go_binary(
    name = "sites",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_image(
    name = "image",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
