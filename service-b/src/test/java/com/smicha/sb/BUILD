load("//:junit5.bzl", "java_junit5_test")

java_junit5_test(
    name = "sb-unit-tests",
    srcs = glob(["*.java"]
    ),
    resources = glob(["src/test/resources/**"]),
    test_package = "com.smicha.sb",
    deps = [
        "//service-b/src/main/java/com/smicha/sb:sb",
    ],
)