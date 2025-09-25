# BE-GO-GIN

Backend service with Go and Gin framework.

---

## 📦 Project Structure
```
---
BE-RAINVOW/
├── cmd/ # HTTP handler functions
│ └── web
├── internal/ 
│ └── config
│ └── delivery
├── go.mod
├── go.sum
└── .gitignore
---
```

## 🚀 Getting Started

### 1. Install Go
```bash
brew install go     # for Mac
```

### 2. Clone the repo
```bash
git clone https://github.com/RifanAmrozi/be-rainvow.git
cd be-rainvow
go mod tidy
```

### 3. Run the project

```bash
go run cmd/web/main.go 
```
API will be available at: http://localhost:8080/api/v1/

### ✅ Git Commit Convention
https://www.conventionalcommits.org/

Use Conventional Commits example:
feat: add user list endpoint
fix: correct seed data typo
chore: update dependencies

### Pull Request Flow
#### Create new branch:
```
git checkout -b feat/user-endpoint
```

#### Commit changes:
```
git add .
git commit -m "feat: add GET /user endpoint"
```

#### Push and create PR:
```
git push origin feat/user-endpoint
```

#### Open pull request on GitHub, assign reviewer.


### Rebase before merge
```
git checkout feat/user-endpoint
git fetch origin
git rebase origin/main
# resolve any conflicts, then:
git push --force
```