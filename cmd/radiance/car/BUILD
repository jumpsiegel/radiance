load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "car",
    srcs = ["car.go"],
    importpath = "github.com/jumpsiegel/radiance/cmd/radiance/car",
    visibility = ["//visibility:public"],
    deps = [
        "//cmd/radiance/car/create",
        "//cmd/radiance/car/dump",
        "@com_github_spf13_cobra//:cobra",
    ],
)
