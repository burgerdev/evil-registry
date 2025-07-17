# Licensing

## Image

The tarball and config checked into this repo belong to the `amd64/linux` variant of this image:

```
docker.io/library/busybox:1.36.1-glibc@sha256:fc31e45c8ef62a1c88eecf59333a789b165424883b5960947eb52c1d9752de5d
```

To the best of the author's knowledge, that image corresponds to this
[Docker commit], which refers to this [Busybox download link], which is
licensed under GPLv2.

[Docker commit]: https://github.com/docker-library/busybox/commits/c48244ec06971ec9da046a065764cff4d92c3c25
[Busybox download link]: https://busybox.net/downloads/busybox-1.36.1.tar.bz2

The index- and manifest files found in `registry/variations/` are derived from the original ones stored in `registry/busybox/`.
In the case of the indices, all manifests save for `amd64/linux` have been removed, and depending on the purpose of each file,
digests have been modified to no longer fit what is being returned by the registry.

## Code

Copyright 2024 Markus Rudy

Permission to use, copy, modify, and/or distribute this software for any
purpose with or without fee is hereby granted, provided that the above
copyright notice and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED “AS IS” AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
PERFORMANCE OF THIS SOFTWARE.
