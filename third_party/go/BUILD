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
    binary = True,
    extract = False,
    hashes = ["c3e4c8bd035487504e568d791d36a4a45c5b8f59a58dc1d8863ba39338b3fb24"],
    url = f"https://github.com/please-build/puku/releases/download/v{PUKU_VERSION}/puku-{PUKU_VERSION}-linux_amd64",
)

go_repo(
    licences = ["MIT"],
    module = "github.com/stretchr/testify",
    version = "v1.9.0",
)

go_repo(
    licences = ["ISC"],
    module = "github.com/davecgh/go-spew",
    version = "v1.1.1",
)

go_repo(
    licences = ["BSD-3-Clause"],
    module = "github.com/pmezard/go-difflib",
    version = "v1.0.0",
)

go_repo(
    licences = ["MIT"],
    module = "gopkg.in/yaml.v3",
    version = "v3.0.1",
)
