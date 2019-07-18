load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/prksu/publr/cmd/service/users",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/api/users/v1alpha2:go_default_library",
        "//pkg/log:go_default_library",
        "//pkg/service/logging:go_default_library",
        "//pkg/service/server/users:go_default_library",
        "//pkg/storage/database:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
    ],
)

go_binary(
    name = "users",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_image(
    name = "image",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
