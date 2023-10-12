# NewsCLI - Your News Terminal App

NewsCLI is a Go-based terminal application that allows you to read and summarize news from multiple sources such as Globo, UOL, and Terra. Stay updated without leaving your terminal!

## Features

- Fetch latest news headlines from multiple sources.
- View detailed articles.
- User-friendly command line interface.
- Colorful output for better readability.

## Installation

1. Make sure you have Go installed ([download here](https://golang.org/dl/)).
2. Clone this repository:
    ```bash
    git clone https://github.com/yourusername/newscli.git
    ```
3. Navigate to the project directory:
    ```bash
    cd newscli
    ```
4. Build the project:
    ```bash
    go build
    ```
5. Run the executable:
    ```bash
    ./newscli
    ```

## Usage

- Run the program
    ```bash
    ./newscli
    ```
- Follow the on-screen instructions to select a news source.
- Navigate through the news headlines.
- Select a headline to view the detailed article.

## Dependencies

- [Colly](https://github.com/gocolly/colly) for web scraping.
- [Cobra](https://github.com/spf13/cobra) for CLI functionalities.
- [promptui](https://github.com/manifoldco/promptui) for interactive user prompts.
- [color](https://github.com/fatih/color) for colorful output.

## Contributing

1. Fork the project.
2. Create your feature branch (`git checkout -b feature/fooBar`).
3. Commit your changes (`git commit -am 'Add some fooBar'`).
4. Push to the branch (`git push origin feature/fooBar`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.