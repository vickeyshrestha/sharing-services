### Running the example

- Make sure the NATS message bus server is running. You can visit https://docs.nats.io/running-a-nats-service/introduction/installation
- Inside the both pub.go & sub.go, make sure the host and port for NATS is correct. Since I am running under localhost on port 5101, I have setup mine accordingly.
- Run sub.go file. Without the pub.go active at first, Subscriber should not stream any messages.
- While sub.go is actively running, start the pub.go. This should begin publishing numerical data in message bus.
- Check the sub.go console, and the Subscriber should pick the data streamed