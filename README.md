# Customer Schedule API 📅

Esta é uma API REST desenvolvida em **Go** para o gerenciamento de agendamentos de serviços de estética (especializada em extensões de cílios). O projeto utiliza **MongoDB** para persistência de dados e segue boas práticas de validação e estruturação.

## 🚀 Tecnologias Utilizadas

*   **Linguagem:** [Go (Golang)](https://go.dev)
*   **Roteamento:** [Gorilla Mux](https://github.com)
*   **Banco de Dados:** [MongoDB](https://www.mongodb.com) (Driver oficial)
*   **Manipulação de JSON:** Standard Encoding/JSON

## 🛠️ Instalação e Execução

1.  **Clone o repositório:**
    ```bash
    git clone https://github.com
    cd customerscheduleapi
    ```

2.  **Certifique-se de que o MongoDB está rodando:**
    A API espera uma instância local no endereço: `mongodb://localhost:27017`

3.  **Instale as dependências:**
    ```bash
    go mod tidy
    ```

4.  **Execute a aplicação:**
    ```bash
    go run main.go
    ```
    A API iniciará na porta **9437**.

## 🛣️ Endpoints da API

| Método | Rota | Descrição |
| :--- | :--- | :--- |
| `POST` | `/appointments` | Cria um novo agendamento |
| `GET` | `/appointments` | Lista todos os agendamentos |
| `PUT` | `/appointments/{id}` | Atualiza um agendamento existente |
| `DELETE` | `/appointments/{id}` | Remove um agendamento |

### Estrutura do JSON (Exemplo)
```json
{
  "name": "Nome da Cliente",
  "phonenumber": 11988887777,
  "typeservice": 1,
  "date": "2026-10-25T14:00:00Z"
}
