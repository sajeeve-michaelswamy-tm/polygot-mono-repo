load("@rules_java//java:defs.bzl", "java_library", "java_test")

package(default_visibility = ["//visibility:public"])

java_library(
    name = "library-a",
    srcs = glob(["src/main/java/com/smicha/shared/la/**/*.java"]),
    resources = glob(["src/main/resources/**"]),
    deps = ["@maven//:com_google_guava_guava"],
)

java_test(
    name = "tests",
    srcs = glob(["src/test/java/com/smicha/shared/la/**/*.java"]),
    resources = glob(["src/test/resources/**"]),
    deps = [
        "@maven//:com_google_guava_guava",
        "@maven//:org_junit_jupiter_junit-jupiter-api",
    ],
)
