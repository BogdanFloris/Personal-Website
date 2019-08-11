# personal-website

This repository contains the source code for the personal website, built using [Rust](https://www.rust-lang.org/) and the [Rocket](https://rocket.rs/) web framework.

## Build and run using Docker

To build the image, you need to generate a secret key using: `openssl rand -base64 32`. Then run:
```shell script
docker build -t personal-website --build-arg SECRET_KEY=<YOUR_SECRET_KEY> .
```

To run the website:
```shell script
docker run -p 8080:8080 personal-website
```
