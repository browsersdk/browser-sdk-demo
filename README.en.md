# Browser SDK Demo

[![License](https://img.shields.io/github/license/your-org/browser-sdk-demo)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/your-org/browser-sdk-demo)](go.mod)
[![Node.js](https://img.shields.io/badge/node-%3E%3D20.19.0-brightgreen)](https://nodejs.org/)
[![Vue](https://img.shields.io/badge/vue-3.5.27-brightgreen)](https://vuejs.org/)

[中文](./README.md) | **English**

## 🌟 Project Overview

Browser SDK Demo is a comprehensive browser-side SDK demonstration project that includes backend management system, server-side API, and desktop client application. This project showcases best practices of modern web technology stack, suitable for SDK integrators, frontend developers, and platform operators.

## 🏗️ Project Architecture

```
browser-sdk-demo/
├── sdk-admin/     # Backend Management System (Vue 3 + TypeScript)
├── sdk-server/    # Server-side API (Go + Gin)
├── sdk-client/    # Desktop Client (Electron + Vue 3)
```

### Technology Stack Overview

**Frontend Stack:**
- Vue 3 + TypeScript
- Vite Build Tool
- Pinia State Management
- Element Plus UI Library
- Electron Desktop Framework

**Backend Stack:**
- Go Language (1.25+)
- Gin Web Framework
- GORM ORM
- JWT Authentication
- Redis Cache
- MySQL Database

## 🚀 Quick Start

### Environment Requirements

- **Node.js**: >= 20.19.0
- **Go**: >= 1.25
- **pnpm**: Package Manager
- **MySQL**: Database
- **Redis**: Cache Service

### Install Dependencies

```bash
# Clone the project
git clone https://github.com/your-org/browser-sdk-demo.git
cd browser-sdk-demo

# Install frontend dependencies
cd sdk-admin && pnpm install
cd ../sdk-client && npm install

# Install backend dependencies
cd ../sdk-server && go mod tidy
```

### Start Services

**1. Start Backend Service:**
```bash
cd sdk-server
go run main.go start -c resources/config.dev.yaml
```

**2. Start Backend Management:**
```bash
cd sdk-admin
pnpm dev
```

**3. Start Desktop Client:**
```bash
cd sdk-client
npm run dev
```

## 📁 Directory Structure

```
browser-sdk-demo/
├── sdk-admin/              # Backend Management System
│   ├── src/               # Source Code
│   │   ├── components/    # Components
│   │   ├── layout/        # Layout
│   │   ├── router/        # Router
│   │   ├── store/         # State Management
│   │   └── views/         # Page Views
│   ├── package.json       # Dependency Configuration
│   └── vite.config.ts     # Build Configuration
│
├── sdk-server/            # Server Side
│   ├── cmd/               # Command Line Tools
│   ├── common/            # Common Utilities
│   ├── modules/           # Business Modules
│   ├── resources/         # Configuration Resources
│   ├── go.mod             # Go Module
│   └── main.go            # Entry Point
│
├── sdk-client/            # Client
│   ├── src/               # Vue Source
│   ├── electron/          # Electron Main Process
│   ├── package.json       # Dependency Configuration
│   └── vite.config.ts     # Build Configuration
│
└── README.md              # Project Documentation
```

## 🔧 Core Features

### Backend Management System (sdk-admin)
- User authentication and permission management
- Browser environment monitoring
- SDK configuration management
- Data statistics and analysis
- Internationalization support

### Server-side API (sdk-server)
- RESTful API interfaces
- JWT Token authentication
- Database operations
- Log recording
- Middleware support

### Desktop Client (sdk-client)
- Electron desktop application
- Browser environment control
- Real-time status monitoring
- Local data storage
- Cross-platform support

## ⚙️ Configuration Guide

### bro-sdk Configuration

The `bro-sdk` configuration is used to connect to the browser SDK service:

```yaml
bro-sdk:
  endpoint: http://192.168.0.188:9988 # Browser SDK service access address
  api-key: Ls/e7BzPbtHfcmfR3V+kYJpWHsPffdXKWhEYJdABTmYp+CV/FA1G7EHzk6coamX5 # Sensitive information
```

**Configuration Items:**
- `endpoint`: Access address of the browser SDK service
- `api-key`: API key for authentication

**Notes:**
- Please modify the `endpoint` address according to your actual deployment environment
- `api-key` is sensitive information, please keep it secure
- HTTPS protocol is recommended in production environments

### Other Important Configurations

- **Database Configuration**: Supports SQLite and MySQL
- **Cache Configuration**: Supports memory and Redis cache
- **Log Configuration**: Configurable log level and output format
- **JWT Configuration**: Token configuration for user authentication

## 🛠️ Development Guide

### Code Standards

**Frontend Development:**
```bash
# Code checking
pnpm lint

# Code formatting
pnpm format

# Type checking
pnpm type-check
```

**Backend Development:**
```bash
# Code formatting
go fmt ./...

# Code checking
go vet ./...
```

### Build and Deployment

**Build Frontend:**
```bash
# Build backend management
cd sdk-admin && pnpm build

# Build client
cd sdk-client && npm run build
```

**Build Backend:**
```bash
cd sdk-server
go build -ldflags "-w -s" -o bin/server main.go
```

### Testing

```bash
# Frontend testing
cd sdk-admin && pnpm test

# Backend testing
cd sdk-server && go test ./...
```

## 📊 API Documentation

API documentation is automatically generated by Swagger. After starting the service, visit:
```
http://localhost:7888/swagger/index.html
```

## 🤝 Contribution Guidelines

1. Fork the project
2. Create feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to branch (`git push origin feature/AmazingFeature`)
5. Open Pull Request

### Commit Convention

Follow [Conventional Commits](https://www.conventionalcommits.org/) specification:
- `feat`: New features
- `fix`: Bug fixes
- `docs`: Documentation updates
- `style`: Code formatting adjustments
- `refactor`: Code refactoring
- `test`: Test related
- `chore`: Changes to build process or auxiliary tools

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## 📞 Contact

- Project Homepage: [https://github.com/your-org/browser-sdk-demo](https://github.com/your-org/browser-sdk-demo)
- Issues: [https://github.com/your-org/browser-sdk-demo/issues](https://github.com/your-org/browser-sdk-demo/issues)
- Email: your-email@example.com

---

<p align="center">Made with ❤️ by Your Team</p>