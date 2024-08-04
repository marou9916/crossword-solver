# Crossword Solver (voir la version française de ce document en dessous de la version anglaise après la section "Licence")
(See the French version of this document below the English version after the "License" section.)

This project is a renewed attempt to solve the Crossword Solver problem, which I had tackled a few months ago. At that time, I didn't have a good grasp of backtracking, which posed some difficulties. Determined to overcome this challenge, I studied and practiced backtracking through various exercises on GeeksForGeeks. You can find these solutions in my repository `problem-solving-practice`, under the 'backtracking' folder. With this experience, I now feel ready to apply my new skills to solve the Crossword Solver.

## Introduction

The Crossword Solver is a crossword puzzle solver written in Go. It takes a puzzle grid and a set of words as input and finds the positions where the words can be placed in the grid.

## Learning Highlights

Working on the Crossword Solver project has provided valuable insights and hands-on experience in several key areas of software development:

**Unit Testing in Go**: I learned how to write comprehensive unit tests for various functions, ensuring that each component of the system behaves as expected. This practice not only helps in maintaining code quality but also aids in early detection of bugs.

**Iterative Problem Decomposition**: Tackling the Crossword Solver problem required breaking down the problem into smaller, manageable tasks. This iterative decomposition approach helped in systematically addressing each aspect of the problem, leading to a robust solution.

**Logical Grouping of Functions**: I organized functions into logical units using packages. This modular approach enhances code readability, maintainability, and reusability. Each package, such as placement, solver, remover, and validator, serves a distinct purpose, contributing to a well-structured codebase.

**Naming Conventions**: Giving meaningful names to functions, variables, and files is crucial for code clarity. Throughout the project, I adhered to naming conventions that accurately describe the purpose and functionality of each component, making the code easier to understand and navigate.

**Backtracking Algorithm**: Implementing the backtracking algorithm for solving the crossword puzzle was a significant learning experience. I gained a deeper understanding of this algorithmic technique and its application in solving complex problems.

**Practical Go Experience**: This project provided extensive hands-on experience with the Go programming language, including its syntax, standard libraries, and best practices for writing idiomatic Go code.

These learning highlights not only contributed to the successful completion of the Crossword Solver project but also enriched my overall skill set as a software developer.

## Project Structure
```
crossword-solver/
├── cmd/
│   └── crosswordsolver/
│       └── main.go
├── internal/
│   ├── model/
│   │   └── position.go
│   ├── placement/
│   │   ├── can_place_word_horizontally.go
│   │   ├── can_place_word_horizontally_test.go
│   │   ├── place_word_horizontally.go
│   │   ├── place_word_horizontally_test.go
│   │   ├── can_place_word_vertically.go
│   │   ├── can_place_word_vertically_test.go
│   │   └── place_word_vertically.go
│   │   ├── place_word_vertically_test.go
│   ├── remover/
│   │   ├── remove_word_horizontally.go
│   │   ├── remove_word_horizontally_test.go
│   │   ├── remove_word_vertically.go
│   │   └── remove_word_vertically_test.go
│   ├── solver/
│   │   ├── try_to_find_solution_with_backtracking.go
│   │   ├── try_to_find_solution_with_backtracking_test.go
│   │   ├── convert_puzzle_to_grid.go
│   │   ├── convert_puzzle_to_grid_test.go
│   │   ├── fill_grid_with_content.go
│   │   ├── fill_grid_with_content_test.go
│   │   ├── find_positions.go
│   │   ├── find_positions_test.go
│   │   ├── get_dot_positions.go
│   │   ├── get_dot_positions_test.go
│   │   └── solve.go
│   ├── printer/
│   │   └── print_grid.go
│   └── validator/
│       ├── validate_string.go
│       ├── validate_string_test.go
│       ├── validate_setofwords.go
│       └── validate_setofwords_test.go
├── setup_crossWordSolver.sh
├── go.mod
└── go.sum
```

## Features

**Puzzle Format Validation**: Checks that the provided puzzle grid is in a valid format.

**Word Set Validation**: Ensures the provided set of words contains no duplicates.

**Grid Conversion**: Converts the puzzle grid from a string format to a 2D rune slice.

**Dot Position Identification**: Identifies the positions in the grid where words can potentially be placed.

**Word Placement Checking**: Checks if a word can be placed at a specific position horizontally or vertically.

**Word Placement**: Places words in the grid if they can be placed without conflicts.

**Word Removal**: Removes words from the grid, useful for backtracking.

**Backtracking Solver**: Uses backtracking to try to place all words in the grid.

**Grid Printing**: Prints the current state of the grid.

## Usage

1. **Clone the repository**:

    ```bash
    git clone https://github.com/marou9916/crossword-solver.git
    cd crossword-solver
    ```

2. **Install dependencies**:

    ```bash
    go mod tidy
    ```

3. **Run the tests**:

    ```bash
    go test ./...
    ```
**Ensure there is no space after the last slash**

## Contributions

Contributions are welcome! Please submit a pull request or open an issue to discuss any changes you would like to make.

## Licence

This project is licensed under the MIT License. See the LICENSE file for details.




# Crossword Solver

Ce projet est une nouvelle tentative de résoudre le problème du solveur de mots croisés, que j'avais abordé il y a quelques mois. À l'époque, je n'avais pas une bonne compréhension du backtracking, ce qui avait posé certaines difficultés. Déterminé à surmonter ce défi, j'ai étudié et pratiqué le backtracking à travers divers exercices sur GeeksForGeeks. Vous pouvez trouver ces solutions dans mon dépôt `problem-solving-practice`, sous le dossier 'backtracking'. Avec cette expérience, je me sens maintenant prêt à appliquer mes nouvelles compétences pour résoudre le solveur de mots croisés.

