# Documentations

## NATS vs Redis as message broker
When comparing NATS and Redis as message brokers, NATS generally excels in raw performance and extremely low latency, making it ideal for high-throughput, real-time applications, while Redis offers more flexibility with its data structures and can be a good choice for scenarios where you need both messaging and data storage capabilities, but may have slightly higher latency than NATS in pure messaging scenarios.
Key Differences:

__Focus__:
NATS is primarily designed as a high-performance message broker, prioritizing speed and low latency, whereas Redis is a more versatile in-memory data store that also supports messaging functionality through its "Lists" and "Streams" features.

__Data Structures__:
Redis provides a wide range of data structures like strings, lists, sets, sorted sets, and more, which can be used for various data operations beyond just messaging, while NATS primarily focuses on simple message delivery with topics.

__Persistence__:
While both can be configured with persistence options, NATS generally prioritizes in-memory processing for maximum speed, whereas Redis offers more robust persistence mechanisms for data durability.

__When to choose NATS__:
- High-frequency, low-latency messaging systems
- Real-time applications like live chat or stock tickers
- Microservices communication where fast message delivery is critical

__When to choose Redis__:
- Situations where you need both messaging and data storage capabilities within a single system
- Applications requiring flexible data structures for complex data operations
Scenarios where some level of data persistence is necessary 
- Scenarios where some level of data persistence is necessary 
