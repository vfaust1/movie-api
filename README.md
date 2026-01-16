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
* **Base de données** : PostgreSQL (Driver `pgx` pour la performance)
* **Infrastructure** : Docker & Docker Compose
* **Qualité & CI** : GitHub Actions, Tests Unitaires (Mocks), Linter (`go vet`)
* **Architecture** : Dependency Injection & Repository Pattern

## 📂 Architecture du Projet

Le code suit les standards "Project Layout" de Go et une **Clean Architecture** :

```text
movie-api/
├── .github/
│   └── workflows/
│       └── audit.yml       # Pipeline CI (GitHub Actions)
├── cmd/
│   └── api/
│       ├── handlers.go     # Contrôleurs HTTP
│       ├── handlers_test.go # Tests unitaires (Mocking)
│       ├── main.go         # Point d'entrée & Injection de dépendances
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
│       ├── movies_test.go  # Tests d'intégration DB
│       └── storage.go      # Interfaces (Contrats) pour le découplage
├── .dockerignore           # Fichiers ignorés par Docker
├── .env.example            # Variables d'environnement (Template)
├── .gitignore              # Fichiers ignorés par Git
├── docker-compose.yml      # Configuration des conteneurs
├── Dockerfile              # Recette de construction de l'image API
├── go.mod                  # Gestionnaire de dépendances
└── README.md               # Ce fichier
```

### Zoom sur l'Architecture (Injection de Dépendances)
L'application n'utilise pas de variables globales pour la base de données. À la place, elle utilise une interface MovieRepository.

En Production : On injecte la vraie structure qui parle à PostgreSQL.

En Test : On injecte un Mock (fausse BDD) pour tester l'API instantanément sans Docker.

## 📦 Installation & Démarrage (Docker)

Aucun outil (Go/Postgres) n'est nécessaire sur votre machine, juste Docker.

1.  **Cloner le dépôt**
    ```bash
    git clone https://github.com/vfaust1/movie-api.git
    cd movie-api
    ```

2.  **Lancer l'environnement**
    ```bash
    docker compose up --build
    ```
    *👉 Une fois lancé, accédez à la documentation interactive : `http://localhost:8080/swagger/index.html`*

## Tests Unitaires

Grâce à l'architecture découplée, les tests s'exécutent en mémoire.

```bash
# Lancer les tests (nécessite Go installé localement)
go test -v ./cmd/api/
``` 

## 🔌 Utilisation de l'API

### Authentification

L'API utilise une authentification par **Bearer Token**.
* **Lecture (GET)** : Accès public (pas de token requis).
* **Écriture (POST, PUT, DELETE)** : Requiert le header suivant :
    `Authorization: Bearer super-secret-password-123`

### Exemples de Routes

| Méthode | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/movies` | Lister les films (paginé) |
| `GET` | `/movies?title=dune` | Rechercher un film |
| `POST` | `/movies` | Ajouter un film |
| `GET` | `/movies/{id}` | Détails d'un film |
| `PUT` | `/movies/{id}` | Modifier un film |
| `DELETE` | `/movies/{id}` | Supprimer un film |

## 👤 Auteur

**Valentin Faust** 
[Mon Portfolio](https://valentinfaustweb.vercel.app/)

---

*Projet personnel réalisé dans le but d'apprendre les bases du langage Go et le fonctionnement d'une API REST. Le projet intègre une base de données PostgreSQL via Docker, une documentation automatique avec Swagger, et un pipeline d'intégration continue (CI) via GitHub Actions.*