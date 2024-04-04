# Fonctionnement API LXD

L'API RestFULL de LXD permet d'intéragir depuis différentes route HTTP pour permettre une gestion des containeurs, des clusters, des snapshots, ou encore du terminal de commande.

Pour éviter d'exposer toutes l'API de LXD et permettre une gestion des utilisateurs plus souple, nous avons décidé d'ajouter notre propre serveur WEB entre l'utilisateur et l'API.

![Structure basique](./documents/basic_structure.svg)


## Opérations

LXD permet une gestion asynchrone des tâches. En effet certaines tâches peuvent mettre plus ou moins longtemps a se faire (Allumage d'une instance, Création d'un cluster, Restauration d'une snapshot, ...).
Pour éviter d'attendre LXD retourne directement une ``Operation``, qui va permettre avec certaines routes de pouvoir suivre l'avancement de la tâche.

Il existe plusieurs manière de récupérer l'état de l'opération :

- ``/1.0/operations/{ID}`` - Retourne un JSON avec différentes informations (status, description, ...)
- ``/1.0/operations/{ID}/websocket`` - Permet l'ouverture d'un websocket qui va recevoir des messages en fonction de l'avancement
- ``/1.0/operations/{ID}/wait`` - Va attendre la fin (ou le faille) de l'opération pour renvoyer une réponse.



## Terminal

LXD fourni la possibilité de mettre en place un terminal, pour cela il se repose sur le protocole WebSocket.

Le Websocket ouvre un canal de communication bi-directionnel entre le serveur web et le client. Cela à plusieurs avantages comme :

- Pouvoir notifier le client d'une modification
- L'envoie de donnée du serveur vers le client sans requête du client.

Cela permet de pousser l'interactivité au maximum.

LXD permet l'ouverture d'un websocket après l'éxécution d'une commande (ici ``bash``),
après l'ouverture du socket LXD retourne les flux de sortie (stdout, stderr) et récupère le flux d'entrée (stdin)

### Websocket

Pour que le websocket fonctionne, le client doit parler directement a l'API LXD sans passer par notre serveur.
Pour cela nous avons proxifié une partie de l'api LXD pour la faire passer dans notre back :

![Structure avance](./documents/advances_structure.svg)


# Dépendances du projet
## Go-Chi
Go-chi est un routeur léger et facile à prendre en main. Nous avons choisi cette dépendances pour plusieurs points :
- **Simplification :** la syntaxe simple nous permets de nous concentrer sur notre application sans trop nous soucier du fonctionnement du routing.
- **Concurrence :** les autres projets comme *Way*, *Gin* et *Fiber* sont plus complets/complexe ; il était alors plus logique pour notre projet de choisir un routeur léger.

Le service Go-chi nous permet donc de faire le routing général de l'application grâce aux méthodes `GET|POST|PUT|DELETE`.
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
### Structure du code
Puisque nous avons appris la programmation objets lors de nos cours, nous avons architecturé le projet avec une "classe" (fichier)  pour un "objet" (structure).

> **_NOTE :_** cela implique donc une répétition de certains morceau de code.

Le projet est architecturé comme cela :
```bash
.
├── api
├── go.mod
├── go.sum
├── index.html
├── main.go
├── main_test.go
├── models
├── tls
```
Les répertoires sont plutôt éponymes.
Dans `api/` on va y retrouver tous les fichiers concernant la communcation avec l'API de LXD :
- **StandardReturn() :** retour présent sur toutes les structures (conteneurs, clusters, snapshots...);
- **client.go :** créer un client (génération de clefs TLS) pour communiquer avec l'API LXD;

Dans le répertoire `models/` nous avons définis toutes les structures ainsi que les méthodes utilisées pour piloter l'application :
- **server.go :** exécution du serveur web (go-chi) et routing des pages;
- **containers.go :** permet de contrôler les conteneurs;
- **clusters.go :** permet de contrôler les clusters
- **images.go :** contrôle des images LXC;
- **console.go :** récupère les informations pour ouvrir un websocket (token, url de contrôle/envoi...); 

Dans le fichier `main.go` il y a juste le lancement du serveur web.

### Authentification
Pour la partie authentification nous avons créé un middleware (le fichier correspondant est `middleware2.go` ) qui selon certaines conditions va renvoyer un status code http Unauthorized.
- **Si l'en-tête Authorization est vide et que l'URI de la requête n'est pas dans "/auth" ou "/", renvoie une réponse 401.**
- **Sinon, si l'en-tête Authorization n'est pas vide et que l'URI de la requête n'est pas dans "/auth" ou "/", vérifie le token d'authentification dans le cache du serveur** 
- **Si le token d'authentification est invalide, renvoie une réponse 401.**
Pour la connexion au serveur on a créé une fonction nommé 'authenticate' permettant de faire une verification des noms d'utilisateurs et des mots de passe si ces derniers correspondent à un binome <utilisateur:mot-de-passe> du fichier user.txt alors un token est généré et envoyé au client dans l'en-tête,et ce token est rajouté dans le cache du serveur .
Cette fonction est appelé lors du requête sur la route "/auth".
### Cache
Pour la gestion du cache nous avons créé un fichier `cache.go` qui où sont écrits les fonctions permettant de gérer le cache du serveur .
Nous avons un constructeur qui sera appelé par le server lors de son lancement afin d'initialiser le cache.


### Potentielles améliorations
Nous pensons qu'il est possible d'améliorer la code base sur plusieurs point :
- **Structure :** comme dit précedemment, nous avons réalisé le projet en suivant une structure "programmation objet". Il y a donc des répétitions dans le code, surtout entre `container.go` et `clusters.go`. Avec plus de temps, on aurait fait du refactoring ; par exemple définir une seule méthode `Create()` qui, en fonction de la structure/json passé en paramètre, créer un objet correspondant. Cela aurait permis de limiter la répétition du code. La même remarque peut-être faite pour les getters par id/nom ou getter local.
- **Scripting :** pour la génération de certificats TLS, on aurait pus réalisé un script permettant de les générer.

## Tests 
Pour les test on a opté pour des tests fonctionnels, nous avons utilisé la librairie git `github.com/stretchr/testify/assert` permettant d'utiliser des fonctions tel que assert.Equals .
Nous avons jugé que des tests unitaires n'étaient pas intérressant ,globalement les tests se déroulent en 3 étapes.

1- On lance la fonction qui sera appelé lors du clic du bouton 
2- On vérifie à l'aide d'une commande shell et on stock le résultat de la commande
3- On vérifie le résultat attendu 

Les test sont faits de tel sorte qu'au fur et à mesure on puisse utilisé les fonctions déjà testé pour la suite (Ce n'est pas la meilleur des méthodes...)

