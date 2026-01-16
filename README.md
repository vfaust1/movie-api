# 🎬 Movie API

Une API REST robuste pour gérer une collection de films, construite avec **Go**, **PostgreSQL** et **Docker**.
Ce projet démontre une architecture backend moderne, sécurisée et prête pour la production.

![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go) ![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=flat&logo=docker) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-336791?style=flat&logo=postgresql)

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
.
├── cmd/api/                 # Le point d'entrée de l'application
│   ├── main.go              # Initialise la BDD, charge la config et lance le serveur
│   ├── routes.go            # Définit les URLs et applique les middlewares
│   ├── handlers.go          # Contrôleurs : reçoit la requête HTTP -> appelle la BDD -> répond en JSON
│   └── middleware.go        # Intercepteurs : gère l'Authentification (API Key) et les Logs
│
├── internal/store/          # La logique métier et l'accès aux données (interne au projet)
│   ├── db.go                # Configuration BDD : Connexion Postgres, Retry Logic (Docker) et création des tables
│   ├── movies.go            # Modèle : Structures Go et requêtes SQL (CRUD, filtres, pagination)
│   └── movies_test.go       # Tests unitaires pour valider le code
│
├── Dockerfile               # Recette de construction de l'image (Multi-stage build Alpine)
├── compose.yaml             # Orchestrateur : lance l'API et PostgreSQL ensemble
├── .env.example             # Exemple de configuration (copier en .env pour usage hors Docker)
├── go.mod & go.sum          # Gestion des dépendances du module Go
└── .gitignore               # Fichiers à exclure de Git (ex: .env, binaires)
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
*Projet réalisé dans le cadre d'un parcours de formation Backend Go.*