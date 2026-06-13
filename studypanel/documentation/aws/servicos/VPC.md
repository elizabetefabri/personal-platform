# VPC — Virtual Private Cloud

**Incidência:** 🔴 Muito Alta | **Domínio:** Rede

---

## Conceito

VPC é uma rede privada isolada dentro da AWS. Você define o espaço de endereçamento IP, subnets, roteamento e gateways.

Analogia: VPC é o prédio; subnets são os andares; Security Groups são as portas de cada sala.

---

## Arquitetura Típica

```
Internet
    │
    ▼
Internet Gateway (IGW)
    │
    ▼
┌─────────────────────────────┐
│ VPC (10.0.0.0/16)          │
│                             │
│ ┌─────────────┐             │
│ │ Subnet Pub  │ 10.0.1.0/24│
│ │ (ALB, NAT) │             │
│ └──────┬──────┘             │
│        │                    │
│ ┌──────▼──────┐             │
│ │ Subnet Priv │ 10.0.2.0/24│
│ │ (EC2, RDS) │             │
│ └─────────────┘             │
└─────────────────────────────┘
```

---

## Componentes

### CIDR Block
- Define o range de IPs da VPC: ex `10.0.0.0/16`
- Subnets dividem o CIDR em blocos menores: ex `10.0.1.0/24`

### Subnets
| Tipo | Características |
|------|----------------|
| Pública | Tem rota para Internet Gateway |
| Privada | Não tem rota direta para internet |
| Isolada | Sem rota externa, acesso apenas interno |

### Internet Gateway (IGW)
- Conecta a VPC à internet
- Apenas 1 por VPC
- Altamente disponível e escalável automaticamente

### NAT Gateway
- Permite que instâncias privadas acessem a internet (mas não o contrário)
- Deve ser criado em **subnet pública**
- Altamente disponível dentro de uma AZ (crie um por AZ para HA)

### Route Tables
- Define para onde vai o tráfego
- Subnet pública: `0.0.0.0/0 → IGW`
- Subnet privada: `0.0.0.0/0 → NAT Gateway`

### Security Groups
- Firewall **stateful** em nível de instância
- Apenas regras de Allow (sem Deny explícito)
- Suporta referência a outros Security Groups

### NACLs (Network Access Control Lists)
- Firewall **stateless** em nível de subnet
- Suporta Allow e Deny
- Regras avaliadas em ordem numérica

---

## Diferença: Security Group vs NACL

| Característica | Security Group | NACL |
|---------------|---------------|------|
| Nível | Instância | Subnet |
| Estado | Stateful | Stateless |
| Regras | Apenas Allow | Allow e Deny |
| Aplicação | Automática (retorno) | Manual (entrada + saída) |
| Avaliação | Todas as regras | Ordem numérica |

---

## Conectividade

### VPC Peering
- Conecta duas VPCs (mesma ou contas diferentes)
- Não é transitivo: A↔B, B↔C ≠ A↔C

### Transit Gateway
- Hub central para múltiplas VPCs
- Suporta roteamento transitivo
- Mais caro, mas escala melhor

### VPN Site-to-Site
- Conexão criptografada via internet entre on-premises e AWS
- Usa Virtual Private Gateway (VGW) no lado AWS

### Direct Connect
- Conexão física dedicada (não via internet)
- Mais estável, menor latência, maior custo
- Usado para compliance e performance

### VPC Endpoints
| Tipo | Uso |
|------|-----|
| Gateway Endpoint | S3 e DynamoDB (gratuito) |
| Interface Endpoint | Outros serviços via PrivateLink |

---

## Pegadinhas da Prova

1. **NAT Gateway** fica em subnet pública, mas serve instâncias privadas
2. **VPC Peering não é transitivo** — precisa de peering explícito entre cada par
3. **NACLs são stateless** — precisa permitir tráfego de entrada E saída
4. **Security Groups não bloqueiam** — só permitem
5. **IGW é obrigatório** para subnet pública ter acesso à internet
6. **VPC Flow Logs** registram tráfego IP aceito e rejeitado — não registra conteúdo

---

## Casos Reais

### Itaú
Usa VPC com subnets privadas para todos os serviços de backend, com apenas o Load Balancer em subnet pública. Direct Connect para conectar datacenters próprios à AWS com baixa latência.

---

## Lab Relacionado

[Lab 02 — VPC: Subnets, Gateway, Route Tables](../labs/lab-02-vpc-subnets-gateway.md)
