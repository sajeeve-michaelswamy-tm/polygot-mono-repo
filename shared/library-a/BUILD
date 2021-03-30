load("@rules_java//java:defs.bzl", "java_library", "java_test")


java_library(
    name = "la",
    srcs = glob(["src/main/java/com/smicha/shared/la/**/*.java"]),
    resources = glob(["src/main/resources/**"]),
    deps = ["@maven//:com_google_guava_guava"],
    visibility = ["//service-a:__subpackages__"]
)
