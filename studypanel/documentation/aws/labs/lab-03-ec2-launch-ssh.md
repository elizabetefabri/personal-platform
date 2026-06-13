# Lab 03 — EC2: Lançar Instância e Conectar via SSH

**Dia:** 20 (Semana 3) | **Duração:** 45 minutos

---

## Objetivo

Lançar uma instância EC2, anexar volume EBS e conectar via SSH.

---

## Pré-requisitos

- VPC criada (Lab 02) ou usar VPC padrão
- Par de chaves SSH criado

---

## Passo a Passo

### 1. Criar Key Pair

1. EC2 → Key Pairs → Create key pair
2. Name: `studylab-key`
3. Type: RSA
4. Format: `.pem` (Linux/Mac) ou `.ppk` (Windows PuTTY)
5. Download e salve em local seguro

---

### 2. Lançar Instância EC2

1. EC2 → Launch Instance
2. Name: `StudyLab-EC2`
3. AMI: Amazon Linux 2023 (Free Tier)
4. Instance type: `t2.micro` (Free Tier)
5. Key pair: `studylab-key`
6. Network:
   - VPC: `StudyLab-VPC` (ou default)
   - Subnet: `StudyLab-PublicSubnet`
   - Auto-assign public IP: Enable
   - Security Group: `StudyLab-PublicSG` (porta 22 liberada do seu IP)
7. Storage: 8GB gp3 (padrão)
8. User Data (opcional):
```bash
#!/bin/bash
yum update -y
yum install -y httpd
systemctl start httpd
systemctl enable httpd
echo "<h1>StudyPanel Lab EC2</h1>" > /var/www/html/index.html
```
9. Launch Instance

---

### 3. Conectar via SSH

No terminal local:
```bash
# Ajustar permissões do arquivo de chave
chmod 400 studylab-key.pem

# Conectar
ssh -i "studylab-key.pem" ec2-user@<PUBLIC_IP>
```

### Via Console AWS (EC2 Instance Connect)
1. Selecione a instância
2. Connect → EC2 Instance Connect
3. Click Connect

---

### 4. Instalar Software na Instância

```bash
# Atualizar pacotes
sudo yum update -y

# Instalar htop para monitoramento
sudo yum install -y htop

# Ver informações da instância
curl http://169.254.169.254/latest/meta-data/instance-type
curl http://169.254.169.254/latest/meta-data/placement/availability-zone
```

---

### 5. Criar e Anexar Volume EBS Adicional

1. EC2 → Volumes → Create Volume
2. Type: `gp3`
3. Size: 5 GB
4. AZ: Mesma da instância (ex: sa-east-1a)
5. Após criar → Actions → Attach Volume
6. Selecione a instância e Device: `/dev/xvdf`

**Formatar e montar no Linux:**
```bash
# Verificar disco
lsblk

# Formatar
sudo mkfs -t ext4 /dev/xvdf

# Criar diretório e montar
sudo mkdir /data
sudo mount /dev/xvdf /data

# Verificar
df -h
```

---

### 6. Criar AMI da Instância

1. EC2 → Instances → Selecione `StudyLab-EC2`
2. Actions → Image → Create Image
3. Name: `studylab-custom-ami`
4. (Aguardar criação — pode demorar alguns minutos)

---

## Resultado Esperado

- Instância EC2 rodando
- Conexão SSH funcionando
- Volume EBS anexado e montado
- AMI customizada criada

---

## Validação

1. `ssh` funcionando sem erro de permissão
2. `lsblk` mostra o volume adicional montado em `/data`
3. AMI aparece em EC2 → AMIs

---

## Limpeza

1. Terminate Instance (dados da instance store perdidos)
2. Delete Volume (dados do EBS mantidos até deletar)
3. Delete AMI (deregister + delete snapshot)

---

## Perguntas de Revisão

1. Qual a diferença entre Instance Store e EBS?
2. Por que a instância precisa estar na mesma AZ do volume EBS?
3. O que é o EC2 Instance Metadata Service (IMDS)?
4. O que acontece com o EBS quando a instância é terminada?
