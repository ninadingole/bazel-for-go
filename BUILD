load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:go_naming_convention import_alias

gazelle(
    name = "gazelle",
    prefix = "github.com/ninadingole/bazel-for-go",
)

gazelle(
    name = "gazelle-update-repos",
    command = "fix",
    extra_args = [
        "-from_file=go.mod",
        "-to_macro=depz.bzl%go_dependencies",
        "-prune",
    ],
)
