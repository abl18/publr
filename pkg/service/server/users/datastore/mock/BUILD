load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["mock_datastore.go"],
    importpath = "github.com/prksu/publr/pkg/service/server/users/datastore/mock",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/api/users/v1alpha3:go_default_library",
        "//vendor/github.com/golang/mock/gomock:go_default_library",
    ],
)
