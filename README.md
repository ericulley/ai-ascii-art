# AI ASCII ART

CS50 Final Project  
Copyright © 2024 [Eric Culley](https://github.com/ericulley)

## Description

This is an interactive CLI application that uses AI to generate ASCII art. It is written in Go and utilizes SQLite to store and manage the data. Additionally, a couple other libraries are used to help with building the text-based user interface, namely Cobra and BubbleTea.

## Installation

1. Clone this repo using `git clone https://github.com/ericulley/ai-ascii-art.git`
2. Move into the directory using `cd ai-ascii-art`
3. Create an .env file with `make env`
   - The .env file should contain `OPENAI_API_KEY` & `OPENAI_MAX_TOKENS` variables
   - Optional: if you have an [OpenAI API key](https://openai.com/index/openai-api/) with a credit balance (may require payment info), you can provide it here
4. Now run `make install` which will
   - build the binary
   - update its permissions to make it executable (may require password)
   - move the executable into /usr/local/bin
5. Confirm installation by running `ascii` to see the help menu

## Usage

Start by running `ascii create` to open up a prompt window with ChatGPT, and ask it to generate some ASCII art for you. For example, you can try entering the following prompt: `Hi ChatGPT, can you create an ASCII art dog?`

> **_NOTE:_** If you have not added an OpenAI API key to the .env file, a default piece of ASCII art will be returned so that you can still test out the commands of the application.

When art is generated and displayed, you will be asked if you'd like to save the art or not. Try creating a few pieces of art and saving them. Then you can use some of the other commands, like `ascii list` to display your saved ascii art, `ascii update` to update the name of saved art, or `ascii delete` to delete some pieces. Make sure to look into each commands' required flags in the help menu (e.g. `ascii <command> --help`) to utilize each one properly.

Lastly, use the best command, `ascii art`, to display a random piece of saved ASCII art whenever you need a pick-me-up. Happy coding!
