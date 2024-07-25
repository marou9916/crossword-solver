# Crossword Solver

Ce projet est un solveur de mots croisés en Go. Il prend une grille de puzzle et un ensemble de mots comme entrée et trouve les positions où les mots peuvent être placés dans la grille.

## Structure du Projet

crossword-solver/
├── cmd/
│ └── crosswordsolver/
│ └── main.go
├── internal/
│ ├── validator/
│ │ ├── validate_string.go
│ │ ├── validate_string_test.go
│ │ ├── validate_setofwords.go
│ │ └── validate_setofwords_test.go
│ ├── solver/
│ │ ├── backtrack.go
│ │ ├── place.go
│ │ ├── positions.go
│ │ └── solve.go
│ └── model/
│ └── position.go
├── test/
│ └── solver_test.go
├── go.mod
└── go.sum


## Fonctionnalités

- **Validation du format du puzzle**: Vérifie que la grille de puzzle fournie est dans un format valide.
- **Validation de l'ensemble de mots**: Vérifie que l'ensemble de mots fourni ne contient pas de doublons.
- **Placement des mots**: Détermine les positions possibles pour placer les mots dans la grille.

## Utilisation

1. **Cloner le dépôt**:

    ```bash
    git clone https://github.com/marou9916/crossword-solver.git
    cd crossword-solver
    ```

2. **Installer les dépendances**:

    ```bash
    go mod tidy
    ```

3. **Exécuter les tests**:

    ```bash
    go test ./...
    ```

## Contributions

Les contributions sont les bienvenues ! Veuillez soumettre une pull request ou ouvrir une issue pour discuter des modifications que vous souhaitez apporter.

## Licence

Ce projet est sous licence MIT. Voir le fichier [LICENSE](LICENSE) pour plus de détails.



