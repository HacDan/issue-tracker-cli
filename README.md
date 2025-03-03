# 📌 Project: Terminal-Based Issue Tracker CLI
## Tech Stack:
* Language: Golang
* Database: SQLite 
* CLI Framework: Cobra

## 🚀 Key Features:

### 🔹 Core Functionality
* Create, update, delete issues (with priority, status, tags)
* List and filter issues (by status, priority, tag, assignee)
* Search for issues using fuzzy matching
* Assign issues to users

### 🔹 Enhancements
* Sync with a remote PostgreSQL database
* Export issues as JSON/CSV for reporting
* Offline mode with local caching and auto-sync

### 🔹 Bonus Features
* GitHub/Jira integration (fetch issues from repositories)
* Webhook support for notifications
* Terminal UI mode for an interactive interface

## 📌 Basic Commands
### ✅ Create an Issue

```sh
issue add -t "Fix login bug" -d "Users cannot log in with special characters" -p high -s open -a alice
```

 * -t: Title
 * -d: Description
 * -p: Priority (low, medium, high, critical)
 * -s: Status (open, in-progress, closed)
 * -a: Assign to user

### 📝 List Issues
```sh
issue list 
issue list -s open               # Show only open issues  
issue list -p high               # Show high-priority issues  
issue list -a alice              # Show issues assigned to Alice  
```

### 🔍 Search for Issues
```sh
issue search "login bug"
issue search -t "UI glitch"       # Search by title  
issue search -d "performance"     # Search in description  
```

### ✏️ Edit an Issue
```sh
issue edit 42 -t "Fix login error" -p critical -s in-progress
```

* 42 is the issue ID
* Updates the title, priority, and status

### ❌ Delete an Issue
```sh
issue delete 42
issue delete --all-closed       # Delete all closed issues  
```

~~### 📌 Tagging, Filtering, and Organizing~~ Future plans
#### 🏷️ Tag an Issue
```sh
issue tag 42 -a "backend" -a "security"
```
 * -a: Add tags
 * Use -r to remove a tag

```sh
issue tag 42 -r "security"
```

### 📌 Change Issue Status
```sh
issue close 42
issue reopen 42
issue move 42 -s in-progress
```

### 📌 Sync & Export
#### 🔄 Sync with Remote Database
```sh
issue sync          # Sync local issues with remote PostgreSQL  
issue sync --push   # Force push local changes  
issue sync --pull   # Fetch latest issues from remote  
```

### 📤 Export Issues
```sh
issue export -f json > issues.json
issue export -f csv > issues.csv
```

### 📌 Advanced (Optional)
#### 🔗 Fetch Issues from GitHub/Jira
```sh
issue import github --repo "org/project"
issue import jira --project "DEV"
```
