load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["httpbody.pb.go"],
    importmap = "github.com/prksu/publr/vendor/google.golang.org/genproto/googleapis/api/httpbody",
    importpath = "google.golang.org/genproto/googleapis/api/httpbody",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/any:go_default_library",
    ],
)
