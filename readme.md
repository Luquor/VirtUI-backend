# Fonctionnement API LXD
## Opérations

## Websocket


# Dépendances du projet
## Go-Chi
Go-chi est un routeur léger et facile à prendre en main. Nous avons choisi cette dépendances pour plusieurs points :
- **Simplification :** la syntaxe simple nous permets de nous concentrer sur notre application sans trop nous soucier du fonctionnement du routing.
- **Concurrence :** les autres projets comme *Way*, *Gin* et *Fiber* sont plus complets/complexe ; il était alors plus logique pour notre projet de choisir un routeur léger.

Le service Go-chi nous permet donc de faire le routing général de l'application grâce aux méthodes `Get/Post/Put/Delete`.
Voici un exemple de son utilisation :
```golang
r := chi.NewRouter()
r.Get("/", homepage)
r.Post("/container", createContainer)
r.Delete("/container/{name}", deleteContainer)
```

Dans l'exemple ci-dessus on crée un router, et comme avec le paquet `net/http`, on précise son pattern ainsi qu'une fonction utiliser pour cette route.
La fonction `createContainer()` va faire un appel à la fonction éponyme.
Tout les retours des fonctions de handling renvoient du JSON, facilitant ainsi la communication entre le front-end et le back-end. 


# Structure du code
## Code applicatif
Puisque nous avons appris la programmation objets lors de nos cours, nous avons architecturé avec une "classe" (fichier)  pour un "objet" (structure).

### Authentification


## Tests unitaires

