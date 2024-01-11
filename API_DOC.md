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

__Commande de test avec le certificat__
```bash
curl -s -k --cert ~/snap/lxd/common/config/client.crt --key ~/snap/lxd/common/config/client.key -X POST
```

## Documentation temporaire  

## Récupèrer les images

https://127.0.0.1:8443/1.0/instances

## Création d'une instance

fingerprint correspond à l'id d'une image

```bash
curl -s -k --cert ~/snap/lxd/common/config/client.crt --key ~/snap/lxd/common/config/client.key -X POST https://127.0.0.1:8443/1.0/instances -H "Content-Type: application/json" -d '{"name":"test56noah","source":{"type":"image","fingerprint":"1722a71a9f2dc0c68eac142a7d53ec728c15d2379e99f5b5545de99d440e3422"}}'| jq .metadata.id | sed 's/"//g' | curl -s -k --cert ~/snap/lxd/common/config/client.crt --key ~/snap/lxd/common/config/client.key https://127.0.0.1:8443/1.0/operations/$(</dev/stdin) | jq .
```
```bash
curl -s -k --cert ~/snap/lxd/common/config/client.crt --key ~/snap/lxd/common/config/client.key -X POST https://127.0.0.1:8443/1.0/instances -H "Content-Type: application/json" -d '{"name":"test56noah","source":{"type":"image","fingerprint":"712a58368655"}}'| jq .metadata.id | sed 's/"//g' | 
curl --include \
--no-buffer \
--header "Connection: Upgrade" \
--header "Upgrade: websocket" \
--header "Host: example.com:80" \
--header "Sec-WebSocket-Version: 13" \
-s -k --cert ~/snap/lxd/common/config/client.crt --key ~/snap/lxd/common/config/client.key https://127.0.0.1:8443/1.0/operations/$(</dev/stdin)
```
