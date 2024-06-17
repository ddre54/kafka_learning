// See: https://github.com/kafka-rust/kafka-rust/blob/master/examples/example-produce.rs
use std::process;
use serde_json::json;

use std::time::Duration;
use kafka::producer::{Producer, Record, RequiredAcks};
// use kafka::error::Error as KafkaError;


fn main() {
    // tracing_subscriber::fmt::init();

    let broker = "localhost:9092";
    let topic = "my-topic";

    // println!("About to publish a message at {:?} to: {}, key: {:?},  msg: {:?}", vec![broker.to_owned()], topic, key_str, msg_str);
    let mut producer = Producer::from_hosts(vec![broker.to_owned()])
        // ~ give the brokers one second time to ack the message
        .with_ack_timeout(Duration::from_secs(1))
        // ~ require only one broker to ack the message
        .with_required_acks(RequiredAcks::One)
        // ~ build the producer with the above settings
        // .create()?;
        // Panic on error
        .create().unwrap();


    // ~ we leave the partition "unspecified" - this is a negative
    // partition - which causes the producer to find out one on its
    // own using its underlying partitioner.
    // producer.send(&Record {
    //     topic,
    //     partition: -1,
    //     key: (),
    //     value: data,
    // })?;

    let key = process::id().to_string();
    let key_json = json!(key);
    let key_str = key_json.to_string();
    let key_data = key_str.as_bytes();

    for i in 1..101 {
        // let msg = format!("Hello kafka from Rust:{} with json", key);
        let msg = format!("{key} Rust message #{i}");
        let msg_json = json!({
            "msg": msg
        });
        let msg_str = msg_json.to_string();
        let payload = msg_str.as_bytes();

        // Discarding error for now.
        let _ = producer.send(&Record {
            topic,
            partition: -1,
            key: key_data,
            value: payload,
        });
    }


    // ~ we can achieve exactly the same as above in a shorter way with
    // the following call
    // producer.send(&Record::from_value(topic, data))?;
    // Discarding error for now, by removing ?
    // let _ = producer.send(&Record::from_value(topic, payload));
}
