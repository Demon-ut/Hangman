# Jeu du Pendu en Go

Bienvenue dans le jeu du pendu, développé en Go ! Ce projet est un jeu de type console qui propose une expérience de jeu de pendu, incluant un système de score, des indices colorés, la sauvegarde de l’état de la partie, et des fonctionnalités spéciales.

## Fonctionnalités

- **Jeu de pendu classique** : Le jeu propose un mot aléatoire à deviner, avec un nombre limité de tentatives.
- **Easter Egg** : Tapez `Mario` pendant le jeu pour découvrir un bonus visuel caché !
- **Sauvegarde de la progression** : Tapez `stop` pour sauvegarder votre état de jeu (mot actuel et nombre de tentatives restantes).
- **Indications colorées** : Chaque action (erreur, tentative réussie, mot déjà deviné) est accompagnée de codes couleur pour une meilleure lisibilité.
- **Affichage ASCII Art** : Le mot et le statut de la partie sont affichés en ASCII Art.

## Prérequis

- Go doit être installé sur votre système.
  Voici le site --> [Golang](https://go.dev/)

## Installation

1. **Cloner le dépôt :**

   ```sh
   git clone https://ytrack.learn.ynov.com/git/gtimothe/hangman-classic
   cd hangman
   ```
2. **Lancer le jeu :**
   Vérifier que vous êtes bien dans le dossier du hangman avent d'utiliser cette commande pour cela utiliser :

```sh
cd Votre chemin d'acces
```

Lancer le jeux avec cette commande :

```sh
   go run main.go
```

### Détails des fichiers

- **General.go** : La fonction `General()` est le cœur du jeu. Elle initialise les variables de couleur, le score, et le nombre de tentatives. Elle contrôle également la logique principale du jeu et vérifie les conditions de victoire ou de défaite.
- **StopSave.go** : Cette fonction sauvegarde l'état actuel du jeu (mot à deviner, mot partiellement découvert, et nombre de tentatives restantes) dans `save.txt`.
- **PrintHG.go** : Affiche le menu principal avec des options et un message de bienvenue.
- **ChoisiMot.go** : Sélectionne un mot au hasard depuis `words.txt`. En cas d'erreur (fichier non trouvé), un message s’affiche.
- **RevLettres.go** : Révèle un certain nombre de lettres aléatoirement au début du jeu pour aider le joueur.
- **Pendu.go** : Affiche l’état du pendu selon le nombre de tentatives restantes.
- **ChangeValue.go** : Affiche le mot partiellement découvert en ASCII Art.

## Instructions de jeu

1. **Démarrage** :

   - L’écran d’accueil s’affiche avec le titre du jeu en couleurs.
   - Pour commencer une nouvelle partie, entrez n’importe quelle lettre (ou tapez `q` pour quitter).
2. **Gameplay** :

   - **Devinez une lettre** : Tapez une lettre pour essayer de la deviner. Le jeu vérifie si elle est présente dans le mot et l’affiche si c’est le cas.
   - **Commandes spéciales** :
     - `stop` : Sauvegarde l’état actuel du jeu dans `save.txt` et termine la session.
     - `Mario` : Affiche un easter egg spécial dans la console.
3. **Indicateurs de statut** :

   - **Couleurs** :
     - Vert : Indique un bon choix de lettre.
     - Rouge : Indique un mauvais choix ou une erreur.
     - Blanc : Indique que la lettre a déjà été devinée.
   - **ASCII Art** : Le mot et les lettres devinées sont affichés en ASCII pour une expérience visuelle plus immersive.
4. **Conditions de victoire et défaite** :

   - **Victoire** : Si toutes les lettres du mot sont devinées avant l’épuisement des tentatives, le jeu affiche un message de victoire et augmente le score.
   - **Défaite** : Si le nombre de tentatives atteint zéro, le jeu affiche le mot complet et réinitialise le score.

## Contributeurs

- **Anthony Goasdoué** - Développeurs principal et auteur du code.
- **Timothé Guinard** - Développeurs principal et auteur du code.

## Licence

Le projet est sous licence MIT - vous êtes libre de l’utiliser, de le modifier et de le partager selon les termes de cette licence.
