# 📦 Secondary Device Sync Storage

A distributed storage system that utilizes your **secondary device** (smartphone, tablet, second laptop, etc.) as temporary or extended storage. This app buffers data locally and automatically syncs it to the cloud once your secondary device comes online. Built with a modern, scalable stack including Flutter, Go, Docker, Kubernetes, and more.

---

## 🚀 Features

- Use *any secondary device* as a storage buffer
- Offline-first design: data is safely buffered when offline
- Automatic sync when the device reconnects to the internet
- Mobile-friendly UI built with Flutter
- Robust backend using Go and MongoDB
- Containerized and orchestrated via Docker & Kubernetes
- Deployed on DigitalOcean

---

## 🛠 Tech Stack

### Frontend
- **Flutter**: Cross-platform mobile & web interface
- **Riverpod**: State management
- **GoRouter**: Declarative routing for Flutter

### Backend
- **Golang**: RESTful API services
- **MongoDB**: Document-based database for scalable data storage

### DevOps & Deployment
- **Docker**: Containerization for backend services
- **Kubernetes**: Orchestration and scaling of services
- **Nginx**: Reverse proxy and load balancing
- **DigitalOcean**: Cloud hosting and deployment

---

## 📦 Architecture Overview

```
[ User Device (Flutter App) ]
        |
        v
[ Go API Server (Docker) ] <-> [ MongoDB ]
        |
        v
[ Nginx Proxy ]
        |
        v
[ Kubernetes Cluster on DigitalOcean ]
```

- Data is temporarily stored in the device buffer.
- Once internet is available, the data is synced via Go APIs.
- MongoDB stores the data permanently.
- Nginx handles incoming traffic and routes it efficiently.
- All services are containerized and deployed via Kubernetes.

---

## 🧪 Getting Started

### Prerequisites
- Flutter SDK
- Go 1.20+
- Docker
- Kubernetes (Minikube or DigitalOcean Kubernetes Cluster)

### Running Locally

#### Frontend
```bash
cd frontend
flutter run
```

#### Backend
```bash
cd backend
go run main.go
```

#### Docker Compose (Optional for local testing)
```bash
docker-compose up --build
```

---

## 🌐 Deployment

1. Build and push Docker images to your container registry.
2. Apply Kubernetes manifests to your DigitalOcean cluster:
```bash
kubectl apply -f k8s/
```
3. Setup Nginx Ingress Controller for routing.

---

## 🤝 Contributing

Contributions, issues and feature requests are welcome! Feel free to fork the repo and submit a PR.

---

## 📄 License

This project is licensed under the MIT License.

---

## 📬 Contact

Feel free to reach out for support or collaboration:
- GitHub: [manab-pr](https://github.com/manab-pr)
- Email: manab2001maity@gmail.com

