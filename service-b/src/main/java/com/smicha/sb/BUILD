load("@rules_java//java:defs.bzl", "java_binary")

package(default_visibility = ["//visibility:public"])

java_binary(
    name = "sb",
    srcs = glob(["**/*.java"]),
    main_class = "com.smicha.sb.App",
    resources = glob(["src/main/resources/**"]),
    runtime_deps = ["//shared/library-b:lb"],
)
