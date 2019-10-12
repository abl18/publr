load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["mock.go"],
    importpath = "github.com/prksu/publr/pkg/service/mock/posts",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/api/posts/v1alpha3:go_default_library",
        "//vendor/github.com/golang/mock/gomock:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/empty:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
    ],
)
