load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "client_mock.go",
        "datastore_mock.go",
    ],
    importpath = "github.com/prksu/publr/test/mock/sites",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/api/sites/v1alpha3:go_default_library",
        "//vendor/github.com/golang/mock/gomock:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/empty:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
    ],
)
