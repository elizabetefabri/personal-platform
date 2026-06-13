# Lab 04 — S3: Bucket, Policies e Site Estático

**Dia:** 28 (Semana 4) | **Duração:** 40 minutos

---

## Objetivo

Criar bucket S3, configurar policies, versionamento e hospedar um site estático.

---

## Passo a Passo

### 1. Criar Bucket S3

1. S3 → Create bucket
2. Name: `studylab-[SEU-NOME]-2024` (único globalmente)
3. Region: `sa-east-1`
4. Block Public Access: **Desmarcar** "Block all public access" (para site estático)
5. Versioning: Habilitar
6. Create bucket

---

### 2. Fazer Upload de Arquivos

Crie localmente um `index.html`:
```html
<!DOCTYPE html>
<html>
<head><title>StudyPanel AWS Lab</title></head>
<body>
  <h1>Elizabete — AWS SAA-C03 Study</h1>
  <p>S3 Static Website Lab</p>
</body>
</html>
```

1. Acesse o bucket criado
2. Upload → Selecione `index.html`
3. Confirmar upload

---

### 3. Configurar Versionamento

1. Crie um arquivo `index.html` com conteúdo diferente
2. Faça upload novamente (sobrescreve com nova versão)
3. Em Objects → Selecione `index.html` → Versions
4. Observe as versões criadas

---

### 4. Aplicar Bucket Policy

1. Permissions → Bucket Policy → Edit
2. Cole a policy abaixo:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "PublicReadGetObject",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::studylab-[SEU-NOME]-2024/*"
    }
  ]
}
```

---

### 5. Configurar Site Estático

1. Properties → Static website hosting → Edit
2. Enable
3. Index document: `index.html`
4. Error document: `error.html` (opcional)
5. Save

Acesse o endpoint fornecido (ex: `http://studylab-xxx.s3-website-sa-east-1.amazonaws.com`)

---

### 6. Configurar Lifecycle Rule

1. Management → Lifecycle rules → Create
2. Name: `MoveToGlacier`
3. Prefix: (vazio = aplica a todos)
4. Transitions:
   - Current versions: S3 Standard-IA após 30 dias
   - Current versions: S3 Glacier Flexible após 90 dias
5. Create

---

### 7. Criar Presigned URL

Via AWS CLI:
```bash
# Configurar credenciais se necessário
aws configure

# Gerar presigned URL (válida por 1 hora)
aws s3 presign s3://studylab-[SEU-NOME]-2024/index.html --expires-in 3600
```

A URL gerada permite acesso temporário sem credenciais.

---

## Resultado Esperado

- Site estático acessível publicamente
- Versionamento habilitado com múltiplas versões
- Lifecycle rule configurada
- Presigned URL gerada

---

## Validação

1. Acesse o endpoint do site estático e veja a página carregando
2. Verifique as versões do `index.html` em Objects
3. Acesse a presigned URL antes de expirar

---

## Limpeza

1. Deletar objetos (incluindo versões)
2. Deletar bucket

---

## Perguntas de Revisão

1. Por que o bucket name deve ser único globalmente?
2. Qual a diferença entre SSE-S3, SSE-KMS e SSE-C?
3. Quando usar S3 Transfer Acceleration?
4. Qual é o tamanho máximo de um objeto no S3?
