import { Client } from 'basic-ftp';
import * as path from 'path';
import * as fs from 'fs';
import { fileURLToPath } from 'url';
import { dirname } from 'path';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

// Load from .env.deployment or use default
const loadConfig = () => {
  const envPath = path.join(__dirname, '..', '.env.deployment');

  if (fs.existsSync(envPath)) {
    const envContent = fs.readFileSync(envPath, 'utf-8');
    const env = {};

    envContent.split('\n').forEach((line) => {
      const [key, value] = line.split('=');
      if (key && value) {
        env[key.trim()] = value.trim();
      }
    });

    return {
      host: env.FTP_HOST || 'ftp.elizabetesousafabri.com.br',
      user: env.FTP_USER || 'u485760756.elizabete',
      password: env.FTP_PASSWORD || 'Elizabetefabri123*',
      port: parseInt(env.FTP_PORT || '21', 10),
      remotePath: env.FTP_REMOTE_PATH || '/public_html/studypanel/',
      localPath: path.join(__dirname, '..', 'dist', 'frontend', 'browser'),
    };
  }

  throw new Error(
    '❌ Arquivo .env.deployment não encontrado!\n' +
    'Crie o arquivo com as seguintes variáveis:\n' +
    'FTP_HOST=\n' +
    'FTP_USER=\n' +
    'FTP_PASSWORD=\n' +
    'FTP_PORT=21\n' +
    'FTP_REMOTE_PATH=/public_html/studypanel/'
  );
};

async function deploy() {
  const client = new Client();
  client.ftp.verbose = true;

  try {
    const config = loadConfig();

    console.log('\n🔌 Conectando ao servidor FTP...');
    console.log(`📍 Host: ${config.host}`);
    console.log(`👤 User: ${config.user}`);

    await client.access({
      host: config.host,
      user: config.user,
      password: config.password,
      port: config.port,
      secure: false,
    });

    console.log('✅ Conectado ao servidor FTP!\n');

    // Ensure the remote directory exists
    console.log(`📂 Navegando para diretório remoto: ${config.remotePath}`);
    await client.ensureDir(config.remotePath);
    await client.cd(config.remotePath);

    // Clear existing files (optional, uncomment if desired)
    // console.log('🗑️  Limpando diretório remoto...');
    // await client.clearWorkingDir();

    console.log(`📤 Enviando arquivos de ${config.localPath}...`);
    await client.uploadFromDir(config.localPath);

    console.log('\n✅ Deploy concluído com sucesso!');
    console.log('🌐 Acesse sua aplicação para validar.\n');
  } catch (err) {
    if (err instanceof Error) {
      console.error('\n❌ Erro no deploy:', err.message);
    } else {
      console.error('\n❌ Erro desconhecido:', err);
    }
    process.exit(1);
  } finally {
    client.close();
  }
}

// Run deployment
deploy();
