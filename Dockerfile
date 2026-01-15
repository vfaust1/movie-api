# 1. Image de base officielle qui contient déjà Go
FROM golang:alpine

# 2. Création d'un dossier de travail à l'intérieur du conteneur
WORKDIR /app

# 3. Copie des fichiers de dépendances (go.mod et go.sum)
COPY go.mod go.sum ./

# 4. Téléchargement des librairies
RUN go mod download

# 5. Copie de tout le reste du code source dans le conteneur
COPY . .

# 6. Compilation de l'application
RUN go build -o main ./cmd/api

# 7. Expose du port 8080
EXPOSE 8080

# 8. Commande lancé au démarrage du conteneur
CMD ["./main"]