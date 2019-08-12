FROM rustlang/rust:nightly-slim
ARG SECRET_KEY

COPY . /personal-website
WORKDIR /personal-website

EXPOSE 8080

RUN cargo build --release

ENV PW_SECRET_KEY=$SECRET_KEY
CMD ROCKET_SECRET_KEY=$PW_SECRET_KEY cargo run --release
