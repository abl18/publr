load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "internal.go",
        "sanitize.go",
        "traceinternals.go",
    ],
    importmap = "github.com/prksu/publr/vendor/go.opencensus.io/internal",
    importpath = "go.opencensus.io/internal",
    visibility = ["//visibility:public"],
    deps = ["//vendor/go.opencensus.io:go_default_library"],
)
