load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cloudsql.go",
        "cloudsql_vm.go",
    ],
    importmap = "github.com/prksu/publr/vendor/google.golang.org/appengine/cloudsql",
    importpath = "google.golang.org/appengine/cloudsql",
    visibility = ["//visibility:public"],
)
