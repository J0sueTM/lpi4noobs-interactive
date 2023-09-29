# LPI 4 Noobs Interactive

Este projeto visa disponibilizar uma aplicação voltada para o aprendizado para a certificação LPI, baseado no projeto LPI4Noobs. Utilizando do LPI4Noobs, o Interactive conteineriza e automatiza o todo o processo necessário, fazendo com que o estudante não dependa de muito além da própria aplicação.

## Como usar

Antes de tudo, é necessário que a linguagem [Go](https://go.dev/dl/) e o [Docker](https://www.docker.com/) estejam instalados no seu computador. Depois de clonar, é necessário que gere o bundle do CSS utilizando o tailwind. Você pode adquiri-lo tanto como binário do seu gerenciador de pacotes ou pelo npm:

```sh
npx tailwindcss -i views/style/index.css -o views/style/bundle.css
```

Finalmente, você pode tanto compilar quanto rodar diretamente.

```sh
# compilando
go build
./lpi4noobs-interactive

# ou interpretando
go run .
```

## Licença

Copyright © 2023 Josué Teodoro Moreira et. al <j0suetm.com> <teodoro.josue@pm.me>
