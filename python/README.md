# Python Producer/Consumer

Kafka Producer/Consumer code in python

NOTES:
- Serialization for payload: JSON -> Bytes
- Deserialization for payload: Bytes -> JSON

# Requirements

List environments and installed versions:

```bash
pyenv versions
```

This is how to install required python version and create a virtual environment
for the project

```bash
# virtual environments from pyenv
pyenv install 3.12.2
pyenv virtualenv 3.12.2 new-env
pyenv activate new-env
pyenv deactive
# You can also use `pyenv local`
```

Example:
```bash
# virtual environments from pyenv
pyenv install 3.12.2
pyenv virtualenv 3.12.2 kafka_learning
pyenv activate kafka_learning
pyenv deactive
# You can also use `pyenv local`
```


## Dependencies

### Install `kafka-python`

```bash
pip install kafka-python
```

Error with `ModuleNotFoundError: No module named 'kafka.vendor.six.moves'`, install latest version

- ["import kafka" fails with "ModuleNotFoundError: No module named 'kafka.vendor.six.moves'" under Python 3.12](https://github.com/dpkp/kafka-python/issues/2412#issuecomment-1806341342)

```bash
pip install git+https://github.com/dpkp/kafka-python.git
```

Example:
```bash
pyenv activate kafka_learning
pip install kafka-python
```

# Running Consumer
```bash
python consumer.py
```

# Running Producer
```bash
python producer.py
```
