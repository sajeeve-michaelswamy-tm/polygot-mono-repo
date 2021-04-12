# bazel-mono-repo

## bazel-mono-repo
###       - service-a
###       - service-b
###       - service-c
###       - shared
####         - library-a
####         - library-b



Gradle
    - groovy or kotlin.
    - supports incremental builds ( but it has to scripted out properly , lot of boiler plate code involved in getting incremental builds working) 
    - difficult to support polygon monorepos ( multiple languages ) 

Bazel
    - starlark ( a subset of python )
    - incremental builds are truly incremental
    - the way bazel is designed - it just support polygot mono repos out of the box
    - builds are highly deterministic
    - highly parallel , especially if you run things on a build farm ( set of workers)  - both across workers and well as within a single worker ( multiplexed )

Paradigm shift

    -  We can stop following any gradle or maven conventions
    -  https://docs.bazel.build/versions/master/build-ref.html 
    -  https://blog.bazel.build/2019/09/29/intellij-bazel-sync.html
    -  https://docs.bazel.build/versions/master/visibility.html


CI : 
bazel build //...
bazel test  //...