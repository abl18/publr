load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["bundler.go"],
    importmap = "github.com/prksu/publr/vendor/google.golang.org/api/support/bundler",
    importpath = "google.golang.org/api/support/bundler",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/sync/semaphore:go_default_library"],
)
