# StudyPanel - Deployment Guide (FTP)

This document explains how to deploy the StudyPanel frontend to a Hostinger FTP server using the included TypeScript deployment script.

## Files

- `deploy.ts` — TypeScript script that uses `basic-ftp` to upload the built `dist/frontend/browser/` folder to the remote server.
- `.env.deployment.example` — template file with the necessary variables. Copy to `.env.deployment` and fill with real credentials.

## Required environment

- Node.js
- npm
- `npx` (comes with npm)
- (Optional) `ts-node` to run TypeScript scripts directly

## Configure

1. Copy the example file:

```bash
cp .env.deployment.example .env.deployment
```

2. Edit `.env.deployment` and fill the values:

```
FTP_HOST=ftp.seudominio.com.br
FTP_USER=usuario_ftp
FTP_PASSWORD=senha_ftp_aqui
FTP_PORT=21
FTP_REMOTE_PATH=/public_html/studypanel/
```

Ensure this file is NOT committed to Git.

## Deploy steps

1. Build production bundle:

```bash
npm run build:prod
```

2. Run deploy script (two options):

- Using npx (no global install):

```bash
npx ts-node deploy.ts
```

- Or using the package.json script:

```bash
npm run deploy
```

The script will:

- Read `.env.deployment` for credentials
- Connect to FTP
- Ensure the target directory exists
- Upload files from `dist/frontend/browser/`

## Troubleshooting

- If connection fails, check FTP credentials and port.
- If upload fails, check remote permissions and remote path.
- Enable verbose logging in `deploy.ts` by setting `client.ftp.verbose = true`.

## Security

Never commit `.env.deployment` to the repository. Use `.env.deployment.example` as the template.
