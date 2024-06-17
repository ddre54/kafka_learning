# Kafka - Apache (local)

Running Kafka locally in Dockers to learn using different code for Producer/Consumer.

NOTES:
- Producer: Serialization for payload: JSON -> Bytes
- Consumer: Deserialization for payload: Bytes -> JSON
- Handle errors

TODO:
- Better error handling
- Improve code quality

# Requirements

- [Dockers](https://docs.docker.com/get-docker/?_gl=1*d1jyzk*_ga*ODU1MTAxMzQ1LjE3MDM2NzU0NDM.*_ga_XJWPQMJYHQ*MTcwMzY3NTQ0My4xLjEuMTcwMzY3NTQ0My42MC4wLjA.)

# References

- [Getting Started with Apache Kafka on Docker: A Step-by-Step Guide](https://medium.com/@amberkakkar01/getting-started-with-apache-kafka-on-docker-a-step-by-step-guide-48e71e241cf2)

# Running Environment

## Kafka

```bash
docker compose up -d
```

NOTE:
- Taking into consideration the platform (arch)

```
# Dockerfile
FROM --platform=linux/amd64

# Docker image buld platform specific
docker buildx build --platform=linux/amd64 -t $DOCKER_TAG_SHA -t $DOCKER_TAG_LATEST .
```

### Create a Kafka topic

Get the `<kafka-container-id>` with:

```bash
docker ps
```

Replace the `<kafka-container-id>` with the `CONTAINER ID` for Kafka from the previous command:

```bash
docker exec -it <kafka-container-id> /opt/kafka/bin/kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic <new-topic>
```

Example:
```bash
docker exec -it 4ef041bba319 /opt/kafka/bin/kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic my-topic
```

### Produce and Consume messages

#### Produce

```bash
docker exec -it <kafka-container-id> /opt/kafka/bin/kafka-console-producer.sh --broker-list localhost:9092 --topic <topic>
```

#### Consume

```bash
docker exec -it <kafka-container-id> /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic <topic> --from-beginning
```

### Stop and remove Kafka containers

```bash
docker compose down
```
