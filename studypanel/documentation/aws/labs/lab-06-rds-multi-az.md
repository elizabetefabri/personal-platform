# Lab 06 — RDS Multi-AZ e Read Replica

**Dia:** 34 (Semana 5) | **Duração:** 60 minutos

---

## Objetivo

Criar um banco RDS MySQL com Multi-AZ e uma Read Replica, entendendo a diferença entre HA e escala de leitura.

---

## Atenção sobre Custos

RDS tem custo por hora. Após o lab, **delete a instância** imediatamente.

---

## Passo a Passo

### 1. Criar Subnet Group para RDS

1. RDS → Subnet Groups → Create
2. Name: `studylab-rds-subnet-group`
3. VPC: `StudyLab-VPC`
4. Adicionar subnets em pelo menos 2 AZs diferentes
5. Create

---

### 2. Criar Security Group para RDS

1. EC2 → Security Groups → Create
2. Name: `StudyLab-RDS-SG`
3. VPC: `StudyLab-VPC`
4. Inbound: MySQL/Aurora (3306) do Security Group da sua EC2

---

### 3. Criar Instância RDS MySQL

1. RDS → Create database
2. Engine: MySQL 8.0
3. Template: Free Tier (desabilita Multi-AZ por limitação, mas veja opção)
4. DB identifier: `studylab-mysql`
5. Master username: `admin`
6. Password: (anote em local seguro)
7. DB instance class: `db.t3.micro`
8. Storage: 20 GB gp2
9. **Multi-AZ**: Sim (cria réplica síncrona em outra AZ)
10. VPC: `StudyLab-VPC`
11. Subnet group: `studylab-rds-subnet-group`
12. Public access: No (apenas acesso interno)
13. Security Group: `StudyLab-RDS-SG`
14. Create database

*Aguardar ~5-10 minutos para available*

---

### 4. Conectar ao RDS via EC2

Na instância EC2 do Lab 03:
```bash
# Instalar cliente MySQL
sudo yum install -y mysql

# Conectar (substitua pelo endpoint do RDS)
mysql -h studylab-mysql.xxxxx.sa-east-1.rds.amazonaws.com \
      -u admin \
      -p

# Criar tabela de teste
CREATE DATABASE studylab;
USE studylab;

CREATE TABLE estudos (
  id INT AUTO_INCREMENT PRIMARY KEY,
  topico VARCHAR(100),
  status VARCHAR(50),
  data_estudo DATE
);

INSERT INTO estudos (topico, status, data_estudo) 
VALUES ('IAM', 'Concluído', '2024-01-15'),
       ('VPC', 'Concluído', '2024-01-22'),
       ('EC2', 'Em andamento', '2024-01-29');

SELECT * FROM estudos;
```

---

### 5. Criar Read Replica

1. RDS → Databases → Selecione `studylab-mysql`
2. Actions → Create read replica
3. DB identifier: `studylab-mysql-replica`
4. Region: Mesma região (ou outra para cross-region)
5. Instance class: `db.t3.micro`
6. Create

*Aguardar ~5-10 minutos*

**Conectar à Read Replica:**
```bash
# Use o endpoint da Read Replica (diferente do principal)
mysql -h studylab-mysql-replica.xxxxx.sa-east-1.rds.amazonaws.com \
      -u admin \
      -p -e "SELECT * FROM studylab.estudos;"
```

---

### 6. Observar o Failover Multi-AZ

1. RDS → Selecione o banco principal
2. Actions → Reboot → Reboot with failover
3. Observe que o endpoint permanece o mesmo
4. O banco ficará indisponível por ~1-2 minutos
5. Após failover, a réplica se torna o primário

---

### 7. Criar Snapshot Manual

1. RDS → Selecione `studylab-mysql`
2. Actions → Take snapshot
3. Name: `studylab-backup-manual`
4. Take snapshot

---

## Resultado Esperado

- RDS MySQL rodando com Multi-AZ
- Dados inseridos via EC2
- Read Replica funcionando e respondendo leituras
- Snapshot manual criado

---

## Validação

1. Read Replica retorna dados corretamente
2. Após failover, endpoint principal ainda funciona (mas com nova instância)
3. Snapshot aparece em RDS → Snapshots

---

## Limpeza (Importante!)

1. Delete Read Replica primeiro
2. Delete instância principal (sem snapshot final para economizar)
3. Delete Subnet Group
4. Delete Security Group

---

## Perguntas de Revisão

1. Qual o propósito do Multi-AZ vs Read Replica?
2. O que acontece com o endpoint ao fazer failover Multi-AZ?
3. Qual o limite de Read Replicas por instância RDS?
4. Como funciona o Point-in-Time Recovery no RDS?
