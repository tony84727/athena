mod protocol;

use protocol::{chat_grpc, chat};
use protocol::chat_grpc::Chat;
use std::sync::mpsc::channel;
use std::sync;
use std::iter;
use grpc::ClientStub;

fn main() {
    let grpc_client = sync::Arc::new(grpc::Client::new_plain("0.0.0.0", 30000, grpc::ClientConf::new()).unwrap());
    let client = chat_grpc::ChatClient::with_client(grpc_client);
    let (sender, receiver) = channel::<String>();
    let request = grpc::StreamingRequest::iter(iter::from_fn(move ||{
        let content = receiver.recv().unwrap();
        print!("received: {}", content);
        let mut m = chat::Message::new();
        m.set_content(String::from(content));
        Some(m)
    }));

    std::thread::spawn(move || {
        sender.send(String::from("hi")).unwrap();
    });

    let response = client.connect(grpc::RequestOptions::new(), request);
    match response.wait() {
        Err(e) => panic!(e),
        Ok((_, stream)) => {
            for s in stream {
                match s {
                    Err(e) => panic!(e),
                    Ok(s) => {
                        println!("sender: {}, message: {}", s.sender, s.content);
                    }
                };
            }
        },
    }
}
