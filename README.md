# Console

Console is a library designed for Dragonfly servers, enabling the use of commands through command-line
## Getting Started

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
