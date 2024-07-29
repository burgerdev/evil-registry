# evil-registry

OCI registry clients are supposed to check that digests of retrieved resources match expectations:
<https://github.com/opencontainers/distribution-spec/blob/main/spec.md#pull>.

This repo simulates an OCI registry that maliciously serves false content.
The `registry` dir contains a server that hosts a single registry and repo: `localhost:80/busybox`.
It will respond with the same image manifest regardless of the requested digest, and will add the `Docker-Content-Digest` header that was requested, not the actual one.

```sh
go run ./registry/main.go --addr :31313 &
curl -v localhost:31313/v2/busybox/manifests/sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef | sha256sum
```

```data
[...]
< HTTP/1.1 200 OK
< Content-Type: application/vnd.docker.distribution.manifest.v2+json
< Docker-Content-Digest: sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef
[...]
d319b0e3e1745e504544e931cde012fc5470eba649acc8a7b3607402942e5db7  -
```

`rust` and `go` contain example clients, to evaluate whether these clients validate digests.
