# IAM — Identity and Access Management

**Incidência:** 🔴 Muito Alta | **Domínio:** Segurança

---

## Conceito

IAM é o serviço que controla **quem pode fazer o quê** na sua conta AWS.

Analogia: pense no IAM como um sistema de crachás em uma empresa. Cada pessoa tem um crachá com permissões específicas — só entra nos andares que tem acesso.

---

## Componentes Principais

### Usuários (Users)
- Representam uma pessoa ou aplicação
- Têm credenciais permanentes (senha ou access keys)
- Evite usar o usuário root para tarefas diárias

### Grupos (Groups)
- Coleção de usuários
- Herdam políticas do grupo
- Um usuário pode estar em múltiplos grupos

### Políticas (Policies)
- Documento JSON que define permissões
- Podem ser: AWS Managed, Customer Managed, Inline
- Estrutura: Effect (Allow/Deny) + Action + Resource

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::meu-bucket/*"
    }
  ]
}
```

### Roles (Funções)
- Identidade temporária, sem senha ou access key permanente
- Usada por serviços AWS (EC2, Lambda, ECS) ou usuários externos
- Assume role via STS (Security Token Service)

---

## Regras de Ouro

| Situação | Solução |
|----------|---------|
| Aplicação rodando em EC2 precisa acessar S3 | IAM Role na instância |
| Usuário de outra conta precisa de acesso | Cross-account Role |
| Lambda precisa escrever no DynamoDB | IAM Role na Lambda |
| Acesso temporário para terceiros | Role com condições de tempo |
| Usuário humano precisa de acesso | IAM User no grupo correto |

---

## Princípio do Menor Privilégio

Sempre conceda **apenas as permissões mínimas necessárias**.

Nunca use `"Action": "*"` em produção.

---

## Pegadinhas da Prova

1. **Deny sempre sobrepõe Allow** — se houver Deny explícito, ele vence
2. **Root account** não pode ser restringida por políticas IAM
3. **Access Keys** nunca devem ser commitadas no código — use Roles
4. **IAM é global** — não é por região
5. **MFA é obrigatório** para proteger contas privilegiadas
6. **Policy Evaluation**: Deny explícito > Allow explícito > Deny implícito (padrão)

---

## Casos Reais

### Nubank
O Nubank usa IAM Roles para separar permissões entre times. Cada squad tem uma Role com acesso apenas aos recursos do seu serviço, impedindo acesso cruzado acidental.

### Mercado Livre
Usa IAM com condições (`Condition`) para restringir acesso por IP e horário, garantindo que sistemas internos só acessem a AWS a partir da rede corporativa.

---

## Comparações Importantes

| Característica | IAM User | IAM Role |
|---------------|----------|----------|
| Credenciais | Permanentes | Temporárias (STS) |
| Uso | Pessoas / Aplicações externas | Serviços AWS / Cross-account |
| MFA | Suportado | Via STS |
| Recomendação para EC2/Lambda | ❌ Evitar | ✅ Preferido |

---

## Lab Relacionado

[Lab 01 — IAM: Usuários, Grupos, Políticas](../labs/lab-01-iam-usuarios-grupos.md)

---

## Links de Estudo

- AWS Docs: Identity and Access Management
- AWS Skill Builder: "AWS Security Fundamentals"
- Udemy (Stephane Maarek): Seção IAM
