load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "grocksdb",
    srcs = [
        "array.go",
        "backup.go",
        "c.h",
        "cache.go",
        "cf_handle.go",
        "cf_metadata.go",
        "checkpoint.go",
        "compaction_filter.go",
        "comparator.go",
        "cow.go",
        "cuckoo_table.go",
        "db.go",
        "dbpath.go",
        "doc.go",
        "env.go",
        "filter_policy.go",
        "grocksdb.c",
        "grocksdb.h",
        "iterator.go",
        "jemalloc.go",
        "mem_alloc.go",
        "memory_usage.go",
        "merge_operator.go",
        "non_builtin.go",
        "optimistic_transaction_db.go",
        "options.go",
        "options_backup.go",
        "options_block_based_table.go",
        "options_compaction.go",
        "options_compression.go",
        "options_env.go",
        "options_flush.go",
        "options_ingest.go",
        "options_read.go",
        "options_restore.go",
        "options_transaction.go",
        "options_transactiondb.go",
        "options_write.go",
        "perf_context.go",
        "perf_level.go",
        "ratelimiter.go",
        "slice.go",
        "slice_transform.go",
        "snapshot.go",
        "sst_file_writer.go",
        "transaction.go",
        "transactiondb.go",
        "util.go",
        "wal_iterator.go",
        "write_batch.go",
        "write_batch_wi.go",
    ],
    cdeps = ["@com_github_facebook_rocksdb//:rocksdb"],
    cgo = True,
    importpath = "github.com/linxGnu/grocksdb",
    visibility = ["//visibility:public"],
)
