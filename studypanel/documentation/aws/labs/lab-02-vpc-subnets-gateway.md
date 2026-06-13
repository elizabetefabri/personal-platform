# Lab 02 — VPC: Subnets, Internet Gateway e Route Tables

**Dia:** 13 (Semana 2) | **Duração:** 50 minutos

---

## Objetivo

Criar uma VPC completa com subnets públicas e privadas, Internet Gateway e NAT Gateway.

---

## Arquitetura que Será Criada

```
Internet
    │
    ▼
Internet Gateway
    │
    ▼
VPC: 10.0.0.0/16
    │
    ├── Subnet Pública (10.0.1.0/24) — AZ: sa-east-1a
    │       └── NAT Gateway
    │
    └── Subnet Privada (10.0.2.0/24) — AZ: sa-east-1a
            └── (recursos que não devem ser expostos)
```

---

## Passo a Passo

### 1. Criar a VPC

1. VPC → Your VPCs → Create VPC
2. Name: `StudyLab-VPC`
3. IPv4 CIDR: `10.0.0.0/16`
4. Tenancy: Default
5. Create VPC

---

### 2. Criar as Subnets

**Subnet Pública:**
1. VPC → Subnets → Create subnet
2. VPC: `StudyLab-VPC`
3. Name: `StudyLab-PublicSubnet`
4. AZ: `sa-east-1a`
5. IPv4 CIDR: `10.0.1.0/24`

**Subnet Privada:**
1. Name: `StudyLab-PrivateSubnet`
2. AZ: `sa-east-1a`
3. IPv4 CIDR: `10.0.2.0/24`

---

### 3. Criar e Anexar Internet Gateway

1. VPC → Internet Gateways → Create
2. Name: `StudyLab-IGW`
3. Após criar → Actions → Attach to VPC → `StudyLab-VPC`

---

### 4. Configurar Route Tables

**Route Table Pública:**
1. VPC → Route Tables → Create
2. Name: `StudyLab-PublicRT`
3. VPC: `StudyLab-VPC`
4. Routes → Edit → Add route:
   - Destination: `0.0.0.0/0`
   - Target: `StudyLab-IGW`
5. Subnet Associations → Associar `StudyLab-PublicSubnet`

**Route Table Privada:**
1. Name: `StudyLab-PrivateRT`
2. VPC: `StudyLab-VPC`
3. (Rota para NAT Gateway será adicionada após criar o NAT)
4. Subnet Associations → Associar `StudyLab-PrivateSubnet`

---

### 5. Criar NAT Gateway

1. VPC → NAT Gateways → Create
2. Name: `StudyLab-NAT`
3. Subnet: `StudyLab-PublicSubnet` (deve ficar na pública!)
4. Connectivity: Public
5. Elastic IP: Allocate Elastic IP → Aguardar status Available

**Adicionar rota para NAT na Route Table Privada:**
1. `StudyLab-PrivateRT` → Routes → Add route
2. Destination: `0.0.0.0/0`
3. Target: `StudyLab-NAT`

---

### 6. Criar Security Groups

**SG para servidores públicos (ex: ALB):**
1. VPC → Security Groups → Create
2. Name: `StudyLab-PublicSG`
3. VPC: `StudyLab-VPC`
4. Inbound: HTTP (80) e HTTPS (443) de `0.0.0.0/0`

**SG para servidores privados (ex: EC2):**
1. Name: `StudyLab-PrivateSG`
2. Inbound: SSH (22) de `StudyLab-PublicSG` apenas

---

## Resultado Esperado

- VPC com CIDR /16 criada
- 2 subnets (pública e privada) em subnets distintas
- IGW anexado e roteado para subnet pública
- NAT Gateway na subnet pública
- Route Tables corretas para cada subnet
- Security Groups configurados

---

## Validação

1. Confirme que a subnet pública tem rota `0.0.0.0/0 → IGW`
2. Confirme que a subnet privada tem rota `0.0.0.0/0 → NAT`
3. (Opcional) Lance uma EC2 em cada subnet e valide conectividade

---

## Limpeza (Importante — Custo!)

**NAT Gateway tem custo por hora + GB transferido.**

1. Delete NAT Gateway
2. Release Elastic IP
3. Delete IGW (desannexar antes)
4. Delete Subnets
5. Delete Route Tables (exceto main)
6. Delete VPC

---

## Perguntas de Revisão

1. Por que o NAT Gateway deve ficar na subnet pública?
2. Qual a diferença entre Security Group e NACL?
3. O que acontece com uma instância privada se o NAT Gateway for deletado?
4. Por que VPC Peering não é transitivo?
