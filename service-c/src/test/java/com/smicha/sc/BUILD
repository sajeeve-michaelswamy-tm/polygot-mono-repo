load("//:junit5.bzl", "java_junit5_test")

java_junit5_test(
    name = "sc-unit-tests",
    srcs = glob(
        ["*.java"],
    ),
    resources = glob(["src/test/resources/**"]),
    test_package = "com.smicha.cb",
    deps = [
        "//service-c/src/main/java/com/smicha/sc",
    ],
)
