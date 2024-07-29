package main

import (
	"context"
	"flag"
	"log"

	"github.com/regclient/regclient"
	"github.com/regclient/regclient/types/ref"
)

func main() {
	flag.Parse()
	ctx := context.Background()

	for _, img := range flag.Args() {
		c := regclient.New()
		ref, err := ref.New(img)
		if err != nil {
			log.Fatalf("Could not parse image reference %q: %v", img, err)
		}
		m, err := c.ManifestGet(ctx, ref)
		if err != nil {
			log.Fatalf("Could not pull manifest for %q: %v", img, err)
		}
		log.Printf("ManifestGet(%q) -> %q", img, m.GetDescriptor().Digest)
	}
}
