# DadMidia 🎬

Microserviço RESTful desenvolvido em **Go** para a disciplina de **Desenvolvimento de Aplicações Distribuídas (DAD)**. A API gerencia mídias com autenticação JWT, suporte a CORS e integração com banco de dados relacional (MySQL).

---

## Tecnologias

- **Go** 1.25+
- **Gin** — framework HTTP
- **MySQL** — banco de dados relacional
- **CORS** — habilitado para todas as origens

---

## Estrutura do Projeto

```
DadMidia/
├── db/             # Inicialização e configuração dos bancos de dados
├── middlewares/    # Middlewares (ex: autenticação JWT)
├── models/         # Structs e modelos de dados
├── routes/         # Definição das rotas da API
├── utils/          # Funções auxiliares
├── main.go         # Ponto de entrada da aplicação
├── go.mod          # Módulo Go e dependências
└── go.sum          # Checksums das dependências
```

---

## Pré-requisitos

Antes de iniciar, certifique-se de ter instalado:

- [Go](https://go.dev/doc/install) 1.25 ou superior
- [MySQL](https://dev.mysql.com/downloads/) (ou um servidor MySQL acessível)
- [Git](https://git-scm.com/)

---

## Instalação e Execução

### 1. Clonar o repositório

```bash
git clone https://github.com/ODanielFernandes/DadMidia.git
cd DadMidia
```

### 2. Configurar variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto com as configurações dos bancos de dados e do JWT:

```env
# MySQL
DB_HOST=localhost
DB_PORT=3306
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=dad_midia

```

> O projeto utiliza o pacote `godotenv` para carregar automaticamente o arquivo `.env`.

### 3. Instalar as dependências

```bash
go mod tidy
```

### 4. Executar a aplicação

```bash
go run main.go
```

O servidor será iniciado em `http://localhost:8080`.

---

## Endpoints

A API estará disponível em `http://localhost:8080`. As rotas são registradas pelo pacote `routes` — consulte os arquivos em `routes/` para ver a lista completa de endpoints disponíveis.

---

## Compilando para produção

Para gerar um binário executável:

```bash
go build -o dad_midia .
./dad_midia
```

---

## Licença

Projeto desenvolvido para fins acadêmicos na disciplina de DAD.