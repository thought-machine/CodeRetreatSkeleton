package(default_visibility = ["PUBLIC"])

filegroup(
    name = "modfile",
    srcs = ["go.mod"],
)

go_toolchain(
    name = "toolchain",
    install_std = False,
    version = "1.22.5",
)

go_stdlib(
    name = "std",
    tags = [
        "netgo",
        "osusergo",
    ],
)

PUKU_VERSION = "1.11.0"

remote_file(
    name = "puku",
    out = "puku",
    url = f"https://github.com/please-build/puku/releases/download/v{PUKU_VERSION}/puku-{PUKU_VERSION}-linux_amd64",
    hashes = ["c3e4c8bd035487504e568d791d36a4a45c5b8f59a58dc1d8863ba39338b3fb24"],
    extract = False,
    binary = True,
)
