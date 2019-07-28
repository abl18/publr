load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "rpcz.go",
        "templates.go",
        "tracez.go",
        "zpages.go",
    ],
    importmap = "github.com/prksu/publr/vendor/go.opencensus.io/zpages",
    importpath = "go.opencensus.io/zpages",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opencensus.io/internal:go_default_library",
        "//vendor/go.opencensus.io/plugin/ocgrpc:go_default_library",
        "//vendor/go.opencensus.io/stats/view:go_default_library",
        "//vendor/go.opencensus.io/trace:go_default_library",
        "//vendor/go.opencensus.io/zpages/internal:go_default_library",
    ],
)
