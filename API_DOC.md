# Pour la création d'un certificat pour l'API Rest

```bash
lxc config trust add
# Récupèrer le token et le nom du client
lxc remote add <nom> <token>
```
A partir d'ici un certificat est une clé vont être généré dans ~/lxd/ ou ~/snap/lxd/common/config/
Pour interroger l'api il faut donc spécifié la clé et le cert : 

```
curl -s -k --cert ~/snap/lxd/common/config/client.crt --key ~/snap/lxd/common/config/client.key https://127.0.0.1:8443/1.0/[REQUETE] | jq .
```

Documentation API [ici](https://documentation.ubuntu.com/lxd/en/latest/api/#/)