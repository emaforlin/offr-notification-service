# Gjallarhorn

Gjallarhorn is a powerful and scalable notification service designed to handle multiple communication channels such as email (by the moment) ~~and push notifications~~.

## Features

* **Multi-Channel Notifications**: Supports email, and will support push notifications.
* **gRPC API**: Provides a high-performance gRPC-based API for efficient communication.
* **Configuration file**: Easy setup via `config.yaml` file.


## API Endpoints

| Method | Description |
| --- | --- |
| SendEmail (Unary) | Send a single email |
<!-- | SendEmailBatch (Stream) | Send a batch of emails | -->


## License
This project is licensed under the MIT License.