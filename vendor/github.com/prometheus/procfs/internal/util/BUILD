load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "parse.go",
        "sysreadfile.go",
        "sysreadfile_compat.go",
        "valueparser.go",
    ],
    importmap = "github.com/prksu/publr/vendor/github.com/prometheus/procfs/internal/util",
    importpath = "github.com/prometheus/procfs/internal/util",
    visibility = ["//vendor/github.com/prometheus/procfs:__subpackages__"],
)
