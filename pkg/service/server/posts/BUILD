load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["server.go"],
    importpath = "github.com/prksu/publr/pkg/service/server/posts",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/api/posts/v1alpha2:go_default_library",
        "//pkg/service/server/posts/datastore:go_default_library",
        "//pkg/service/util:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/empty:go_default_library",
        "//vendor/google.golang.org/grpc/codes:go_default_library",
        "//vendor/google.golang.org/grpc/status:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["server_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//pkg/api/posts/v1alpha2:go_default_library",
        "//pkg/service/server/posts/datastore/mock:go_default_library",
        "//pkg/service/util:go_default_library",
        "//vendor/github.com/golang/mock/gomock:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/empty:go_default_library",
        "//vendor/google.golang.org/grpc/codes:go_default_library",
        "//vendor/google.golang.org/grpc/status:go_default_library",
    ],
)
