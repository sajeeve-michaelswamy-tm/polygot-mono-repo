load("@rules_java//java:defs.bzl", "java_binary")

package(default_visibility = ["//visibility:public"])

java_binary(
    name = "sa",
    srcs = glob(["**/*.java"]),
    main_class = "com.smicha.sa.App",
    resources = glob(["src/main/resources/**"]),
    runtime_deps = ["//shared/library-a:la"],
)
