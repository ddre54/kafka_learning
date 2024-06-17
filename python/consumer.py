from kafka import KafkaConsumer

# consumer = KafkaConsumer('my-topic')
# consumer = KafkaConsumer(
#     'my-topic',
#     bootstrap_servers='localhost:9092',
#     key_deserializer=lambda k: k.decode('utf-8'),
#     value_deserializer=lambda v: v.decode('utf-8')
# )

import json
consumer = KafkaConsumer(
    'my-topic',
    bootstrap_servers='localhost:9092',
    key_deserializer=lambda k: json.loads(k.decode('utf-8')),
    value_deserializer=lambda v: json.loads(v.decode('utf-8'))
)


for msg in consumer:
    # print("String value: {}, Byte value: {}, ConsumerRecord: {}".format(msg.value.decode('utf-8'), msg.value, msg))
    # print("String value: {}, Byte value: {}".format(msg.value.decode('utf-8'), msg.value))
    # print("String value: {}, Byte value: {}".format(msg.value, msg.value))
    print("String value[msg]: {}, value: {}".format(msg.value['msg'], msg.value))
    print("ConsumerRecord: {}".format(msg))
