load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "argument.go",
        "driver.go",
        "expectations.go",
        "expectations_before_go18.go",
        "expectations_go18.go",
        "options.go",
        "query.go",
        "result.go",
        "rows.go",
        "rows_go18.go",
        "sqlmock.go",
        "sqlmock_go18.go",
        "statement.go",
    ],
    importmap = "github.com/prksu/publr/vendor/github.com/DATA-DOG/go-sqlmock",
    importpath = "github.com/DATA-DOG/go-sqlmock",
    visibility = ["//visibility:public"],
)
