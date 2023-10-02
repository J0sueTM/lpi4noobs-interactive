# LPI 4 Noobs Interactive

Este projeto visa disponibilizar uma aplicação voltada para o aprendizado para a certificação LPI. Baseade no projeto [LPI4Noobs](https://github.com/lanjoni/lpi4noobs), o Interactive conteineriza e automatiza todo o processo, entregando uma experiência mais intuitiva e beginner-friendly aos estudantes.

## Como usar

Antes de tudo, é necessário que a linguagem [Go](https://go.dev/dl/) e o [Docker](https://www.docker.com/) estejam instalados no seu computador.

Depois de clonar, gere o bundle do CSS utilizando o tailwind. Você pode adquiri-lo tanto como binário do seu gerenciador de pacotes ou pelo npm:

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
