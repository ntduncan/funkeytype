# FunkeyType

A terminal-based typing test application written in Go.

## Features

-   A typing test with a variety of words.
-   Calculates your words per minute (WPM) and accuracy.
-   Keeps track of your top score in each mode.
-   Simple and intuitive user interface.

## Installation
```
brew tap ntduncan/ntduncan
brew install --cask funkeytype
```

## Usage

Run the application from your terminal:

```
funkeytype
```

## Configuration

FunkeyType creates a configuration file to store your top high scores.

-   **File Name:** `.scores.json`
-   **Location:** The file is created in your home directory (`~/.config/funkeytype/scores.json`).

The file is automatically created on your first run.

## License

This project is licensed under the MIT License.
