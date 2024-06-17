import os
from kafka import KafkaProducer

# producer = KafkaProducer(bootstrap_servers='localhost:9092')
# producer = KafkaProducer(bootstrap_servers='localhost:9092', key_serializer=str.encode, value_serializer=lambda v: json.dumps(v).encode('utf-8'))

# String to Bytes serialization
# producer = KafkaProducer(
#     bootstrap_servers='localhost:9092',
#     key_serializer=lambda k: k.encode('utf-8'),
#     value_serializer=lambda v: v.encode('utf-8')
# )
#

# Json Serialization
import json
producer = KafkaProducer(
    bootstrap_servers='localhost:9092',
    key_serializer=lambda k: json.dumps(k).encode('utf-8'),
    value_serializer=lambda v: json.dumps(v).encode('utf-8')
)

for i in range(1, 101):
    # producer.send('my-topic', b'Message #{}'.format(i))
    # payload = bytes('Message #{}'.format(i), encoding='utf-8')
    # payload = '{} Python message #{}'.format(os.getpid(), i).encode('utf-8')
    msg = '{} Python message #{}'.format(os.getpid(), i)
    payload = { 'msg': msg }
    # payload = '{} Python message #{}'.format(os.getpid(), i)
    print("payload: ", payload)
    # future = producer.send('my-topic', key=b'foo', value=b'bar')
    # key = '{}'.format(os.getpid()).encode('utf-8')
    key = '{}'.format(os.getpid())
    future = producer.send('my-topic', key=key, value=payload)
    # future = producer.send('my-topic', payload)
    result = future.get(timeout=60)
