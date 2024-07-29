package main

import (
	"crypto/sha256"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

const digestHeader = "docker-content-digest"

var (
	addr = flag.String("addr", ":80", "address to serve on")
)

var (
	//go:embed manifest.json
	manifest       []byte
	manifestDigest string
	//go:embed blob.tar.gz
	blob       []byte
	blobDigest string
	//go:embed config.json
	config       []byte
	configDigest string
)

func digest(data []byte) string {
	return fmt.Sprintf("sha256:%x", sha256.Sum256(data))

}

func init() {
	manifestDigest = digest(manifest)
	blobDigest = digest(blob)
	configDigest = digest(config)
}

func manifestHandler(rw http.ResponseWriter, req *http.Request) {
	respDigest := manifestDigest
	if reqDigest := req.PathValue("digest"); strings.HasPrefix(reqDigest, "sha256:") {
		respDigest = reqDigest
	}
	rw.Header().Set(digestHeader, respDigest)
	rw.Header().Set("Content-Type", "application/vnd.docker.distribution.manifest.v2+json")
	rw.Write(manifest)
}

func blobHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.PathValue("digest") {
	case configDigest:
    rw.Header().Set(digestHeader, configDigest)
		rw.Write(config)
	case blobDigest:
    rw.Header().Set(digestHeader, blobDigest)
		rw.Write(blob)
	default:
		rw.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	flag.Parse()

	http.DefaultServeMux.Handle("/v2/busybox/manifests/{digest}", http.HandlerFunc(manifestHandler))
	http.DefaultServeMux.Handle("/v2/busybox/blobs/{digest}", http.HandlerFunc(blobHandler))

	listener, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("Could not listen on %s: %v", *addr, err)
	}
	log.Printf("serving a registry on %s ...", *addr)
	http.Serve(listener, http.DefaultServeMux)
}
