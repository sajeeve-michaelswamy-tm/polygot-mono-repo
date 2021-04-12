load("@rules_java//java:defs.bzl", "java_binary")

package(default_visibility = ["//visibility:public"])

java_binary(
    name = "sb",
    srcs = glob(["**/*.java"]),
    resources = glob(["src/main/resources/**"]),
    main_class = "com.smicha.sb.App",
    runtime_deps = ["//shared/library-b:lb"],
)