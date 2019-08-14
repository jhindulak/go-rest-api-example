FROM scratch

WORKDIR /app
COPY /dist/go-rest-api /app

CMD ["/app/go-rest-api"]