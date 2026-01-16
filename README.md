# 🎬 Movie API

Une API REST robuste pour gérer une collection de films, construite avec **Go**, **PostgreSQL** et **Docker**.
Ce projet démontre une architecture backend moderne, sécurisée et prête pour la production.

![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go) ![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=flat&logo=docker) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-336791?style=flat&logo=postgresql) ![Go Audit](https://github.com/vfaust1/movie-api/actions/workflows/audit.yml/badge.svg)

## 🚀 Fonctionnalités

* **CRUD Complet** : Création, Lecture, Mise à jour, Suppression de films.
* **Base de Données Relationnelle** : Modèle complexe avec relation *Many-to-Many* (Films ↔ Genres).
* **Recherche Avancée** : Filtrage par titre, tri dynamique et pagination (`Metadata`).
* **Sécurité** : Authentification via API Key (Middleware personnalisé).
* **Architecture** : Structure modulaire `cmd/internal` respectant les standards Go.
* **Résilience** : Gestion des *Race Conditions* au démarrage avec Docker (Retry Logic).

## 🛠️ Stack Technique

* **Langage** : Go (Golang) 1.23
* **Base de données** : PostgreSQL
* **Driver SQL** : `pgx` (Performance & Sécurité)
* **Infrastructure** : Docker & Docker Compose
* **Router** : Standard library `net/http` (ServeMux 1.22+)

## 📂 Architecture du Projet

Voici comment le code est organisé, suivant les standards "Project Layout" de Go :

```text
movie-api/
├── .github/
│   └── workflows/
│       └── audit.yml       # Pipeline CI (GitHub Actions)
├── cmd/
│   └── api/
│       ├── handlers.go     # Contrôleurs HTTP
│       ├── main.go         # Point d'entrée
│       ├── middleware.go   # Sécurité et logs
│       └── routes.go       # Définition des URLs
├── docs/                   # Documentation générée par Swagger
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   └── store/
│       ├── db.go           # Connexion à la base de données PostgreSQL
│       ├── movies.go       # Logique métier des films
│       └── movies_test.go  # Fichier test unitaire pour movies.go
├── .dockerignore           # Fichiers ignorés par Docker
├── .env.example            # Variables d'environnement (Template)
├── .gitignore              # Fichiers ignorés par Git
├── docker-compose.yml      # Configuration des conteneurs
├── Dockerfile              # Recette de construction de l'image API
├── go.mod                  # Gestionnaire de dépendances
└── README.md               # Ce fichier
```

## 📦 Installation & Démarrage (Docker)

La méthode recommandée. Aucun outil (Go/Postgres) n'est nécessaire sur votre machine, juste Docker.

1.  **Cloner le dépôt**
    ```bash
    git clone https://github.com/vfaust1/movie-api.git
    cd movie-api
    ```

2.  **Lancer l'environnement**
    ```bash
    docker compose up --build
    ```
    *L'API sera accessible sur `http://localhost:8080`.*

## 🔌 Utilisation de l'API

### Authentification
Toutes les requêtes doivent inclure le header suivant :
`Authorization: Bearer super-secret-password-123`

### Exemples de Routes

| Méthode | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/health` | Vérifier l'état du service |
| `GET` | `/movies` | Lister les films (paginé) |
| `GET` | `/movies?title=dune` | Rechercher un film |
| `POST` | `/movies` | Ajouter un film |
| `GET` | `/movies/{id}` | Détails d'un film |
| `PUT` | `/movies/{id}` | Modifier un film |
| `DELETE` | `/movies/{id}` | Supprimer un film |

---
*Projet personnel réalisé dans le but d'apprendre les bases du langage Go et le fonctionnement d'une API REST. Le projet intègre une base de données PostgreSQL via Docker, une documentation automatique avec Swagger, et un pipeline d'intégration continue (CI) via GitHub Actions.*