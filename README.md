# KurnodeBot [Twitch Chatbot]

A self-hosted twitch chatbot

[![Golang 1.9.4](https://img.shields.io/badge/Go-1.9.4-69D7E2.svg)](https://golang.org/)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

This chatbot is using the native Go TCP package "net" connect to Twitch IRC. To handle the user input used BleveSearch for full text searching. Weather information from Open Weather Map.

Why I will make this project? Since I want to make my own chatbot on my Twitch Channel for fun =P

## Build & Run from source

### Pre-requirment

- Go 1.9.3 or latest (<https://golang.org/>)
- Go dep (Dependency management) <https://github.com/golang/dep>
- A Twitch account
- [Optional] A Open Weather Map account

### Start

1. Clone this project
2. Change directory to project directory
    ```bash
    cd /project/root/dir
    ```
3. Download dependencies
    ```bash
    dep ensure
    ```
4. Modify `config.yml` to fit your requirment
5. Finally, start the chatbot
    ```bash
    go build -o dist/kurnodebot && ./dist/kurnodebot
    ```

## Args

./kurnodebot [-config [config.yml path]]

## CLI commands

```text
send [message]      | Send message to chat room
test-rule [message] | Test rule response (Simulate user message)
stop                | Stop server
help                | Show this help
```

## Config

Here is the configuration example in config.yml

If you not familiar with yaml, you can visit official YAML website <http://www.yaml.org/spec/1.2/spec.html> to learn the syntax, that is a very friendly configuration format =)

```yaml
# Chatbot server setting
server:
  username: <Twitch username>
  password: <Twitch password>
  channel: <Twitch channel>
  # Limit rate for your account, references https://dev.twitch.tv/docs/irc#twitch-irc-capability-membership
  limit-send-count: 100
# External API settings
external:
  # Open weather map
  open-weather-map:
    api-key: <Open weather map api key>
rules:
  # User input message
  - input:
      type: MESSAGE
      content: Hello
    output:
      type: MESSAGE
      content: Hello, {{username}}!
  # User input command (prefix with ~)
  - input:
      type: COMMAND
      content: hi
    output:
      type: MESSAGE
      content: Hello, {{username}}!
  # Show time
  - input:
      type: COMMAND
      content: time
    output:
      type: TIME
      time-format: 2006 Jan _2 15:04:05 MST
      content: "Current time: {{time}}"
  # Show weather using open weather map API
  - input:
      type: COMMAND
      content: weather
    output:
      type: WEATHER_CURRENT
      default-location: Hong Kong
      units: metric
      content: "Current weather in {{name}} - Temp: {{temp}}°C, Min temp: {{temp_min}}°C, Max temp: {{temp_max}}°C"
# Message will be send at a time
schedulers:
  - interval: 120 # Interval in seconds
    output:
      type: TIME
      time-format: 2006 Jan _2 15:04:05 MST
      content: "Current time: {{time}}"
```

### Input types

#### MESSAGE

User send message normally

#### COMMAND

User send a message with `~` prefix

e.g. ~time

### Output types

#### MESSAGE

Send message as config

##### Parameters

| Name    | Type   | Example |
|---------|--------|---------|
| content | string | Hello!  |

##### Placeholders

| Name         | Example     |
|--------------|-------------|
| {{username}} | sender_name |

#### TIME

Show current time from chatbot timezone

##### Parameters

| Name        | Type   | Example                                   |
|-------------|--------|-------------------------------------------|
| content     | string | Current time {{time}}                     |
| time-format | string | 2006 Jan _2 15:04:05 MST (Go time format) |

##### Placeholders

| Name         | Example                  |
|--------------|--------------------------|
| {{time}}     | 2018 Feb 17 09:00:00 HKT |

#### WEATHER_CURRENT

Show current weather

It will change the timezone if a user input the city name after the command

##### Parameters

| Name             | Type   | Example                              |
|------------------|--------|--------------------------------------|
| content          | string | Current weather {{temp}}             |
| default-location | string | Hong Kong                            |
| units            | string | metric (Open weather map units)      |

##### Placeholders

| Name         | Example |
|--------------|---------|
| {{temp}}     | 20      |
| {{temp_min}} | 18      |
| {{temp_max}} | 22      |

## Support Languages

If your server is running in Unicode, any langauges are support =D

## License

Licensed under GPL-3.0

## Authors

- [OnikurYH](https://github.com/OnikurYH)