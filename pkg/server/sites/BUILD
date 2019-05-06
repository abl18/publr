load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "datastore.go",
        "server.go",
    ],
    importpath = "github.com/prksu/publr/pkg/server/sites",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/api/sites/v1alpha1:go_default_library",
        "//pkg/storage/database:go_default_library",
        "//vendor/github.com/go-sql-driver/mysql:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/empty:go_default_library",
        "//vendor/google.golang.org/grpc/codes:go_default_library",
        "//vendor/google.golang.org/grpc/status:go_default_library",
    ],
)
