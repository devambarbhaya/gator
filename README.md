# Gator - The Blog Aggregator

Gator is an RSS feed aggregator built in Go! 🐊 It is a CLI tool that allows users to:

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

RSS feeds are a way for websites to publish updates to their content. You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!

## Prerequisites

Before using Gator, ensure you have the following installed:

- **Go**: Install Go from [golang.org](https://golang.org/dl/)
- **PostgreSQL**: Install PostgreSQL from [postgresql.org](https://www.postgresql.org/download/)

## Installing Gator

To install Gator using `go install`, run the following command:

```sh
go install github.com/devambarbhaya/gator@latest
```

This will install Gator as a binary, which you can run from anywhere.

## Configuring Gator

Gator requires a configuration file to connect to PostgreSQL. Create a `config.yaml` file in your home directory (`~/.gator/config.yaml`) with the following structure:

```yaml
postgres:
  host: "localhost"
  port: 5432
  user: "your-username"
  password: "your-password"
  dbname: "your-database"
user: ""
```

Make sure to replace the placeholders with your actual PostgreSQL credentials.

## Running Gator

Once installed and configured, you can run Gator using:

```sh
gator
```

## User Management Commands

Gator supports user management through the CLI:

```sh
gator login <username>  # Sets the current user in the config
gator register          # Adds a new user to the database
gator users             # Lists all the users in the database
```

## Feed Management Commands

Gator allows users to add and manage RSS feeds.

```sh
gator addfeed "Feed Name" "https://example.com/rss"  # Add an RSS feed
gator feeds  # List all added feeds
```

When adding a feed, it will be linked to the currently logged-in user.

## Following Feeds

Users can follow feeds that other users have added.

```sh
gator follow "https://example.com/rss"  # Follow an existing feed by URL
gator following  # List all feeds the current user is following
gator unfollow "https://example.com/rss"  # Unfollow a feed
```

## Reset Command

For development and testing purposes, Gator provides a `reset` command to clear all users from the database:

```sh
gator reset  # Deletes all users from the database
```

This is useful for resetting the database state while testing, but should not be used in production environments.

## Aggregating RSS Feeds

To test RSS feed parsing, use the `agg` command:

```sh
gator agg 1m  # Fetches new posts every 1 minute
gator browse 5  # View the 5 latest posts from followed feeds
```

**Production**: Build the binary using `go install` and run `gator` directly.

## Submitting Your Repo

Push Gator to GitHub, then submit the link to your remote repo. Your link should look like this:

```
https://github.com/github-username/repo-name
```