## Introduction

Le Crossword Solver est un solveur de mots croisés écrit en Go. Il prend une grille de puzzle et un ensemble de mots comme entrée et trouve les positions où les mots peuvent être placés dans la grille.

## Points d'Apprentissage

Travailler sur le projet Crossword Solver m'a permis d'acquérir des connaissances et une expérience précieuses dans plusieurs domaines clés du développement logiciel :

**Tests Unitaires en Go** : J'ai appris à écrire des tests unitaires complets pour diverses fonctions, garantissant que chaque composant du système fonctionne comme prévu. Cette pratique aide non seulement à maintenir la qualité du code, mais aussi à détecter les bogues plus tôt.

**Décomposition Itérative du Problème** : Aborder le problème du Crossword Solver a nécessité la division du problème en tâches plus petites et gérables. Cette approche itérative de décomposition a permis de traiter systématiquement chaque aspect du problème, conduisant à une solution robuste.

**Regroupement Logique des Fonctions** : J'ai organisé les fonctions en unités logiques en utilisant des packages. Cette approche modulaire améliore la lisibilité, la maintenabilité et la réutilisabilité du code. Chaque package, tel que placement, solver, remover, et validator, a un objectif distinct, contribuant à une base de code bien structurée.

**Conventions de Nommage** : Donner des noms significatifs aux fonctions, variables et fichiers est crucial pour la clarté du code. Tout au long du projet, j'ai respecté des conventions de nommage décrivant précisément le but et la fonctionnalité de chaque composant, rendant le code plus facile à comprendre et à naviguer.

**Algorithme de Backtracking** : L'implémentation de l'algorithme de backtracking pour résoudre le puzzle de mots croisés a été une expérience d'apprentissage significative. J'ai approfondi ma compréhension de cette technique algorithmique et de son application pour résoudre des problèmes complexes.

**Expérience Pratique avec Go** : Ce projet a fourni une expérience pratique étendue avec le langage de programmation Go, y compris sa syntaxe, ses bibliothèques standard, et les meilleures pratiques pour écrire un code Go idiomatique.

Ces points d'apprentissage ont non seulement contribué à l'achèvement réussi du projet Crossword Solver, mais ont également enrichi mon ensemble de compétences global en tant que développeur logiciel.


## Structure du projet
```
crossword-solver/
├── cmd/
│   └── crosswordsolver/
│       └── main.go
├── internal/
│   ├── model/
│   │   └── position.go
│   ├── placement/
│   │   ├── can_place_word_horizontally.go
│   │   ├── can_place_word_horizontally_test.go
│   │   ├── place_word_horizontally.go
│   │   ├── place_word_horizontally_test.go
│   │   ├── can_place_word_vertically.go
│   │   ├── can_place_word_vertically_test.go
│   │   └── place_word_vertically.go
│   │   ├── place_word_vertically_test.go
│   ├── remover/
│   │   ├── remove_word_horizontally.go
│   │   ├── remove_word_horizontally_test.go
│   │   ├── remove_word_vertically.go
│   │   └── remove_word_vertically_test.go
│   ├── solver/
│   │   ├── try_to_find_solution_with_backtracking.go
│   │   ├── try_to_find_solution_with_backtracking_test.go
│   │   ├── convert_puzzle_to_grid.go
│   │   ├── convert_puzzle_to_grid_test.go
│   │   ├── fill_grid_with_content.go
│   │   ├── fill_grid_with_content_test.go
│   │   ├── find_positions.go
│   │   ├── find_positions_test.go
│   │   ├── get_dot_positions.go
│   │   ├── get_dot_positions_test.go
│   │   └── solve.go
│   ├── printer/
│   │   └── print_grid.go
│   └── validator/
│       ├── validate_string.go
│       ├── validate_string_test.go
│       ├── validate_setofwords.go
│       └── validate_setofwords_test.go
├── setup_crossWordSolver.sh
├── go.mod
└── go.sum
```

## Fonctionnalités

**Validation du Format du Puzzle** : Vérifie que la grille de puzzle fournie est dans un format valide.

**Validation de l'Ensemble de Mots** : Assure que l'ensemble de mots fourni ne contient pas de doublons.

**Conversion de la Grille** : Convertit la grille du puzzle du format chaîne de caractères en une tranche 2D de runes.

**Identification des Positions des Points** : Identifie les positions dans la grille où les mots peuvent potentiellement être placés.

**Vérification de Placement des Mots** : Vérifie si un mot peut être placé à une position spécifique horizontalement ou verticalement.

**Placement des Mots** : Place les mots dans la grille s'ils peuvent être placés sans conflits.

**Suppression des Mots** : Supprime les mots de la grille, utile pour le backtracking.

**Solveur avec Backtracking** : Utilise le backtracking pour essayer de placer tous les mots dans la grille.

**Impression de la Grille** : Imprime l'état actuel de la grille.

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
**Assurez-vous qu'il n'y a pas d'espace après la dernière barre oblique**

## Contributions

Les contributions sont les bienvenues ! Veuillez soumettre une demande de tirage ou ouvrir un problème pour discuter de tout changement que vous souhaitez apporter.

## Licence

Ce projet est sous licence MIT. Voir le fichier LICENSE pour plus de détails.