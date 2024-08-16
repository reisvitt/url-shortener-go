<h2 align="center">
  Desafio Backend: Encurtador de URL
</h2>

Resolução do desafio proposto pelo repositorio Backend Brasil, confira detalhes [neste link](https://github.com/backend-br/desafios/blob/master/url-shortener/PROBLEM.md).

## :rocket: Tecnologias utilizadas

- Golang 1.22.6
- Gin Gonic
- Docker
- MongoDB
- Mongo Driver

## Instalação

:mag: Baixe o projeto e teste você.

1. Clone o repositório:

   ```bash
   git clone https://github.com/reisvitt/url-shortener-go.git
   ```

2. Navegue até a pasta da aplicação:

   ```bash
   cd url-shortener-go
   ```

3. Faça uma cópia do .env.example e insirá suas variáveis de ambiente.

4. Execute o comando docker compose up (Opcional)

   _Certifique-se de ter o Docker instalado em sua máquina_

   ```bash
     docker compose up -d
   ```

   _Caso for utilizar um banco MongoDB já existente, insirá o DB_CONNECTION no arquivo .env_

5. Inicie o servidor

   ```bash
     go run main.go
   ```

## Uso

### 1. Encurtando uma URL

Envie uma requisição POST para o endpoint /shorten-url, incluindo um corpo no formato JSON com o seguinte campo:

- **url**: A URL que você deseja encurta

**Exemplo de Requisição via cURL:**

```bash
curl -X POST -H "Content-Type: application/json" -d '{"url":"https://example.com"}' http://localhost:8080/shorten-url
```

_Alternativamente, você pode usar um cliente HTTP de sua preferência, como Postman, Insomnia, ou qualquer outro._

### 2. Recebendo a URL Encurtada

O servidor retornará uma URL encurtada no formato JSON. Você pode utilizar essa URL encurtada diretamente em seu navegador para redirecionar para a URL original.

**Exemplo de Resposta:**

```json
{
  "shortUrl": "http://localhost:8080/abc123"
}
```

### 3. Usando a URL Encurtada

Copie a URL encurtada retornada e insira-a em seu navegador. Você será automaticamente redirecionado para a URL original.

## Licença

Este projeto está licenciado sob os termos da licença MIT. Veja o arquivo `LICENSE` para mais detalhes.
