# tetrigo
Puzzle League / Tetris Attack is a great game that hasn't had any love for a long time. As a technical exercise and for fun I want to make a multiplayer online version for the web.

# Contributions
This follows the standard of [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/)

This is to enable the usage of [release-please](https://github.com/googleapis/release-please)

# Requirements
- Golang 1.21 
- Mongo (Optional—Set config.yaml mongo to false if not wanted.)

# Repository Structure
- The main repository consists of the server written in go, and the website written with react.
- `web/` contains the react website.
- The rest is all go and the server.

# Setup 
## Running Golang Project
From the root of the project:
```bash
go run main.go
# Or
make dev
```

## Running React Project
From the `web/` directory:
```bash
# If you haven't ran it
npm install

# Running the website for development
npm run start
```
