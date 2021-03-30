load("@rules_java//java:defs.bzl", "java_binary")

package(default_visibility = ["//visibility:public"])

java_binary(
    name = "sa",
    srcs = glob(["**/*.java"]),
    resources = glob(["src/main/resources/**"]),
    main_class = "com.smicha.sa.App",
    runtime_deps = ["//shared/library-a:la"],
)
