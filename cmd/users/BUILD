load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/prksu/publr/cmd/users",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/api/users/v1alpha3:go_default_library",
        "//pkg/service:go_default_library",
        "//service/users:go_default_library",
    ],
)

go_binary(
    name = "users",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_image")

go_image(
    name = "go_image",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

container_image(
    name = "image",
    base = ":go_image",
    stamp = True,
    ports = [
        "9000",
    ],
)
