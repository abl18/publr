load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "annotations.pb.go",
        "client.pb.go",
        "field_behavior.pb.go",
        "http.pb.go",
        "resource.pb.go",
    ],
    importmap = "github.com/prksu/publr/vendor/google.golang.org/genproto/googleapis/api/annotations",
    importpath = "google.golang.org/genproto/googleapis/api/annotations",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/descriptor:go_default_library",
    ],
)
