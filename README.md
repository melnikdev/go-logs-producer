# Kafka Log Producer

## 📌 Description

This **Kafka Producer** is written in **Go**, responsible for sending logs to a **Kafka** topic. It supports **gRPC streaming** for efficient log transmission and can handle high-throughput logging.

---

## 🚀 Running the Producer

### 1️⃣ Install dependencies
```sh
go mod tidy
```

### 2️⃣ Configure `.env`
Create a `.env` file and set the variables:
```ini
KAFKA_BROKER=localhost:9092
KAFKA_TOPIC=logs
SERVER_PORT=50051
SERVER_NETWORK=tcp
```

### 3️⃣ Start the Producer
```sh
go run main.go
```

---

## ⚙ Configuration

| Variable        | Description                     | Default Value     |
|----------------|---------------------------------|------------------|
| `KAFKA_BROKER` | Kafka broker address           | `localhost:9092` |
| `KAFKA_TOPIC`  | Kafka topic for logs           | `logs`           |
| `SERVER_PORT`  | gRPC server port               | `50051`          |

---

## 🛠 Architecture
1. **gRPC Client** sends logs to the producer.
2. The **Kafka Producer** processes and pushes logs to a Kafka topic.
3. Logs are consumed by the **Kafka Consumer** and stored in **ClickHouse**.

```plaintext
gRPC Client -> gRPC Server (Producer) -> Kafka -> Consumer -> ClickHouse
```

---



## 📜 License
MIT License.

