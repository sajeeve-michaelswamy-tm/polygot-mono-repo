load("@rules_java//java:defs.bzl", "java_library", "java_test")

java_library(
    name = "lb",
    srcs = glob(["src/main/java/com/smicha/shared/lb/**/*.java"]),
    resources = glob(["src/main/resources/**"]),
    visibility = [
        "//service-b:__subpackages__",
        "//service-c:__subpackages__",
    ],
    deps = ["@maven//:com_google_guava_guava"],
)
