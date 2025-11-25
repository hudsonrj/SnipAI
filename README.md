![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) ![SQLite](https://img.shields.io/badge/SQLite-003B57?style=for-the-badge&logo=sqlite&logoColor=white) ![License](https://img.shields.io/badge/license-MIT-green?style=for-the-badge) ![Version](https://img.shields.io/badge/version-1.1.0-blue?style=for-the-badge) ![GitHub stars](https://img.shields.io/github/stars/matheuzgomes/snip?style=for-the-badge&label=Stars)




<div align="center" style="margin-bottom: 15px; display: flex; align-items: center; justify-content: center; gap: 15px;">
  <img src="assets/snip_logo.png" alt="Snip Logo" width="120" height="130" style="border-radius: 16px; border: 2px solid #e0e0e0;">
  <h1 style="margin: 0;">Snip</h1>
</div>

A fast and efficient command-line note-taking tool built with Go. Snip helps you capture, organize, and search your notes with AI-powered features, project management, tasks, and checklists.

## 🎬 Demo

![Snip Demo](assets/snip_demo.gif)

## ✨ Features

### 📝 Notes Management

- **Create Notes**: Quickly create new notes with title and content
- **List Notes**: View all your notes with chronological sorting options
- **Search Notes**: Full-text search across all notes using SQLite FTS4
- **Edit Notes**: Update existing notes using your preferred editor
- **Get Notes**: Retrieve specific notes by ID with markdown rendering support
- **Delete Notes**: Remove notes you no longer need
- **Tags**: Organize notes with custom tags
- **Patch Notes**: Update note titles and manage tags
- **Export Notes**: Export notes to JSON and Markdown formats
- **Import Notes**: Import notes (markdown) from files and directories
- **Markdown Preview**: Render markdown content beautifully in the terminal
- **Fast Performance**: SQLite database with optimized indexes (90-127ns operations)
- **Editor Integration**: Supports nano, vim, vi, or custom `$EDITOR`
- **Comprehensive Testing**: Full test coverage with performance benchmarks

### 🤖 AI-Powered Features

- **AI Create Notes**: Generate notes with AI-powered content based on topics
- **AI Code Generation**: Generate code in multiple languages with AI
- **AI Search Enhancement**: Improve search queries using AI
- **AI Q&A**: Ask questions to AI based on your notes context
- **AI Project Planning**: Generate detailed project plans with AI
- **AI Checklist Generation**: Create checklists with AI-generated items

### 📁 Project Management

- **Projects**: Create and manage projects with descriptions and status
- **Tasks**: Create tasks within projects with priorities and due dates
- **Task Status**: Track tasks (pending, in_progress, completed)
- **Task Priorities**: Set task priorities (low, medium, high)
- **Checklists**: Create checklists for projects or tasks
- **Checklist Items**: Manage checklist items with completion tracking
- **Progress Tracking**: Visual progress indicators for checklists

### Command Examples

#### 📝 Basic Notes

```bash
# Create a new note
snip create "Meeting Notes"

# Create a new note quickly
snip create "World" --message "Hello!"

# List all notes (newest first)
snip list

# List notes chronologically (oldest first)
snip list --asc

# List with verbose information
snip list --verbose

# Search for notes containing specific terms
snip find "meeting"

# Edit an existing note
snip update 1

# Get a specific note by ID
snip show 1

# Get a note with markdown rendering
snip show 1 --render

# Delete a specific note by ID
snip delete 1

# Patch/update a note's title
snip patch 1 --title "New Title"

# Patch/update a note's tags
snip patch 1 --tag "work important"

# List notes with tags
snip list --tag "work"

# Export notes to JSON format
snip export --format json

# Export notes to Markdown format
snip export --format markdown

# Export notes created since a specific date
snip export --since "2024-01-01"

# Import notes from a directory
snip import /path/to/notes/directory

# Show editor information and available options
snip editor
```

#### 🤖 AI Features

```bash
# Create a note with AI-generated content
snip ai-create "Python Decorators" --tag "programming"

# Generate code with AI
snip ai-code "function to reverse a string" --lang "python"

# Improve search query with AI
snip ai-search "meeting notes"

# Ask questions to AI based on your notes
snip ai-ask "What did I write about Python?"
```

#### 📁 Project Management

```bash
# Create a project
snip project create "Web Application" --description "New web app project"

# Create a project with AI-generated plan
snip project ai-create "Mobile App" --description "iOS and Android app"

# List all projects
snip project list

# Show project details with tasks
snip project show 1

# Update project
snip project update 1 "Updated Name" --status "active"

# Delete a project
snip project delete 1
```

#### ✅ Tasks

```bash
# Create a task
snip task create "Implement authentication" --project 1 --priority high --due 2025-12-15

# List all tasks
snip task list

# List tasks for a specific project
snip task list --project 1

# List tasks by status
snip task list --status pending

# Show task details
snip task show 1

# Update a task
snip task update 1 "New Title" --status in_progress --priority medium

# Toggle task completion
snip task toggle 1

# Delete a task
snip task delete 1
```

#### 📋 Checklists

```bash
# Create a checklist
snip checklist create "Deployment Checklist" --project 1

# Create a checklist with AI-generated items
snip checklist ai-create "Pre-launch Checklist" --items 10 --project 1

# List all checklists
snip checklist list

# List checklists for a project
snip checklist list --project 1

# Show checklist with progress
snip checklist show 1

# Add item to checklist
snip checklist item-add 1 "Test database connection"

# Toggle checklist item completion
snip checklist item-toggle 5

# Delete checklist item
snip checklist item-delete 5

# Delete a checklist
snip checklist delete 1
```

## 🚀 Installation

### Package Managers

#### Scoop (Windows)
```bash
# Add the bucket
scoop bucket add snip https://github.com/matheuzgomes/Snip

# Install snip
scoop install snip

# Update snip
scoop update snip
```

#### Homebrew (macOS/Linux)
```bash
# Add the tap
brew tap matheuzgomes/homebrew-Snip

# Install snip
brew install --cask snip-notes

# Update snip
brew upgrade --cask snip-notes
```

**⚠️ macOS Security Note:**

If macOS blocks the app with "cannot be opened because the developer cannot be verified":

```bash
# Option 1: Remove quarantine attribute
xattr -d com.apple.quarantine /opt/homebrew/bin/snip

# Option 2: Allow in System Settings
# Go to: System Settings > Privacy & Security > Allow "snip"
```

### Direct Download

Pre-compiled binaries are available in the [releases](https://github.com/matheuzgomes/Snip/releases) page for:
- **Linux**: AMD64 and ARM64
- **Windows**: AMD64

### From Source

#### Prerequisites

- **Go 1.21 or later** - [Download Go](https://go.dev/dl/)
- **SQLite3 development libraries** (for CGO builds)
  - Windows: Included with Go or install via [SQLite](https://www.sqlite.org/download.html)
  - Linux: `sudo apt-get install libsqlite3-dev` (Debian/Ubuntu) or `sudo yum install sqlite-devel` (RHEL/CentOS)
  - macOS: Usually pre-installed or via Homebrew: `brew install sqlite`

#### Compilation

```bash
# Clone the repository
git clone https://github.com/hudsonrj/SnipAI.git
cd SnipAI

# Download dependencies
go mod download

# Build for your platform
go build -o snip.exe main.go

# For Windows (explicit)
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=1
go build -o snip.exe main.go

# For Linux
go build -o snip main.go

# For macOS
go build -o snip main.go

# Install to system path (Linux/macOS)
sudo mv snip /usr/local/bin/
```

#### Windows Build Notes

If you encounter issues running `snip.exe` directly, you can use:

```powershell
# Option 1: Use go run
go run main.go --help

# Option 2: Create an alias in PowerShell profile
# Add to $PROFILE:
function snip { 
    Set-Location "C:\repositorio\SnipAI\SnipAI"
    go run main.go $args
}
```

## 🗄️ Data Storage

Snip stores your notes in a SQLite database located at `~/.snip/notes.db`. The database includes:

- **Main Table**: Stores notes with metadata (ID, title, content, timestamps)
- **Tags Table**: Stores custom tags for organizing notes
- **Notes-Tags Table**: Many-to-many relationship between notes and tags
- **FTS Table**: Full-text search index for fast searching
- **Automatic Triggers**: Keeps search index synchronized with your notes

## 🔧 Configuration

### 🤖 AI Configuration (Groq API)

To use AI-powered features, you need to configure the `GROQ_API_KEY` environment variable.

#### Get Your API Key

1. Visit [Groq Console](https://console.groq.com/keys)
2. Sign up or log in
3. Generate a new API key
4. Copy the key

#### Set Environment Variable

**Windows (PowerShell):**
```powershell
# Temporary (current session only)
$env:GROQ_API_KEY="your_api_key_here"

# Permanent (add to user profile)
[Environment]::SetEnvironmentVariable("GROQ_API_KEY", "your_api_key_here", "User")
```

**Windows (CMD):**
```cmd
# Temporary
set GROQ_API_KEY=your_api_key_here

# Permanent: Control Panel > System > Advanced Settings > Environment Variables
```

**Linux/macOS:**
```bash
# Temporary
export GROQ_API_KEY="your_api_key_here"

# Permanent (add to ~/.bashrc or ~/.zshrc)
echo 'export GROQ_API_KEY="your_api_key_here"' >> ~/.bashrc
source ~/.bashrc
```

**Verify Configuration:**
```bash
# Windows PowerShell
echo $env:GROQ_API_KEY

# Linux/macOS
echo $GROQ_API_KEY
```

For detailed instructions, see [README_API_KEY.md](README_API_KEY.md).

### Editor Selection

Snip automatically detects your preferred editor with cross-platform support:

**Windows:**
- Visual Studio Code, Notepad++, Sublime Text, Atom, Micro, Nano, Vim, Notepad

**macOS:**
- Visual Studio Code, Sublime Text, Atom, Nano, Vim, Vi, Open

**Linux:**
- Nano, Vim, Vi, Micro, Visual Studio Code

**Priority Order:**
1. `$EDITOR` environment variable
2. Platform-specific editor detection
3. Smart fallback to basic editors

**Check Available Editors:**
```bash
snip editor
```

### Database Location

The database is automatically created at `~/.snip/notes.db`. The database includes:

- **Notes Table**: Your notes with metadata
- **Tags Table**: Custom tags
- **Projects Table**: Project information
- **Tasks Table**: Task details
- **Checklists Table**: Checklist definitions
- **Checklist Items Table**: Individual checklist items
- **FTS Table**: Full-text search index

You can backup your data by copying the `~/.snip/notes.db` file.

## 🛠️ Development

### Prerequisites

- Go 1.21 or later
- SQLite3 development libraries (for CGO builds)
- mingw-w64 (for Windows cross-compilation)

### Building

```bash
git clone https://github.com/matheuzgomes/Snip.git
cd Snip
go mod download
go build -o snip main.go
```

### Running Tests

```bash
# Run all tests
make test

# Run performance benchmarks
make bench

# Run tests with verbose output
go test -v ./internal/test/...
```

## 🗺️ Roadmap

### ✅ Completed Features

- ~~**🗑️ Delete Notes**: Remove notes you no longer need~~ ✅ Done!
- ~~**🏷️ Tags**: Organize notes with custom tags~~ ✅ Done!
- ~~**✏️ Patch Notes**: Update note titles and manage tags~~ ✅ Done!
- ~~**📤 Export**: Export notes to various formats (Markdown, JSON, etc.)~~ ✅ Done!
- ~~**📥 Import**: Import notes from files and directories~~ ✅ Done!
- ~~**🧪 Testing**: Comprehensive test suite with benchmarks~~ ✅ Done!
- ~~**🖼️ Markdown Preview**: Visualize rendered Markdown so you can see your notes as they'd appear formatted~~ ✅ Done!
- ~~**🤖 AI Features**: AI-powered note creation, code generation, search enhancement, and Q&A~~ ✅ Done!
- ~~**📁 Project Management**: Create and manage projects with tasks and checklists~~ ✅ Done!
- ~~**✅ Checklists**: Create checklists with AI-generated items and track progress~~ ✅ Done!

### Performance Metrics

Snip v1.1.0 delivers exceptional performance:

- **⚡ Sub-microsecond Operations**: Core operations run in 90-127 nanoseconds
- **💾 Memory Efficient**: Only 56 bytes per operation with 3 allocations
- **🧪 100% Test Coverage**: Comprehensive test suite with performance benchmarks
- **📊 Benchmarking**: Built-in performance monitoring with `make bench`

### Release Automation

We're using [GoReleaser](https://goreleaser.com/) for:

- ✅ **Automated Builds**: Cross-platform binary generation (Linux AMD64/ARM64, Windows AMD64)
- ✅ **Release Management**: Automated GitHub releases
- ✅ **Package Distribution**: Scoop, Homebrew, and Winget package managers
- ✅ **Cross-compilation**: Windows binaries built with mingw-w64
- ✅ **CGO Support**: SQLite integration with proper CGO compilation
- ✅ **CI/CD Pipeline**: Automated testing and release pipeline

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI functionality
- Uses [SQLite](https://sqlite.org/) with FTS4 for fast text search
- Inspired by modern note-taking tools and CLI utilities

**Made with ❤️ for anyone who wants to take notes**
