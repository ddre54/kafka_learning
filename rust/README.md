# Rust Producer/Consumer

Kafka Producer/Consumer code in rust

NOTES:
- Serialization for payload: JSON -> Bytes
- Deserialization for payload: Bytes -> JSON
- Handle errors

# Requirements

- cargo 1.79.0

- [Install Rust](https://www.rust-lang.org/learn/get-started)
- [Kafka crate](https://crates.io/crates/kafka)
- [Kafka producer/consumer](https://dev.to/ciscoemerge/how-to-build-a-simple-kafka-producerconsumer-application-in-rust-3pl4)

## Dependencies

# Running Consumer

```bash
cd consumer
cargo run
```

Reference:
- [Kafka Consumer](https://github.com/kafka-rust/kafka-rust/blob/master/examples/example-consume.rs)

# Running Producer

```bash
cd producer
cargo run
```
Reference:
- [Kafka Producer](https://github.com/kafka-rust/kafka-rust/blob/master/examples/example-produce.rs)
