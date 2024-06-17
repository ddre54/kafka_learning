// See: https://github.com/kafka-rust/kafka-rust/blob/master/examples/example-consume.rs

use kafka::consumer::{Consumer, FetchOffset, GroupOffsetStorage};
use serde_json::Value;
use std::str;
// use std::time::{Duration, Instant};
use std::time::Instant;
// use std::time::SystemTime;

fn main() {
    let broker = "localhost:9092".to_owned();
    let topic = "my-topic".to_owned();

    // let mut con = Consumer::from_hosts(brokers)
    //     .with_topic(topic)
    //     .with_group(group)
    //     .with_fallback_offset(FetchOffset::Earliest)
    //     .with_offset_storage(Some(GroupOffsetStorage::Kafka))
    //     .create()?;

    let mut consumer = Consumer::from_hosts(vec![broker])
        .with_topic(topic)
        // .with_group(group)
        .with_fallback_offset(FetchOffset::Earliest)
        .with_offset_storage(Some(GroupOffsetStorage::Kafka))
        .create()
        .unwrap();

    let mut now = Instant::now();
    loop {
        let mss = consumer.poll().unwrap();
        if mss.is_empty() {
            println!("No messages available right now. Last message {:.2?} ago.", now.elapsed());
        } else {
            now = Instant::now();
        }
        for ms in mss.iter() {
            for m in ms.messages() {
                // Convert bytes to string
                let msg = str::from_utf8(&m.value).unwrap();
                println!(
                    "{}:{}@{}: {:?}",
                    ms.topic(),
                    ms.partition(),
                    m.offset,
                    // m.value
                    msg
                );
                // println!("String value[msg]: {:?}, value: {:?}", m.value['msg'], m.value);
                // let v: Value = serde_json::from_slice(m.value).unwrap();
                // let v: Value = serde_json::from_str(&msg).expect("REASON - Expected json");
                // let v: Value = serde_json::from_slice(&bytes).unwrap();
                // Convert Bytes to JSON
                // Don't error out - if Json parse from bytes -> json succeed
                if let Ok(value) = serde_json::from_slice::<Value>(m.value) {
                    // use the returned value
                    // Valid JSON object: Bytes to JSON
                    println!("String value: {:#?}, value: {}", msg, value["msg"]);
                } else {
                    // Not a valid JSON object
                    println!("Not a valid JSON: {:#?}", msg);
                }
            }
            // Don't error out
            let _ = consumer.consume_messageset(ms);
        }
        // consumer.commit_consumed()?;
        // Don't error out
        let _ = consumer.commit_consumed();
        // consumer.commit_consumed().unwrap();
    }
}
