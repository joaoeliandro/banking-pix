<h1 align="center">
    <br>
    Banking Pix Microservice
    <br>
</h1>

<h4 align="center">
   Microservice repository for pix transactions created in Go
</h4>

<p align="center">
  <img alt="GitHub top language" src="https://img.shields.io/github/languages/top/joaoeliandro/banking-pix-microservice.svg">

  <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/joaoeliandro/banking-pix-microservice.svg">

  <a href="https://www.codacy.com/app/joaoeliandro/banking-pix-microservice?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=joaoeliandro/banking-pix-microservice&amp;utm_campaign=Badge_Grade">
    <img alt="Codacy grade" src="https://api.codacy.com/project/badge/Grade/691b85e51bf240b997ae6ff82ea41590">
  </a>

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/joaoeliandro/banking-pix-microservice.svg">
  <a href="https://github.com/joaoeliandro/banking-pix-microservice/commits/master">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/joaoeliandro/banking-pix-microservice.svg">
  </a>

  <a href="https://github.com/joaoeliandro/banking-pix-microservice/issues">
    <img alt="Repository issues" src="https://img.shields.io/github/issues/joaoeliandro/banking-pix-microservice.svg">
  </a>

  <a href="https://github.com/joaoeliandro/banking-pix-microservice/blob/master/LICENSE">
    <img alt="GitHub License" src="https://img.shields.io/github/license/joaoeliandro/banking-pix-microservice.svg">
  </a>
</p>

<p align="center">
  <a href="#octocat-the-project">The Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#rocket-technologies">Technologies</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#memo-license">License</a>
</p>

## About

**Hello, friends! Baking Pix's microservice is responsible for intermediating bank transfers, receives the transfer transaction, forwards the transaction to the destination bank with "pending" status, receives confirmation from the destination bank with "confirmed" status and sends confirmation to the bank of origin with the processing information and delivering and marking with the status "completed".**

> In **Banking Pix Microservice**, focuses on using the most current technologies on the market, here is used:
> - Golang
> - PostgreSQL
> - Docker
> - gRPC
> - Apache Kafka
>
> Among other technologies for building the application.

---

## :octocat: The Project

Made with Hexagonal Architecture/Ports and Adapters, in addition to being able to work with a [gRPC][gRPC] server, it consumes and publishes messages on [Apache Kafka][ApacheKafka], with simultaneous operations when running the service. It works with a design focused on solving the domain problem and leaves the technical complexity to the "application layer", responsible for the [gRPC][gRPC] and [Kafka][ApacheKafka] server. It is also flexible for implementing other communication formats, such as API Rest, CLI clients, etc. Without changing any application component or domain model 

---

## :rocket: Technologies

the solution is developed with the technologies below:

> - [Golang](https://golang.org/)
> - [PostgreSQL](https://www.postgresql.org/)
> - [gRPC][gRPC]
> - [Apache Kafka][ApacheKafka]
> - [Docker](https://www.docker.com/)
> - [VS Code](https://code.visualstudio.com/)

## :memo: License

This project is under the MIT license. See the [LICENSE](https://github.com/joaoeliandro/banking-pix-microservice/blob/master/LICENSE) for more information.

[gRPC]: https://grpc.io/
[ApacheKafka]: https://kafka.apache.org/
