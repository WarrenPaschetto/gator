# 🐊 Gator - RSS Aggregator CLI

Gator is a command-line RSS feed aggregator written in Go. It fetches posts from RSS feeds, stores them in a PostgreSQL database, and allows users to interact with feeds, follow blogs, and browse posts—all from the terminal.

---
##📡 About
This project was built to learn how to build real-world command-line apps in Go using:

PostgreSQL

Goose (migrations)

sqlc (safe SQL queries)

XML + RSS parsing

Ticker loops and middleware

---
## 🚀 Requirements

To run Gator, make sure you have:

- **Go** (1.21+ recommended)  
  [Install Go](https://go.dev/doc/install)

- **PostgreSQL** (15+ recommended)  
  [Install PostgreSQL](https://www.postgresql.org/download/)

- **Goose** (for database migrations)  
  Install Goose globally:

```bash
  go install github.com/pressly/goose/v3/cmd/goose@latest
```

---

## 🛠 Installation

Once Go is installed, you can install Gator globally using:

```bash
go install github.com/YOUR_GITHUB_USERNAME/gator@latest
```

**This will create a gator binary in your $GOPATH/bin, which you can run from anywhere.**

---

## 📦 Development Setup
If you're contributing or developing locally:
```bash
git clone https://github.com/YOUR_GITHUB_USERNAME/gator.git
cd gator
go run . <command>
```

---

## ⚙️ Configuration
Before running Gator, you need to create a config file:

Create a .gatorconfig.json in your home directory (this is managed automatically after register)

Inside, it will store your current user and database connection URL:
```json
{
  "db_url": "postgres://postgres:yourpassword@localhost:5432/gator?sslmode=disable",
  "current_user_name": "yourusername"
}
```
Make sure the gator database exists in Postgres:
```bash
createdb gator
```

Then run migrations:
```bash
goose -dir ./sql/schema postgres "your-db-url" up
```

---

## 🧪 Common Commands
```bash
gator register alice            # Create and login as user "alice"
gator login bob                 # Switch to another user
gator addfeed "Go Blog" "https://blog.golang.org/feed.atom"
gator follow "https://hnrss.org/newest"
gator following                 # See all feeds you're following
gator browse 5                  # View recent 5 posts
gator agg 1m                    # Start background feed scraper loop (every 1 minute)
```

## 🛠 Building the Binary
To build the CLI binary manually:
```bash
go build -o gator .
./gator browse
```
Or install globally:
```bash
go install .
```

---

## 🧼 Notes
go run . is great for development

gator (the compiled binary) is what you'd run in production or share with others

To reset the DB during dev: gator reset

