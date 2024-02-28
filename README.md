# Online Payment Platform

## About

Online Payment Platform is a web service that allows merchants (e.g. online stores) to accept payments from customers. The platform provides a RESTful API for merchants to create and manage payment requests, and for customers to pay for these requests.

## How to get started

You will need:

- [Go](https://golang.org/)
- [GVM (optional)](https://github.com/moovweb/gvm)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Git](https://git-scm.com/)
- [Make](https://www.gnu.org/software/make/)
- [Postman (optional)](https://www.postman.com/)
- [Curl (optional)](https://curl.se/)

### Running the application (docker-compose)

To run the application using `docker-compose`, you can use the following command:

```bash
make app_start
```

The command `make app_start` will build the application and start the services using `docker-compose`. The application will be available at `http://localhost:8080`.

### Running the application (local)

To run the application locally, you need to copy the `.env.example` file to `.env` and set the environment variables accordingly. Then, you can use the following command to start the dependencies:

```bash
make services_start
```

After starting the dependencies, you can use the following command to start the application:

```bash
make serve
```

> Please check the `Makefile` for more commands.

## How to use the application

To have a guide of how to use the application, please check the [Guide](./GUIDE.md).

The api specification is available at [api.yaml](./api/openapi.yml).

## Documentation

The documentation is organized as follows:

- [Home](./docs/README.md)
- [Architecture](./docs/architecture.md)
- [Requirements](./docs/requirements.md)
- [Database](./docs/database.md)
