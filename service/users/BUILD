load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "datastore.go",
        "server.go",
    ],
    importpath = "github.com/prksu/publr/service/users",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/api/users/v1alpha3:go_default_library",
        "//pkg/storage/database:go_default_library",
        "//pkg/util:go_default_library",
        "//vendor/github.com/go-sql-driver/mysql:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/empty:go_default_library",
        "//vendor/google.golang.org/grpc/codes:go_default_library",
        "//vendor/google.golang.org/grpc/status:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "datastore_test.go",
        "server_test.go",
    ],
    embed = [":go_default_library"],
    deps = [
        "//pkg/api/users/v1alpha3:go_default_library",
        "//pkg/util:go_default_library",
        "//test/mock/users:go_default_library",
        "//vendor/github.com/DATA-DOG/go-sqlmock:go_default_library",
        "//vendor/github.com/golang/mock/gomock:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/empty:go_default_library",
        "//vendor/google.golang.org/grpc/codes:go_default_library",
        "//vendor/google.golang.org/grpc/status:go_default_library",
    ],
)
