# Footer Component

## Informações Básicas

- **Selector:** `app-footer`
- **Tipo:** Standalone Component
- **Localização:** `src/app/shared/components/footer/`
- **Dependências:** Nenhuma

## Descrição

Componente de rodapé simples da aplicação. Atualmente é um placeholder para expansão futura.

## Funcionalidades

- ✅ Estrutura base para rodapé
- 📋 Pronto para expansão com links, copyright, etc.

## Conteúdo

Atualmente vazio, consultear `footer.component.html` para layout customizado.

## Arquivo de Estilos

- `footer.component.scss`

## Exemplo de Uso

```typescript
// Em app.component.ts
<app-footer></app-footer>
```

## Possíveis Expansões

- Links de redes sociais
- Copyright e informações legais
- Links de contato/suporte
- Informações de versão da aplicação

## Integração com Banco de Dados

Para o futuro banco de dados, considerar:

```sql
CREATE TABLE footer_content (
  id INT PRIMARY KEY AUTO_INCREMENT,
  section VARCHAR(50),
  content TEXT,
  link VARCHAR(255),
  order INT,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## Notas Técnicas

- Componente simples sem lógica complexa
- Standalone, sem dependências externas
