use glob::glob;
use std::fs::File;

fn main() {
    let mut files = Vec::new();
    for f in glob("forge-gRPC/src/main/proto/*.proto").expect("fail to find proto files. is the submodules initialized?") {
        let f = f.unwrap();
        files.push(f);
    }
    let paths: Vec<&str> = files.iter().map(|x| x.to_str().unwrap()).collect();
    let names: Vec<&str> = files.iter().map(|x| x.file_stem().map(|x| x.to_str().unwrap()).unwrap()).collect();
    let modules_location = std::path::PathBuf::from("src/protocol");
    std::fs::create_dir_all(&modules_location).expect("unable to create src/protocol directory");
    protoc_rust_grpc::run(protoc_rust_grpc::Args{
        out_dir: modules_location.to_str().unwrap(),
        includes: &[],
        input: &paths,
        rust_protobuf: true,
    }).expect("fail to run protoc-rust-grpc");
    let mut f = File::create(modules_location.join("mod.rs")).expect("unable to create mod files for grpc");
    for n in names {
        use std::io::Write;
        f.write_all(b"// GENERATED. DO NOT EDIT!\n").unwrap();
        f.write_all(format!("pub mod {};\n", n).as_bytes()).unwrap();
        f.write_all(format!("pub mod {}_grpc;\n", n).as_bytes()).unwrap();
    }
}