load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["h2c.go"],
    importmap = "github.com/prksu/publr/vendor/golang.org/x/net/http2/h2c",
    importpath = "golang.org/x/net/http2/h2c",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/net/http/httpguts:go_default_library",
        "//vendor/golang.org/x/net/http2:go_default_library",
        "//vendor/golang.org/x/net/http2/hpack:go_default_library",
    ],
)
