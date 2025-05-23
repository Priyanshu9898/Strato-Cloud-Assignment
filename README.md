# Strato-Cloud-Assignment

![Build Status](https://img.shields.io/github/actions/workflow/status/Priyanshu9898/Strato-Cloud-Assignment/ci.yml?branch=main) ![Coverage Status](https://img.shields.io/codecov/c/github/Priyanshu9898/Strato-Cloud-Assignment?branch=main) ![Commit Activity](https://img.shields.io/github/commit-activity/y/Priyanshu9898/Strato-Cloud-Assignment) ![Issues](https://img.shields.io/github/issues/Priyanshu9898/Strato-Cloud-Assignment) ![Pull Requests](https://img.shields.io/github/issues-pr/Priyanshu9898/Strato-Cloud-Assignment) ![License](https://img.shields.io/github/license/Priyanshu9898/Strato-Cloud-Assignment)

A take-home test for Strato-Cloud: a full-stack demo showing a real-time user table powered by a Go backend and a React frontend.

---

## 🛠️ Tech Stack

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white) ![React](https://img.shields.io/badge/React-20232A?logo=react&logoColor=61DAFB) ![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-38B2AC?logo=tailwind-css&logoColor=white) ![date-fns](https://img.shields.io/badge/date--fns-3178C6?logo=date-fns&logoColor=white) ![Vite](https://img.shields.io/badge/Vite-646CFF?logo=vite&logoColor=white)

---

## 🚀 Quick Start

### Prerequisites

- Go 1.18+
- Node.js 16+ / npm 8+

### Setup

1. **Clone the repo**:
   ```sh
   git clone https://github.com/Priyanshu9898/Strato-Cloud-Assignment.git
   cd Strato-Cloud-Assignment
   ```

2. **Start the backend server**:
   ```sh
   cd server
   go run ./cmd/server
   ```

3. **Start the frontend app**:
   ```sh
   cd client
   npm install
   npm run dev
   ```

4. Open your browser at `http://localhost:3000` to view the live user table.

---

## 📝 Features

- Fetches user data from a Go `/api/users` endpoint
- Calculates **Days Since Last Password Change** and **Days Since Last Access** dynamically
- Displays real-time updates via Server-Sent Events (bonus)
- Filtering by MFA status and highlighting stale users (bonus)

---

## 🧪 Testing & Coverage

![Coverage Status](https://img.shields.io/codecov/c/github/Priyanshu9898/Strato-Cloud-Assignment?branch=main)

### Run Tests

- **Go**:
  ```sh
  cd server
  go test ./... -cover
  ```

- **React** (Jest + React Testing Library):
  ```sh
  cd client
  npm run test -- --coverage
  ```

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/YourFeature`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature/YourFeature`)
5. Open a Pull Request

Please ensure all tests pass and code follows the existing style.

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

*Feel free to adjust badge links and branch names if your CI workflows differ.*
