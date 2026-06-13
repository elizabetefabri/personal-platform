# CloudFront e Route 53 — CDN e DNS

**Incidência:** 🟠 Alta | **Domínio:** Rede / Edge

---

## CloudFront — Content Delivery Network

### Conceito
CloudFront distribui conteúdo via rede de edge locations globais, reduzindo latência.

Analogia: em vez de buscar pizza de São Paulo quando está no Rio, o CloudFront coloca uma réplica da pizza em uma "cozinha local" mais perto de você.

### Componentes

| Componente | Descrição |
|-----------|-----------|
| **Distribution** | Configuração do CloudFront |
| **Origin** | Fonte dos dados (S3, ALB, EC2, HTTP custom) |
| **Edge Location** | Ponto de presença (400+ globalmente) |
| **Regional Edge Cache** | Cache intermediário (13 locais) |
| **Cache Behavior** | Regras de cache por path |

### Origins Suportadas
- S3 Bucket (com OAC para acesso privado)
- ALB / NLB
- EC2 instance
- Custom HTTP server

### OAC — Origin Access Control
- Substitui OAI (legado)
- Restringe acesso ao S3 apenas via CloudFront
- Bucket policy permite apenas o CloudFront como principal

### Cache

| Configuração | Descrição |
|-------------|-----------|
| **TTL** | Tempo de vida do cache (default 24h) |
| **Cache Policy** | Define quais headers/query strings afetam o cache |
| **Cache Invalidation** | Força remoção do cache (custo por path) |

### Segurança com CloudFront
- **HTTPS**: Certificate Manager (ACM) em us-east-1
- **WAF**: Web ACL aplicado na distribuição
- **Geo Restriction**: bloquear/permitir países
- **Signed URLs**: acesso a objetos privados por tempo limitado
- **Signed Cookies**: acesso a múltiplos objetos privados

---

## Route 53 — DNS

### Conceito
Route 53 é o serviço de DNS da AWS — traduz nomes de domínio em endereços IP.

### Record Types

| Record | Uso |
|--------|-----|
| **A** | Nome → IPv4 |
| **AAAA** | Nome → IPv6 |
| **CNAME** | Nome → outro nome (não funciona para root domain) |
| **Alias** | Nome → recurso AWS (funciona para root domain) |
| **MX** | Email |
| **TXT** | Verificação de domínio |

**Alias vs CNAME:**
- **Alias** é gratuito e funciona para `example.com` (root/apex)
- **CNAME** tem custo e não funciona para root domain
- Use Alias para ALB, CloudFront, S3 static website, RDS

### Routing Policies

| Política | Comportamento |
|---------|--------------|
| **Simple** | Retorna sempre o mesmo IP |
| **Weighted** | Distribui por peso (ex: 70% / 30%) |
| **Latency** | Roteia para região com menor latência |
| **Failover** | Primary → Secondary em falha |
| **Geolocation** | Por país ou continente |
| **Geoproximity** | Por proximidade + bias configurável |
| **Multi-value** | Retorna até 8 IPs saudáveis |

### Health Checks
- Route 53 verifica saúde dos endpoints
- Suporte a HTTP, HTTPS, TCP
- Pode ser associado a registros para failover automático
- CloudWatch Alarm → Health Check (para recursos sem IP público)

---

## Casos de Uso Combinados

### S3 + CloudFront (Site Estático)
```
Usuário → CloudFront → S3 (privado via OAC)
```
- CloudFront serve HTTPS
- S3 privado, só acessível pelo CloudFront
- Melhor performance e segurança

### Multi-Region com Route 53
```
Usuário → Route 53 (Latency Policy)
              │
              ├─→ us-east-1 (ALB → EC2)
              └─→ sa-east-1 (ALB → EC2)
```
- Route 53 direciona para a região mais próxima
- Failover automático em caso de falha

---

## Pegadinhas da Prova

1. **CloudFront cache em edge** — mudanças na origin não refletem imediatamente
2. **Invalidação de cache tem custo** — planeje well ou use versionamento de URL
3. **Route 53 Alias é gratuito** — CNAME tem custo por query
4. **CNAME não funciona para apex domain** — `example.com` precisa de Alias
5. **CloudFront em us-east-1** — o certificado ACM deve ser criado em `us-east-1`
6. **Geolocation ≠ Latency** — geo é por país; latency é por velocidade de rede
7. **Failover Policy** requer health checks configurados

---

## Caso Real

### Amazon Prime
O site usa CloudFront para servir assets estáticos (JS, CSS, imagens) de edge locations globais. Route 53 com política de latency garante que usuários de cada região conectem na infraestrutura mais próxima.
