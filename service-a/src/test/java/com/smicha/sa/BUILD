load("//:junit5.bzl", "java_junit5_test")

java_junit5_test(
    name = "unit-tests",
    srcs = glob(["src/test/java/com/smicha/sa/**/*.java"]
    ),
    resources = glob(["src/test/resources/**"]),
    test_package = "com.smicha.sa",
    deps = [
        "//service-a/src/main/java/com/smicha/sa:sa",
    ],
)
