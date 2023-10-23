# Chat server TCP em Golang

Primeiramente em um terminal devemos executar o nosso programa para isso utilize o esses dois comandos:
- `go build .` 
- `./GoChat`

Em seguida em outro terminal devemos conectar ao server para isso deve-se utilizar o seguinte comando:
- `telnet localhost 8888` - o numero da porta retorna ao executar o nosso programa
# Comandos

- `/nick <seu nome>` - escolher o seu nick
- `/change <mudar seu nome>` - mudar o nick
- `/msg <destinatario, mensagem>` - enviar uma mensagem
- `/quit` - desconecta do server