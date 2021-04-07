/**
 * https://docs.docker.com/engine/extend/
 *
 */
plugins {
    id("com.github.blindpirate.gogradle")
}

golang {
    packagePath = "main"
}


dependencies {
    golang {
        build(
            mapOf(
                "name" to "github.com/coreos/go-systemd",
                "version" to "64d5cd7cb947834ef93874e82745c42ad6de4d0e",
                "subpackages" to arrayOf("activation", "util", "."),
                "transitive" to false
            )
        )
        build(
            mapOf(
                "name" to "github.com/coreos/pkg",
                "version" to "447b7ec906e523386d9c53be15b55a8ae86ea944",
                "subpackage" to arrayOf("dlopen", "."),
                "transitive" to false
            )
        )
        build(
            mapOf(
                "name" to "github.com/docker/distribution",
                "version" to "fbe6e8d212ed880cf7f7c9f876ba9b15c8221c5f",
                "subpackage" to arrayOf("digest", "reference", "."),
                "transitive" to false
            )
        )
        build(
            mapOf(
                "name" to "github.com/docker/engine-api",
                "version" to "3d1601b9d2436a70b0dfc045a23f6503d19195df",
                "subpackage" to arrayOf(
                    "types/time",
                    "types",
                    "types/filters",
                    "types/container",
                    "types/reference",
                    "types/swarm",
                    "types/network",
                    "client/transport/cancellable",
                    "types/versions",
                    ".",
                    "types/strslice",
                    "types/blkiodev",
                    "client/transport",
                    "client",
                    "types/registry"
                ),
                "transitive" to false
            )
        )
        build(
            mapOf(
                "name" to "github.com/docker/go-connections",
                "version" to "f512407a188ecb16f31a33dbc9c4e4814afc1b03",
                "subpackage" to arrayOf("nat", "tlsconfig", "sockets", "."),
                "transitive" to false
            )
        )
        build(
            mapOf(
                "name" to "github.com/docker/go-plugins-helpers",
                "version" to "60d242cfd0fb30e5002fbf76bf6872e81e85adba",
                "subpackage" to arrayOf("volume", ".", "sdk"),
                "transitive" to false
            )
        )
        build(
            mapOf(
                "name" to "github.com/docker/go-units",
                "version" to "8a7beacffa3009a9ac66bad506b18ffdd110cf97",
                "transitive" to false
            )
        )
        build(
            mapOf(
                "name" to "github.com/fatih/color",
                "version" to "dea9d3a26a087187530244679c1cfb3a42937794",
                "transitive" to false
            )
        )
        build(
            mapOf(
                "name" to "github.com/mattn/go-colorable",
                "version" to "6e26b354bd2b0fc420cb632b0d878abccdc6544c",
                "transitive" to false
            )
        )
        build(
            mapOf(
                "name" to "github.com/mattn/go-isatty",
                "version" to "66b8e73f3f5cda9f96b69efd03dd3d7fc4a5cdb8",
                "transitive" to false
            )
        )
        build(
            mapOf(
                "name" to "github.com/Microsoft/go-winio",
                "version" to "ce2922f643c8fd76b46cadc7f404a06282678b34",
                "transitive" to false
            )
        )
        build(
            mapOf(
                "name" to "github.com/opencontainers/runc",
                "version" to "49ed0a10e4edba88f9221ec730d668099f6d6de8",
                "subpackage" to arrayOf(".", "libcontainer/user"),
                "transitive" to false
            )
        )
        build(
            mapOf(
                "name" to "github.com/Sirupsen/logrus",
                "version" to "55eb11d21d2a31a3cc93838241d04800f52e823d",
                "subpackage" to arrayOf("formatters/logstash", "."),
                "transitive" to false
            )
        )
        build(
            mapOf(
                "name" to "golang.org/x/net",
                "version" to "4876518f9e71663000c348837735820161a42df7",
                "subpackage" to arrayOf("context", "proxy", "."),
                "transitive" to false
            )
        )
        build(
            mapOf(
                "name" to "golang.org/x/sys",
                "version" to "c200b10b5d5e122be351b67af224adc6128af5bf",
                "subpackage" to arrayOf("windows", ".", "unix"),
                "transitive" to false
            )
        )
    }
}
