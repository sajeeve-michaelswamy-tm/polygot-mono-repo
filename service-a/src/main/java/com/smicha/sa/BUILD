load("@rules_java//java:defs.bzl", "java_binary")
load("@rules_java//java:defs.bzl", "java_proto_library")

package(default_visibility = ["//visibility:public"])


java_binary(
    name = "sa",
    srcs = glob(["**/*.java"]),
    main_class = "com.smicha.sa.App",
    resources = glob(["src/main/resources/**"]),
    runtime_deps = ["//shared/library-a:la",":hello_java_proto"],
)

java_proto_library(
    name = "hello_java_proto",
    deps = ["//shared/proto:hello_proto"],
)
