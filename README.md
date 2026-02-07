# Arc

**Arc** is a powerful, cross-platform CLI tool for cloud migrations and automation. Built with Go, it provides unlimited runtime (no Apps Script limits), native Google API access, and seamless transfers between cloud providers.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Features

- 🚀 **Multi-Cloud Support** - Google Drive, AWS S3, Dropbox, OneDrive, and local filesystem
- 📁 **Bulk Operations** - Rename, move, and migrate thousands of files
- 📊 **Analytics** - File statistics by type (images, videos, documents)
- 🔍 **Smart Search** - Find files by type, size, date, or name pattern
- 🔄 **Duplicate Detection** - Find and clean up duplicate files
- ⚡ **No Limits** - Unlike Apps Script, no 6-minute timeout
- 🔐 **Secure OAuth** - Native OAuth 2.0 authentication

## Installation

### From Source

```bash
git clone https://github.com/atheeque/arc.git
cd arc
go build -o arc ./cmd/arc
```

### Homebrew (coming soon)

```bash
brew install arc
```

## Quick Start

### 1. Authenticate with Google

```bash
# Just run this - no setup required!
arc auth login google
```

Your browser will open automatically. Sign in with your Google account and click "Allow".

### 2. List Files

```bash
arc ls google:/                    # List root of Google Drive
arc ls google:/Documents           # List a folder
arc ls local:~/Downloads           # List local directory
arc ls google:/ -l                 # Long listing with details
```

### 3. View Statistics

```bash
arc stats google:/                 # See file counts by type
```

Example output:
```
┌──────────────────────────────────────────┐
│ Google Statistics                        │
├──────────────────────────────────────────┤
│ Total Files:     12,456                  │
│ Total Folders:   1,234                   │
│ Total Size:      45.6 GB                 │
│                                          │
│ By Type:                                 │
│   📷 Images:     4,523 files (15.2 GB)   │
│   🎥 Videos:       234 files (20.1 GB)   │
│   📄 Documents:  5,678 files (8.3 GB)    │
│   📁 Other:      2,021 files (2.0 GB)    │
└──────────────────────────────────────────┘
```

## Commands

### File Operations

```bash
arc ls <path>                      # List files
arc cp <src> <dest>                # Copy files
arc mv <src> <dest>                # Move/rename files
```

### Bulk Operations

```bash
# Bulk rename with patterns
arc rename "google:/photos/*" --pattern "{date}_{name}.{ext}"
arc rename "google:/docs/*" --prefix "2024_" --dry-run

# Find files
arc find google:/ --type image
arc find google:/ --size ">100MB" --type video
```

### Duplicate Detection

```bash
arc duplicates google:/            # Find duplicates by hash
arc duplicates google:/ --delete --keep newest --dry-run
```

### Authentication

```bash
arc auth login google              # Login to Google
arc auth status                    # Check auth status
arc auth logout                    # Logout
```

## Path Syntax

Arc uses a unified path syntax: `provider:/path`

| Provider | Example Path |
|----------|-------------|
| Google Drive | `google:/Documents/file.pdf` |
| AWS S3 | `s3://bucket-name/prefix/` |
| Dropbox | `dropbox:/folder/file.txt` |
| OneDrive | `onedrive:/Documents/` |
| Local | `local:/Users/me/file.txt` or `~/file.txt` |

## Google Authentication

Arc comes with built-in OAuth credentials - **no setup required!** Just run:

```bash
arc auth login google
```

### Using Custom OAuth Credentials (Optional)

If you prefer to use your own OAuth credentials:

1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create a project and enable **Google Drive API**
3. Create OAuth credentials (Desktop application)
4. Set environment variables:

```bash
export ARC_GOOGLE_CLIENT_ID="your-client-id.apps.googleusercontent.com"
export ARC_GOOGLE_CLIENT_SECRET="your-client-secret"
arc auth login google
```

## Configuration

Arc stores configuration in `~/.arc/`:

```
~/.arc/
├── config.yaml      # Global configuration
├── google.json      # Google OAuth tokens
├── aws.json         # AWS credentials
└── jobs/            # Running/completed jobs
```

### config.yaml

```yaml
default_provider: google
concurrency: 10
verbose: false

providers:
  google:
    client_id: "your-client-id"
    client_secret: "your-client-secret"
  aws:
    region: us-east-1
    profile: default
```

## Roadmap

- [x] Google Drive support
- [x] Local filesystem support
- [x] Bulk rename
- [x] File statistics
- [x] Duplicate detection
- [ ] AWS S3 support
- [ ] Dropbox support
- [ ] OneDrive support
- [ ] Background jobs
- [ ] Web UI

## Contributing

Contributions are welcome! Please read our contributing guidelines first.

```bash
# Clone and build
git clone https://github.com/atheeque/arc.git
cd arc
go build -o arc ./cmd/arc

# Run tests
go test ./...
```

## License

MIT License - see [LICENSE](LICENSE) for details.
