package registry

import (
	"crypto/sha256"
	_ "embed"
	"fmt"
	"net"
	"net/http"
	"strings"
)

const digestHeader = "docker-content-digest"

var (
	//go:embed busybox/index.json
	index            []byte
	IndexDigest      string
	WrongIndexDigest string
	//go:embed busybox/manifest.json
	manifest       []byte
	ManifestDigest string
	//go:embed busybox/blob.tar.gz
	blob       []byte
	BlobDigest string
	//go:embed busybox/config.json
	config       []byte
	ConfigDigest string

	//go:embed variations/index_for_evil_manifest.json
	indexForEvilManifest       []byte
	IndexForEvilManifestDigest string

	//go:embed variations/index_for_manifest_for_evil_blob.json
	indexForManifestForEvilBlob       []byte
	IndexForManifestForEvilBlobDigest string

	//go:embed variations/manifest_for_evil_blob.json
	manifestForEvilBlob       []byte
	ManifestForEvilBlobDigest string
)

func digest(data []byte) string {
	return fmt.Sprintf("sha256:%x", sha256.Sum256(data))
}

func init() {
	IndexDigest = digest(index)
	WrongIndexDigest = "sha256:1111111111111111111111111111111111111111111111111111111111111111"
	ManifestDigest = digest(manifest)
	BlobDigest = digest(blob)
	ConfigDigest = digest(config)

	IndexForEvilManifestDigest = digest(indexForEvilManifest)
	IndexForManifestForEvilBlobDigest = digest(indexForManifestForEvilBlob)
	ManifestForEvilBlobDigest = digest(manifestForEvilBlob)
}

func manifestHandler(rw http.ResponseWriter, req *http.Request) {
	respDigest := ManifestDigest
	if reqDigest := req.PathValue("digest"); strings.HasPrefix(reqDigest, "sha256:") {
		respDigest = reqDigest
	}

	contentType := "application/vnd.docker.distribution.manifest.v2+json"
	contentTypeIndex := "application/vnd.oci.image.index.v1+json"

	var response []byte
	switch respDigest {
	case IndexForManifestForEvilBlobDigest:
		response = indexForManifestForEvilBlob
		contentType = contentTypeIndex
	case IndexForEvilManifestDigest:
		response = indexForEvilManifest
		contentType = contentTypeIndex
	case ManifestForEvilBlobDigest:
		response = manifestForEvilBlob
	case IndexDigest, WrongIndexDigest:
		response = index
		contentType = contentTypeIndex
	default:
		response = manifest
	}

	rw.Header().Set(digestHeader, respDigest)
	rw.Header().Set("Content-Type", contentType)
	rw.Write(response)
}

func blobHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.PathValue("digest") {
	case ConfigDigest:
		rw.Header().Set(digestHeader, ConfigDigest)
		rw.Write(config)
	default:
		rw.Header().Set(digestHeader, BlobDigest)
		rw.Write(blob)
	}
}

func v2Handler(rw http.ResponseWriter, _ *http.Request) {
	rw.Header().Set("Docker-Distribution-API-Version", "registry/2.0")
	rw.WriteHeader(http.StatusOK)
}

func Run(listener net.Listener, mux *http.ServeMux) {
	mux.Handle("/v2/busybox/manifests/{digest}", http.HandlerFunc(manifestHandler))
	mux.Handle("/v2/busybox/blobs/{digest}", http.HandlerFunc(blobHandler))
	mux.Handle("/v2/", http.HandlerFunc(v2Handler))

	http.Serve(listener, mux)
}
