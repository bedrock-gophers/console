# Console

Console is a library designed for Dragonfly servers, providing enhanced console logging functionality and management. With Console, you can easily enable and customize console logging for your server, allowing for improved debugging and monitoring capabilities.
Getting Started

To enable Console for your Dragonfly server, follow these simple steps:

 - Initialize a new Logger instance from the Logrus library.
    ```go
        log := logrus.New()
        log.Formatter = &logrus.TextFormatter{ForceColors: true}
        log.Level = logrus.DebugLevel
    ```
   - Enable the console with the initialized Logger instance.
    ```go
        console.Enable(log)
    ```
