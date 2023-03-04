load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "runtime",
    srcs = [
        "runtime.go",
        "serde.go",
    ],
    importpath = "github.com/jumpsiegel/radiance/pkg/runtime",
    visibility = ["//visibility:public"],
    deps = ["@com_github_gagliardetto_binary//:binary"],
)
