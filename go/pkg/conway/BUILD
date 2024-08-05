go_library(
    name = "conway",
    srcs = ["conway.go"],
    visibility = ["//go/cmd/conway:all"],
)

go_test(
    name = "conway_test",
    srcs = ["conway_test.go"],
    deps = [
        ":conway",
        "///third_party/go/github.com_stretchr_testify//assert",
    ],
)
