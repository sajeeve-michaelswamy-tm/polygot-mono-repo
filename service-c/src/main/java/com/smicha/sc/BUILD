load("@rules_java//java:defs.bzl", "java_binary")

package(default_visibility = ["//visibility:public"])

java_binary(
    name = "sc",
    srcs = glob(["**/*.java"]),
    main_class = "com.smicha.sc.App",
    resources = glob(["src/main/resources/**"]),
    runtime_deps = ["//shared/library-b:lb"],
)
