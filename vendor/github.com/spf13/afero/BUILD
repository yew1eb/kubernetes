package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "afero.go",
        "basepath.go",
        "cacheOnReadFs.go",
        "const_win_unix.go",
        "copyOnWriteFs.go",
        "httpFs.go",
        "ioutil.go",
        "memmap.go",
        "memradix.go",
        "os.go",
        "path.go",
        "readonlyfs.go",
        "regexpfs.go",
        "sftp.go",
        "unionFile.go",
        "util.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/pkg/sftp:go_default_library",
        "//vendor/github.com/spf13/afero/mem:go_default_library",
        "//vendor/github.com/spf13/afero/sftp:go_default_library",
        "//vendor/golang.org/x/text/transform:go_default_library",
        "//vendor/golang.org/x/text/unicode/norm:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//vendor/github.com/spf13/afero/mem:all-srcs",
        "//vendor/github.com/spf13/afero/sftp:all-srcs",
    ],
    tags = ["automanaged"],
)
