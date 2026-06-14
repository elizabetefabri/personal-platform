#!/usr/bin/env bash
# Seed de dados de demonstração para o item "Serviços: EC2"
# Executar com: bash docker/seed-ec2.sh

set -euo pipefail
BASE="http://localhost:8090"
ITEM_ID="6a2dc606d11c04d0ab5e6f41"

echo "🌱 Seed: Anotações..."

curl -s -X POST "$BASE/api/v1/study-items/$ITEM_ID/notes" \
  -H "Content-Type: application/json" \
  -d '{
    "date": "2026-06-13",
    "title": "Tipos de instâncias EC2",
    "description": "General Purpose (t3, m5): balanceado CPU/mem. Compute Optimized (c5): alto CPU para batch. Memory Optimized (r5): bancos in-memory. Storage Optimized (i3): IOPS alto. GPU (p3): ML/deep learning.",
    "progress": 60,
    "status": "Revisado",
    "link": "https://aws.amazon.com/ec2/instance-types/"
  }' | python3 -c "import json,sys; d=json.load(sys.stdin); print('  ✓ Nota criada:', d['data']['title'])"

curl -s -X POST "$BASE/api/v1/study-items/$ITEM_ID/notes" \
  -H "Content-Type: application/json" \
  -d '{
    "date": "2026-06-13",
    "title": "Modelos de preço EC2",
    "description": "On-Demand: paga por hora, sem compromisso. Reserved (1-3 anos): até 72% desconto. Spot: bid no excedente AWS, até 90% desconto mas pode ser interrompido. Dedicated Host: servidor físico dedicado (compliance).",
    "progress": 40,
    "status": "Rascunho",
    "link": "https://aws.amazon.com/ec2/pricing/"
  }' | python3 -c "import json,sys; d=json.load(sys.stdin); print('  ✓ Nota criada:', d['data']['title'])"

echo "🌱 Seed: Recursos..."

curl -s -X POST "$BASE/api/v1/study-items/$ITEM_ID/resources" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Documentação oficial AWS EC2",
    "url": "https://docs.aws.amazon.com/ec2/",
    "type": "Documentação",
    "description": "Guia completo do Amazon EC2 com todos os tipos, preços e funcionalidades."
  }' | python3 -c "import json,sys; d=json.load(sys.stdin); print('  ✓ Recurso criado:', d['data']['title'])"

curl -s -X POST "$BASE/api/v1/study-items/$ITEM_ID/resources" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "AWS EC2 – Curso Stephane Maarek (Udemy)",
    "url": "https://www.udemy.com/course/aws-certified-solutions-architect-associate-saa-c03/",
    "type": "Vídeo",
    "description": "Seção sobre EC2 com labs práticos e exemplos de arquitetura real."
  }' | python3 -c "import json,sys; d=json.load(sys.stdin); print('  ✓ Recurso criado:', d['data']['title'])"

curl -s -X POST "$BASE/api/v1/study-items/$ITEM_ID/resources" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Tutorials Dojo – EC2 Cheat Sheet",
    "url": "https://tutorialsdojo.com/amazon-elastic-compute-cloud-amazon-ec2/",
    "type": "Artigo",
    "description": "Resumo rápido dos pontos mais cobrados na prova SAA-C03 sobre EC2."
  }' | python3 -c "import json,sys; d=json.load(sys.stdin); print('  ✓ Recurso criado:', d['data']['title'])"

echo "🌱 Seed: Quiz (5 flashcards EC2 estilo SAA-C03)..."

curl -s -X POST "$BASE/api/v1/quiz-questions" \
  -H "Content-Type: application/json" \
  -d '{
    "section": "cloud",
    "topic": "aws",
    "question": "Uma empresa precisa executar uma aplicação de processamento de vídeo que demanda alta performance de CPU por curtos períodos. Qual tipo de instância EC2 é mais adequado?",
    "answer": "Compute Optimized (família C) — ex: c5.xlarge",
    "explanation": "As instâncias Compute Optimized são projetadas para cargas de trabalho que exigem alto desempenho de CPU, como codificação de vídeo, processamento em lote e servidores web de alta performance.",
    "difficulty": "Médio",
    "tags": ["EC2", "tipos-de-instancia", "compute"]
  }' | python3 -c "import json,sys; d=json.load(sys.stdin); print('  ✓ Quiz criado')"

