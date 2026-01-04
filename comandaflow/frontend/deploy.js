const ftp = require("basic-ftp");
const path = require("path");

async function deploy() {
  const client = new ftp.Client();
  client.ftp.verbose = true;

  try {
    console.log("🔌 Conectando ao servidor FTP...");
    await client.access({
      host: "ftp.elizabetesousafabri.com.br",
      user: "u485760756.elizabete",
      password: "Bete@2025!",
      port: 21,
      secure: false,
    });

    console.log("✅ Conectado!");
    console.log("📤 Enviando arquivos da pasta out/ para /comandaflow/...");

    await client.ensureDir("/comandaflow");
    await client.clearWorkingDir();
    await client.uploadFromDir(path.join(__dirname, "out"));

    console.log("✅ Deploy concluído com sucesso!");
    console.log(
      "🌐 Site disponível em: https://comandaflow.elizabetesousafabri.com.br/"
    );
  } catch (err) {
    console.error("❌ Erro no deploy:", err);
  }

  client.close();
}

deploy();
