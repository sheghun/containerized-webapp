FROM alpine:3.7

# Work directory
WORKDIR /app

# Copy all files in the root project directory
COPY webapp .
COPY web/build web/build

EXPOSE 4222

CMD ["./webapp"]