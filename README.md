📄 Serviço de Simulação de Empréstimos
Go + Fiber + Cálculo PRICE com big.Float


📌 Visão Geral
Este serviço expõe o endpoint POST /simular, responsável por calcular a parcela fixa (sistema PRICE) de operações de empréstimo.

A implementação utiliza Go com math/big.Float, garantindo:\
Alta performance\
Precisão financeira\
Concorrência leve\
Baixo consumo de CPU e memória\


🧮 1. Cálculo PRICE com big.Float

📝 Função: MonthlyPayment(pv, rate float64, n int) float64

Função de alta precisão para cálculo da parcela do financiamento (PRICE).

📥 Parâmetros de Entrada
| Parâmetro | Tipo      | Descrição                                |
| --------- | --------- | ---------------------------------------- |
| `pv`      | `float64` | Valor financiado (valor presente).       |
| `rate`    | `float64` | Taxa mensal de juros (ex.: `0.02` = 2%). |
| `n`       | `int`     | Quantidade de parcelas.                  |

📤 Retorno
| Tipo      | Descrição                   |
| --------- | --------------------------- |
| `float64` | Valor da parcela calculada. |

🌐 2. API REST (Go + Fiber)
📍 Endpoint
POST /simular

📥 Corpo da Requisição (JSON)

{  
  "amount": 10000,\
  "rate": 0.02,\
  "months": 12\
}

| Campo    | Tipo      | Obrigatório | Descrição               |
| -------- | --------- | ----------- | ----------------------- |
| `amount` | `float64` | ✔           | Valor financiado.       |
| `rate`   | `float64` | ✔           | Taxa mensal de juros.   |
| `months` | `int`     | ✔           | Quantidade de parcelas. |

📤 Corpo da Resposta (JSON)

{\
  "installment": 937.42,\
  "amount": 10000,\
  "rate": 0.02,\
  "months": 12\
}

| Campo         | Tipo      | Descrição                   |
| ------------- | --------- | --------------------------- |
| `installment` | `float64` | Parcela calculada.          |
| `amount`      | `float64` | Valor financiado informado. |
| `rate`        | `float64` | Taxa informada.             |
| `months`      | `int`     | Prazo informado.            |

🧪 Exemplo de Uso — cURL

🔹 Request
curl -X POST http://localhost:8080/simular \
  -H "Content-Type: application/json" \
  -d '{\
        "amount": 10000,\
        "rate": 0.02,\
        "months": 12\
      }'


🔹 Response\
{\
  "installment": 937.42,\
  "amount": 10000,\
  "rate": 0.02,\
  "months": 12\
}


🎯 Conclusão

O serviço oferece:\
🔥 Alta performance com big.Float\
⚡ Baixa latência, ideal para grandes volumes\
🧮 Precisão financeira real\
🧩 API limpa e simples de integrar (Fiber)\
🏗️ Pode ser facilmente containerizado e escalado
