load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "gen.go",
        "resources.go",
    ],
    importmap = "github.com/prksu/publr/vendor/go.opencensus.io/zpages/internal",
    importpath = "go.opencensus.io/zpages/internal",
    visibility = ["//visibility:public"],
)
