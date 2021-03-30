load("//:junit5.bzl", "java_junit5_test")

java_junit5_test(
    name = "sa-unit-tests",
    srcs = glob(["*.java"]
    ),
    resources = glob(["src/test/resources/**"]),
    test_package = "com.smicha.sa",
    deps = [
        "//service-a/src/main/java/com/smicha/sa:sa",
    ],
)
