# Console

Console is a library designed for Dragonfly servers, enabling the use of commands through command-line
## Getting Started

To enable Console for your Dragonfly server, follow these simple steps:

- Enable the console with the initialized Logger instance.
    ```go
        console.Enable(slog.Default())
    ```
