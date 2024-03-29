"""repositories.bzl declares dependencies for the workspace
"""

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _maybe(repo_rule, name, **kwargs):
    if name not in native.existing_rules():
        repo_rule(name = name, **kwargs)

def repositories():
    """repositories loads all dependencies for the workspace
    """
    rules_proto()  # via <TOP>
    io_bazel_rules_go()  # via bazel_gazelle
    bazel_gazelle()  # via <TOP>
    build_stack_rules_proto()
    protobuf_core_deps()

def protobuf_core_deps():
    bazel_skylib()  # via com_google_protobuf
    rules_python()  # via com_google_protobuf
    zlib()  # via com_google_protobuf
    com_google_protobuf()  # via <TOP>

def io_bazel_rules_go():
    # Release: v0.39.1
    # TargetCommitish: release-0.39
    # Date: 2023-04-20 04:35:08 +0000 UTC
    # URL: https://github.com/bazelbuild/rules_go/releases/tag/v0.39.1
    # Size: 1759832 (1.8 MB)
    _maybe(
        http_archive,
        name = "io_bazel_rules_go",
        sha256 = "473a064d502e89d11c497a59f9717d1846e01515a3210bd169f22323161c076e",
        strip_prefix = "rules_go-0.39.1",
        urls = ["https://github.com/bazelbuild/rules_go/archive/v0.39.1.tar.gz"],
    )

def bazel_gazelle():
    # Branch: master
    # Commit: 2d1002926dd160e4c787c1b7ecc60fb7d39b97dc
    # Date: 2022-11-14 04:43:02 +0000 UTC
    # URL: https://github.com/bazelbuild/bazel-gazelle/commit/2d1002926dd160e4c787c1b7ecc60fb7d39b97dc
    #
    # fix updateStmt makeslice panic (#1371)
    # Size: 1859745 (1.9 MB)
    _maybe(
        http_archive,
        name = "bazel_gazelle",
        sha256 = "5ebc984c7be67a317175a9527ea1fb027c67f0b57bb0c990bac348186195f1ba",
        strip_prefix = "bazel-gazelle-2d1002926dd160e4c787c1b7ecc60fb7d39b97dc",
        urls = ["https://github.com/bazelbuild/bazel-gazelle/archive/2d1002926dd160e4c787c1b7ecc60fb7d39b97dc.tar.gz"],
    )

def local_bazel_gazelle():
    _maybe(
        native.local_repository,
        name = "bazel_gazelle",
        path = "/Users/i868039/go/src/github.com/bazelbuild/bazel-gazelle",
    )

def rules_proto():
    # Commit: f7a30f6f80006b591fa7c437fe5a951eb10bcbcf
    # Date: 2021-02-09 14:25:06 +0000 UTC
    # URL: https://github.com/bazelbuild/rules_proto/commit/f7a30f6f80006b591fa7c437fe5a951eb10bcbcf
    #
    # Merge pull request #77 from Yannic/proto_descriptor_set_rule
    #
    # Create proto_descriptor_set
    # Size: 14397 (14 kB)
    _maybe(
        http_archive,
        name = "rules_proto",
        sha256 = "9fc210a34f0f9e7cc31598d109b5d069ef44911a82f507d5a88716db171615a8",
        strip_prefix = "rules_proto-f7a30f6f80006b591fa7c437fe5a951eb10bcbcf",
        urls = ["https://github.com/bazelbuild/rules_proto/archive/f7a30f6f80006b591fa7c437fe5a951eb10bcbcf.tar.gz"],
    )

def build_stack_rules_proto():
    # Branch: master
    # Commit: 3d7eec0c990f6c4c37967ef4a0d58ba31075d5c6
    # Date: 2023-06-12 18:24:59 +0000 UTC
    # URL: https://github.com/stackb/rules_proto/commit/3d7eec0c990f6c4c37967ef4a0d58ba31075d5c6
    #
    # Fix proto_repository for @googleapis (#324)
    #
    # * Include vendor tree in proto_repository_tools_srcs
    # * Rename -override_go_googleapis to -reresolve_known_proto_imports
    # * Add googleapis to CI build
    # * Fix TestOverrideRule
    # Size: 2090870 (2.1 MB)
    _maybe(
        http_archive,
        name = "build_stack_rules_proto",
        sha256 = "92b02fc3040e4fbeba22843aa9b46eeb24b63c05c533bdc13e9afcaedd5a3d4f",
        strip_prefix = "rules_proto-3d7eec0c990f6c4c37967ef4a0d58ba31075d5c6",
        urls = ["https://github.com/stackb/rules_proto/archive/3d7eec0c990f6c4c37967ef4a0d58ba31075d5c6.tar.gz"],
    )
    # native.local_repository(
    #     name = "build_stack_rules_proto",
    #     path = "/Users/pcj/go/src/github.com/stackb/rules_proto",
    # )

def bazel_skylib():
    _maybe(
        http_archive,
        name = "bazel_skylib",
        sha256 = "ebdf850bfef28d923a2cc67ddca86355a449b5e4f38b0a70e584dc24e5984aa6",
        strip_prefix = "bazel-skylib-f80bc733d4b9f83d427ce3442be2e07427b2cc8d",
        urls = [
            "https://github.com/bazelbuild/bazel-skylib/archive/f80bc733d4b9f83d427ce3442be2e07427b2cc8d.tar.gz",
        ],
    )

def rules_python():
    _maybe(
        http_archive,
        name = "rules_python",
        sha256 = "8cc0ad31c8fc699a49ad31628273529ef8929ded0a0859a3d841ce711a9a90d5",
        strip_prefix = "rules_python-c7e068d38e2fec1d899e1c150e372f205c220e27",
        urls = [
            "https://github.com/bazelbuild/rules_python/archive/c7e068d38e2fec1d899e1c150e372f205c220e27.tar.gz",
        ],
    )

def zlib():
    _maybe(
        http_archive,
        name = "zlib",
        sha256 = "c3e5e9fdd5004dcb542feda5ee4f0ff0744628baf8ed2dd5d66f8ca1197cb1a1",
        strip_prefix = "zlib-1.2.11",
        urls = [
            "https://mirror.bazel.build/zlib.net/zlib-1.2.11.tar.gz",
            "https://zlib.net/zlib-1.2.11.tar.gz",
        ],
        build_file = "@build_stack_rules_proto//third_party:zlib.BUILD",
    )

def com_google_protobuf():
    _maybe(
        http_archive,
        name = "com_google_protobuf",
        sha256 = "d0f5f605d0d656007ce6c8b5a82df3037e1d8fe8b121ed42e536f569dec16113",
        strip_prefix = "protobuf-3.14.0",
        urls = [
            "https://github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
        ],
    )
