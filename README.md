# Blog Aggregator CLI (Gator)

Gator is a versatile command-line interface (CLI) tool designed to help you manage and aggregate blog feeds efficiently. Whether you're a developer or a content enthusiast, Gator streamlines the process of handling RSS feeds, user management, and more.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
    - [Project Directory](#project-directory-configuration)
    - [Home Directory](#home-directory-configuration)
- [Usage](#usage)
    - [Running the Program](#running-the-program)
    - [Available Commands](#available-commands)
- [Contributing](#contributing)
- [License](#license)
- [Getting Started](#getting-started)

## Prerequisites
Before you begin, ensure you have the following software installed on your system:

### Go
Gator is built using the Go programming language. You need to have Go installed to run and manage the CLI.

- **Installation Instructions:**
  You can download and install Go from the [official website](https://golang.org/dl/). Follow the installation guide suitable for your operating system.

### PostgreSQL
Gator relies on PostgreSQL for data storage. Ensure you have PostgreSQL installed and running.

- **Installation Instructions:**
  Download PostgreSQL from the [official website](https://www.postgresql.org/download/) and follow the installation steps for your operating system.

## Installation
To install the Gator CLI, use the `go install` command:
This command will download the latest version of Gator and install the executable in your `$GOPATH/bin` directory. Ensure that this directory is included in your system's `PATH` to run Gator from any location.

## Configuration
Before running Gator, you need to set up the configuration file to connect to your PostgreSQL database.

### Project Directory
1. **Create Configuration File:**
   Gator expects a configuration file (e.g., `config.yaml`) in the project directory. Create this file with the following structure:

   ```yaml
   db_url: "postgres://username:password@localhost:5432/gator_db?sslmode=disable"
   current_user: "your_username"
   ```

  - **db_url:** Replace `username`, `password`, `localhost`, `5432`, and `gator_db` with your PostgreSQL credentials and database details.
  - **current_user:** Set this to your desired default username.

2. **Ensure Proper Permissions:**
   Make sure the configuration file has the correct permissions to prevent unauthorized access to sensitive information.

### Home Directory
Alternatively, you can manually create a configuration file in your home directory.

1. **Create Configuration File:**
   Manually create a config file in your home directory, `~/.gatorconfig.json`, with the following content:

   ```json
   {
     "db_url": "postgres://example"
   }
   ```

  - **db_url:** Replace `"postgres://example"` with your actual PostgreSQL connection string.

2. **Set Environment Variable (Optional):**
   If you prefer to use the home directory configuration, ensure that Gator points to this file. You might need to set an environment variable or modify the application's configuration loading logic to prioritize `~/.gatorconfig.json`.

## Usage
Once installed and configured, you can start using Gator to manage your blog feeds and users.

### Running the Program
Navigate to the directory containing the `config.yaml` file (if using project directory configuration) or ensure your home directory configuration is set up, then execute the CLI:
### Available Commands
Here are some of the essential commands you can run with Gator:

- **Register a New User**
  Create a new user account.

  ```shell script
  gator register <username>
  ```

  **Example:**

  ```shell script
  gator register john_doe
  ```

- **Login as a User**
  Switch to an existing user account.

  ```shell script
  gator login <username>
  ```

  **Example:**

  ```shell script
  gator login john_doe
  ```

- **Add a New Feed**
  Add a new RSS feed to your account.

  ```shell script
  gator addfeed <feed_name> <feed_url>
  ```

  **Example:**

  ```shell script
  gator addfeed TechNews https://example.com/rss
  ```

- **List All Feeds**
  Display all the feeds you have added.

  ```shell script
  gator feeds
  ```

- **Browse Posts**
  Browse the latest posts from your subscribed feeds. You can specify the number of posts to display.

  ```shell script
  gator browse [limit]
  ```

  **Example:**

  ```shell script
  gator browse 10
  ```

- **Reset the Database**

  **Warning:** This command will delete all user data.

  ```shell script
  gator reset
  ```

- **Aggregate Feeds Automatically**
  Collect feeds at regular intervals.

  ```shell script
  gator agg <time_between_requests>
  ```

  **Example:**

  ```shell script
  gator agg 1h
  ```

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request with your enhancements.

## License
This project is licensed under the [MIT License](LICENSE).

## Getting Started
1. **Clone the Repository:**

   ```shell script
   git clone https://github.com/marcuschui2022/blog-aggregator.git
   cd blog-aggregator
   ```

2. **Install Dependencies:**

   ```shell script
   go mod download
   ```

3. **Run the CLI:**

   ```shell script
   go run main.go <command> [arguments]
   ```

   Replace `<command>` and `[arguments]` with your desired commands as outlined above.

---
For more detailed information and advanced usage, please refer to the [official documentation](https://github.com/marcuschui2022/blog-aggregator).