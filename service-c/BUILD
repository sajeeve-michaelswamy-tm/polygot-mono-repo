load("@rules_java//java:defs.bzl", "java_binary", "java_test")

package(default_visibility = ["//visibility:public"])


java_binary(
    name = "service-b",
    main_class = "com.smicha.sc.App",
    srcs = glob(["src/main/java/com/smicha/sc/**/*.java"]),
    resources = glob(["src/main/resources/**"]),
    runtime_deps = [":library-a"],
)

java_test(
    name = "tests",
    srcs = glob(["src/test/java/com/smicha/sc/**/*.java"]),
    resources = glob(["src/test/resources/**"]),
    deps = [
        ":library-a",
        "@maven//:com_google_guava_guava",
        "@maven//:org_junit_jupiter_junit-jupiter-api",
    ],
)
