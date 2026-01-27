package main

import "go.uber.org/zap"

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("This is an info message.")

	logger.Info("User logged in", zap.String("username", "John Doe"), zap.String("method", "GET"))

}
