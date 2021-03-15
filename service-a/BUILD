load("@rules_java//java:defs.bzl", "java_binary","java_test")

package(default_visibility = ["//visibility:public"])


java_binary(
    name = "service-a",
    srcs = glob(["src/main/java/com/smicha/sa/**/*.java"]),
    resources = glob(["src/main/resources/**"]),
    main_class = "com.smicha.sa.App",
    runtime_deps = [":library-a",":library-b"],
)

java_test(
    name = "unit-tests",
    srcs = glob(["src/test/java/com/smicha/sa/**/*.java"]),
    resources = glob(["src/test/resources/**"]),
    deps = [
        ":library-a",
        ":library-b",
        "@maven//:com_google_guava_guava",
        "@maven//:org_junit_jupiter_junit-jupiter-api",
    ],
)
