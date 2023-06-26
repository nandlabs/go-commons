# MESSAGING Client

This is a flexible and extensible messaging client designed to provide a unified interface for producing and consuming messages across various messaging platforms. It allows developers to seamlessly integrate messaging functionality into their applications without being tied to a specific messaging service.

---
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Extending the library](#extending-the-library)
---

## Features
* General producer interface for sending messages to different messaging platforms.
* General consumer interface for receiving and processing messages from different messaging platforms.
* Easy-to-use message abstraction for consistent handling of messages across platforms.
* Can be extended to multiple messaging platforms with easily pluggable interfaces, including:
  * AMQP (Advanced Message Queuing Protocol)
  * Apache Kafka
  * AWS SNS (Simple Notification Service)
  * AWS SQS (Simple Queue Service)
  * GCM (Google Cloud Messaging)
  * GCP Pub/Sub (Google Cloud Pub/Sub)

## Installation
To install the messaging client, use the following command:
```bash
go get go.nandlabs.io/commons/clients/messaging
```

## Usage
1. Import the library into your Go project:
    ```go
    import "github.com/your-repo/generic-messaging-client"
    ```
2. Initialize the messaging client for a specific platform. For example, to use the AMQP extension:
    ```go
    type AMQPMessaging struct {} // implements the Message interface defined under the library
    
    amqpMessageClient := &AMQPMessaging{}
    messaging.GetManager().Register(amqpMessageClient)
    ```
3. Send a message
   ```go
   message := &messaging.Message{
     Body: []byte("Hello, World!"), 
	 /// Add any additional properties or metadata
   }
   amqp := &AMQPMessaging{}
   destination := url.Parse("amqp://guest:password@localhost:5672/myvhost")
   err := amqp.Send(destination, message)
   if err != nil {
     // Handle error
   }
   ```
4. Consume a message
   ```go
   // Define the onConsume function
   onConsume := func(msg Message) error {
       // Process the message
       // ...

    return nil
   }
   // Start consuming messages from the channel
   OnMessage(consumerUrl, onConsume)
   ```
5. Repeat steps 2-4 for other messaging platforms by initializing the respective clients.

## Extending the library
To add support for additional messaging platforms, you can create new extensions by implementing the producer, consumer, and message interfaces defined in the library. These interfaces provide a consistent way to interact with different messaging systems.

You can refer to the existing extensions, such as amqp, kafka, etc., as examples for creating your own extensions. Ensure that your extension adheres to the interface definitions and follows the library's conventions for consistency.