curl -s -X POST "$BASE/api/v1/quiz-questions" \
  -H "Content-Type: application/json" \
  -d '{
    "section": "cloud",
    "topic": "aws",
    "question": "Qual modelo de preço EC2 oferece o maior desconto (até 90%) mas pode ser interrompido pela AWS com aviso de 2 minutos?",
    "answer": "Spot Instances",
    "explanation": "As Spot Instances aproveitam a capacidade EC2 não utilizada da AWS. São ideais para cargas tolerantes a falhas como processamento batch, análise de dados e renderização, mas NÃO são adequadas para aplicações que precisam de disponibilidade contínua.",
    "difficulty": "Fácil",
    "tags": ["EC2", "precos", "spot"]
  }' | python3 -c "import json,sys; d=json.load(sys.stdin); print('  ✓ Quiz criado')"

curl -s -X POST "$BASE/api/v1/quiz-questions" \
  -H "Content-Type: application/json" \
  -d '{
    "section": "cloud",
    "topic": "aws",
    "question": "Uma empresa tem workloads previsíveis 24/7 por 3 anos. Qual opção oferece o maior desconto comparado ao On-Demand sem risco de interrupção?",
    "answer": "Reserved Instances — All Upfront, 3 anos (até 72% de desconto)",
    "explanation": "Reserved Instances (RI) com pagamento All Upfront por 3 anos oferecem o maior desconto fixo. Diferente das Spot, as RIs não são interrompidas. Use para workloads estáveis e previsíveis como servidores de banco de dados de produção.",
    "difficulty": "Médio",
    "tags": ["EC2", "precos", "reserved"]
  }' | python3 -c "import json,sys; d=json.load(sys.stdin); print('  ✓ Quiz criado')"

curl -s -X POST "$BASE/api/v1/quiz-questions" \
  -H "Content-Type: application/json" \
  -d '{
    "section": "cloud",
    "topic": "aws",
    "question": "Qual é a diferença entre um Security Group e um Network ACL no contexto de instâncias EC2?",
    "answer": "Security Group: stateful, opera no nível da instância. NACL: stateless, opera no nível da subnet.",
    "explanation": "Security Groups são stateful (retorno automático do tráfego), aplicados à instância, só permitem regras ALLOW. NACLs são stateless (retorno deve ser explicitamente permitido), aplicadas à subnet, permitem regras ALLOW e DENY. A prova frequentemente testa a diferença entre stateful e stateless.",
    "difficulty": "Difícil",
    "tags": ["EC2", "segurança", "networking", "VPC"]
  }' | python3 -c "import json,sys; d=json.load(sys.stdin); print('  ✓ Quiz criado')"

curl -s -X POST "$BASE/api/v1/quiz-questions" \
  -H "Content-Type: application/json" \
  -d '{
    "section": "cloud",
    "topic": "aws",
    "question": "Uma aplicação precisa de armazenamento temporário de alta velocidade para processamento intermediário. Os dados podem ser perdidos se a instância for interrompida. Qual solução usar?",
    "answer": "EC2 Instance Store (armazenamento de instância efêmero)",
    "explanation": "O Instance Store oferece armazenamento temporário de latência muito baixa conectado fisicamente ao host da instância. Os dados NÃO persistem após stop/terminate. Ideal para buffers, caches e dados temporários. Para persistência, use EBS.",
    "difficulty": "Médio",
    "tags": ["EC2", "storage", "instance-store", "EBS"]
  }' | python3 -c "import json,sys; d=json.load(sys.stdin); print('  ✓ Quiz criado')"

echo "🌱 Seed: Histórico de sessão de exemplo..."

curl -s -X POST "$BASE/api/v1/study-items/$ITEM_ID/sessions" \
  -H "Content-Type: application/json" \
  -d '{
    "startedAt": "2026-06-13T20:00:00Z",
    "endedAt": "2026-06-13T21:00:00Z",
    "durationSeconds": 3600,
    "topic": "Tipos e preços de instâncias EC2",
    "notes": "Entendi bem os tipos de instância. Ainda tenho dúvidas sobre quando usar Reserved vs Savings Plans.",
    "rating": 4,
    "milestonesCompleted": 1
  }' | python3 -c "import json,sys; d=json.load(sys.stdin); print('  ✓ Sessão criada')"

echo ""
echo "✅ Seed EC2 completo!"
