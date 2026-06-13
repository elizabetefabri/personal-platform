# Lab 01 — IAM: Usuários, Grupos e Políticas

**Dia:** 6 (Semana 1) | **Duração:** 45 minutos

---

## Objetivo

Criar usuários, grupos e políticas IAM para simular o acesso de uma equipe de desenvolvimento.

---

## Pré-requisitos

- Acesso à conta AWS (AWS Educate ou conta pessoal)
- Usuário com permissão AdministratorAccess

---

## Cenário

Você é a engenheira responsável por configurar acesso para 3 times:
- **Desenvolvedores**: acesso somente leitura ao S3 e CloudWatch
- **DevOps**: acesso de leitura/escrita ao EC2 e S3
- **Leitura Geral**: acesso somente leitura a todos os recursos

---

## Passo a Passo

### 1. Criar Políticas Customizadas

**Política: DeveloperPolicy**
1. IAM → Policies → Create policy
2. Escolha o editor JSON:
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:GetObject",
        "s3:ListBucket"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "cloudwatch:GetMetricData",
        "cloudwatch:ListMetrics",
        "logs:DescribeLogGroups",
        "logs:GetLogEvents"
      ],
      "Resource": "*"
    }
  ]
}
```
3. Nome: `StudyLab-DeveloperPolicy`

---

### 2. Criar Grupos

1. IAM → User groups → Create group
2. **Grupo 1:** `developers` → Attach: `StudyLab-DeveloperPolicy`
3. **Grupo 2:** `devops` → Attach: `AmazonEC2FullAccess` + `AmazonS3FullAccess`
4. **Grupo 3:** `readonly` → Attach: `ReadOnlyAccess`

---

### 3. Criar Usuários

1. IAM → Users → Add users
2. **Usuário 1:** `dev-elizabete`
   - Console access: Sim
   - Grupo: `developers`
3. **Usuário 2:** `devops-deploy`
   - Console access: Sim
   - Grupo: `devops`

---

### 4. Habilitar MFA (Simulação)

1. IAM → Users → `dev-elizabete`
2. Security credentials → MFA device
3. Observe que MFA é uma camada crítica de segurança

---

### 5. Criar IAM Role

1. IAM → Roles → Create role
2. Trusted entity: EC2
3. Attach: `AmazonS3ReadOnlyAccess`
4. Nome: `StudyLab-EC2-S3ReadRole`
5. (Esta role seria usada por instâncias EC2 para acessar S3 sem credenciais)

---

## Resultado Esperado

- 3 grupos criados com políticas corretas
- 2 usuários criados e associados a grupos
- 1 IAM Role para EC2

---

## Validação

1. Login com `dev-elizabete` → tente acessar EC2 → deve ser negado
2. Login com `dev-elizabete` → tente listar objetos S3 → deve funcionar
3. Login com `devops-deploy` → tente criar EC2 → deve funcionar

---

## Limpeza

Após validação, delete os recursos criados para evitar custos.

---

## Perguntas de Revisão

1. Por que usar grupos em vez de políticas diretamente no usuário?
2. Qual a diferença entre política AWS Managed e Customer Managed?
3. Quando uma IAM Role é preferível a um usuário com Access Key?
4. O que acontece se um usuário está em dois grupos com políticas conflitantes?
