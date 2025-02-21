# Test Client
In the context of service-to-service communication using gRPC, a test client typically refers to a client application 
or script used to interact with and test gRPC-based services. This client simulates the behavior of an actual service consumer, 
allowing developers to verify the functionality and performance of the service APIs. Here's a more detailed description:

__Key Roles of a Test Client in gRPC Communication__:
1. Request/Response Simulation:

- The test client sends RPC requests to the gRPC server, which hosts the actual service. These requests are typically defined in Protocol Buffers (proto files) and can include data structures for various service operations.
- It waits for the response from the server to verify correctness, checking if the server is correctly processing the request.

2. Service Discovery:

- In a real-world service-to-service communication setup, the client might need to discover available service endpoints or interact with a service registry. During testing, a test client can simulate these interactions if such service discovery is part of the workflow.

3. Error Handling and Edge Cases:

- The test client can be used to simulate error conditions such as timeouts, network issues, or unexpected responses, ensuring that the service handles edge cases gracefully.
It can also simulate invalid or malformed requests to test how the server responds to bad input.

4. Performance and Load Testing:

- A test client can also be used to measure the performance of the service by sending a high volume of requests to see how well the server handles load, checks response times, and identifies any performance bottlenecks.

5. Interaction with gRPC Services:

- The test client will often have code to directly invoke gRPC methods defined in the service's .proto file, handling serialization/deserialization of messages (via Protocol Buffers) and managing the lifecycle of the connection to the server.
- In typical scenarios, the client will communicate over either Unary RPCs (single request, single response) or Streaming RPCs (one-way or bidirectional streams).

__How It Works__:
- The test client typically connects to the gRPC server over HTTP/2, sends a request based on the method signature defined in the .proto file, and then checks the returned response to confirm the service behavior.
- It may use gRPC libraries available for various programming languages (e.g., Go, Python, Java, C++) to interact with the service.

__Example Scenario__:
- If a gRPC-based service offers a method like GetUserInfo, the test client would invoke this method with sample input (e.g., a user ID), and then compare the response (e.g., user details) to the expected result.

__Benefits of a Test Client__:
- Mocking and Simulations: Can simulate complex client behaviors without needing to set up a full production environment.
- Faster Feedback: Helps developers quickly identify bugs or issues in service logic, network communication, or message formatting.

- Overall, a test client plays a crucial role in ensuring that gRPC services behave as expected during development and in production environments.


## Dependencies
```shell
go get -v github.com/garcios/asset-trak-portfolio/portfolio-service@main
go get -v github.com/garcios/asset-trak-portfolio/currency-service@main
```

## How to test portfolio service
```shell
go run portfolio-client/main.go
```


## How to test currency service
```shell
go run currency-client/main.go
```