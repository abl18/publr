load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["database.go"],
    importpath = "github.com/prksu/publr/pkg/storage/database",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/log:go_default_library",
        "//vendor/contrib.go.opencensus.io/integrations/ocsql:go_default_library",
        "//vendor/github.com/go-sql-driver/mysql:go_default_library",
    ],
)