Nous avons mis en place un CI à l'aide des github actions cependant nous avons rencontré beaucoup de dificulté :
notemment pour le service lxd .A la fin on a du abandonné car on a supposé que c'était probablement impossible (sur la fin du debug on avait une sombre histoire de socket)
```Error: The LXD daemon doesn't appear to be started (socket path: /var/snap/lxd/common/lxd/unix.socket)```
Ceci fut notre dernière erreur et sommes resté bloquer dessus car le service lxd était lancé et que c'était probablement une histoire d'autorisation pas très clair...
On a opté pour un self-hosted runner dans lequel on installé tout ce qu'il fallait sur googlecloud.
On a créé un service pour que le run.sh .
**NOTE :** On aurait pu faire sur une machine en cours mais on ne peut pas travailler depuis chez soi (si jamais la machine est débranché pour x ou y raison)

# Premier itération du front-end (SSR)
Au début du projet, nous avions décidé de faire le front-end avec juste du HTML/CSS/JS (ou HTMX) tout cela géré par le serveur Golang (Server Side Rendering).  
Pour éviter de faire du copier/coller de code, nous voulions utiliser les templates de Golang ce qui permet de réaliser des "composants" réutilisables.
Nous n'avons pas réussi à mettre en place ce système car nous rencontrions des problèmes pour intégrer les templates au serveur web.
Un moyen de contourner ce problème aurait été d'ajouter une dépendance comme `tylermmorton/tmpl` qui permet de gérer les templates HTML en plus de simplifier le workflow.  
Nous avons donc décidé de faire le front-end en utilisant VueJS.