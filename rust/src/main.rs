use oci_distribution::{
    client::{linux_amd64_resolver, ClientConfig, ClientProtocol},
    Client,
    Reference,
};

#[tokio::main]
async fn main() {
    for arg in std::env::args().skip(1) {
        let image: Reference = arg.parse().unwrap();

        let _ = Client::new(ClientConfig {
            protocol: ClientProtocol::HttpsExcept(vec![image.registry().to_owned()]),
            platform_resolver: Some(Box::new(linux_amd64_resolver)),
            ..Default::default()
        });
        let client = Client::default();
        let auth = &oci_distribution::secrets::RegistryAuth::Anonymous;
        let accepted_media_types = vec!["application/vnd.docker.image.rootfs.diff.tar.gzip"];
    
        let (_m, d) = client.pull_manifest(&image, auth).await.unwrap();
    
        println!("pull_manifest({}) -> {}", arg, d);

        let data = client.pull(&image, auth, accepted_media_types).await.unwrap();
    
        println!("pull({}) -> {:?}", arg, data.digest);
    }
}
