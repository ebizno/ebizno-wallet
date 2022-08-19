# e-Bizno
## Descrição

Repositório do Apache Kafka (Mensageria)

**obs**: Roda-lo antes de qualquer apliação

## Configurar /etc/hosts
Em todos os sistemas operacionais é necessário abrir o programa para editar o *hosts* como Administrator da máquina ou root.
```
127.0.0.1 host.docker.internal
```

## Rodar a aplicação

Execute os comandos:

```
docker-compose up
```
***Importante***: Toda vez que para-lo de rodar, rode ```docker-compose down``` para destruir os volumes, senão ao rodar um UP novamente dará erro
