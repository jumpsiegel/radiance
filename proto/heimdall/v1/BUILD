load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "heimdall",
    srcs = ["heimdall.pb.go"],
    importpath = "github.com/jumpsiegel/radiance/proto/heimdall/v1",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)
