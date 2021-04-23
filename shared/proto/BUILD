load("@rules_java//java:defs.bzl", "java_proto_library")

java_proto_library(
    name = "di_java_proto",
    deps = ["//shared/proto/di:di_proto"],
    visibility = ["//service-a:__subpackages__","//service-b:__subpackages__"],
)

java_proto_library(
    name = "payment_java_proto",
    deps = ["//shared/proto/payment:payment_proto"],
    visibility = ["//service-a:__subpackages__","//service-c:__subpackages__"],
)
