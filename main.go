package main

import (
    "github.com/secmohammed/golang-kafka-grpc-poc/config"
    "github.com/secmohammed/golang-kafka-grpc-poc/container"
    "github.com/secmohammed/golang-kafka-grpc-poc/routes"
    "log"
    "os"
)

func main() {
    configType := os.Getenv("CONFIG_TYPE")
    c := config.Factory(config.Type(configType))
    app := container.NewApplication(c)
    r := routes.Factory("all", app)
    if err := r.Expose(); err != nil {
        log.Fatalf("failed to start our http server: %s", err)
        os.Exit(1)
    }
}
