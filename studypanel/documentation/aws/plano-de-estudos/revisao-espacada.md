# Técnica de Revisão Espaçada — AWS SAA-C03

A revisão espaçada aumenta a retenção de 20% para 80%+ do conteúdo.

---

## Como Funciona

Após estudar um tema, revise nos seguintes intervalos:

| Revisão | Quando |
|---------|--------|
| R1 | 1 dia depois |
| R2 | 7 dias depois |
| R3 | 15 dias depois |
| R4 | 30 dias depois |

---

## Calendário de Revisão Automática

| Dia de Estudo | Tema | R1 | R2 | R3 | R4 |
|---------------|------|-----|-----|-----|-----|
| Dia 1 | Fundamentos AWS | Dia 2 | Dia 8 | Dia 16 | Dia 31 |
| Dia 2 | IAM Parte 1 | Dia 3 | Dia 9 | Dia 17 | Dia 32 |
| Dia 3 | IAM Parte 2 | Dia 4 | Dia 10 | Dia 18 | Dia 33 |
| Dia 8 | VPC Parte 1 | Dia 9 | Dia 15 | Dia 23 | Dia 38 |
| Dia 15 | EC2 Parte 1 | Dia 16 | Dia 22 | Dia 30 | Dia 45 |
| Dia 22 | S3 Parte 1 | Dia 23 | Dia 29 | Dia 37 | Dia 52 |
| Dia 29 | RDS Parte 1 | Dia 30 | Dia 36 | Dia 44 | Dia 59 |
| Dia 36 | Lambda Parte 1 | Dia 37 | Dia 43 | Dia 51 | Dia 60 |

---

## Formato da Revisão

Para cada revisão, responda as perguntas abaixo sem consultar o material:

### IAM
1. Qual a diferença entre usuário, grupo e role no IAM?
2. Quando usar IAM Role ao invés de Access Key?
3. O que é uma política baseada em recurso vs baseada em identidade?
4. Como o MFA protege as contas AWS?
5. O que é o princípio do menor privilégio?

### VPC
1. O que é um CIDR block e como dividir em subnets?
2. Diferença entre Security Group e NACL?
3. Quando usar NAT Gateway vs NAT Instance?
4. Para que serve o VPC Peering vs Transit Gateway?
5. O que é uma subnet pública vs privada?

### EC2
1. Quais as diferenças entre On-Demand, Reserved, Spot e Dedicated?
2. O que é uma AMI e como criar uma?
3. Diferença entre EBS e Instance Store?
4. O que é um Placement Group e quando usar cada tipo?
5. Como funciona o Auto Scaling?

### S3
1. Quais as classes de storage do S3 e quando usar cada uma?
2. O que é versionamento e como funciona o lifecycle?
3. Como funciona a replicação S3 (CRR vs SRR)?
4. Diferença entre bucket policy e ACL?
5. O que é um presigned URL?

---

## Regras de Revisão

- Se acertar >80%: avance normalmente, próxima revisão no intervalo padrão
- Se acertar 50-80%: revise de novo em 3 dias
- Se acertar <50%: releia o material completo e comece o ciclo do zero

---

## Flash Cards Rápidos

Crie um arquivo `flashcards.md` com perguntas e respostas para revisar em 10 minutos por dia.

Formato sugerido:

```
P: O que é o RTO?
R: Recovery Time Objective — tempo máximo aceitável para restaurar o sistema.

P: O que é o RPO?
R: Recovery Point Objective — perda máxima de dados aceitável (em tempo).

P: Qual serviço AWS é mais barato para armazenar dados raramente acessados?
R: S3 Glacier Deep Archive.
```
