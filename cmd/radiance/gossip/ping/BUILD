load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ping",
    srcs = ["ping.go"],
    importpath = "github.com/jumpsiegel/radiance/cmd/radiance/gossip/ping",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/gossip",
        "@com_github_spf13_cobra//:cobra",
        "@io_k8s_klog_v2//:klog",
        "@org_golang_x_sync//errgroup",
    ],
)
