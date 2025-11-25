# Configuração da Chave de API Groq

Para usar as funcionalidades de IA do Snip, você precisa configurar a variável de ambiente `GROQ_API_KEY`.

## Como obter a chave

1. Acesse: https://console.groq.com/keys
2. Crie uma conta ou faça login
3. Gere uma nova chave de API
4. Copie a chave

## Como configurar

### Windows (PowerShell)

```powershell
# Temporário (apenas nesta sessão)
$env:GROQ_API_KEY="sua_chave_aqui"

# Permanente (adicionar ao perfil)
[Environment]::SetEnvironmentVariable("GROQ_API_KEY", "sua_chave_aqui", "User")
```

### Windows (CMD)

```cmd
# Temporário
set GROQ_API_KEY=sua_chave_aqui

# Permanente (via GUI)
# Painel de Controle > Sistema > Configurações Avançadas > Variáveis de Ambiente
```

### Linux/macOS

```bash
# Temporário
export GROQ_API_KEY="sua_chave_aqui"

# Permanente (adicionar ao ~/.bashrc ou ~/.zshrc)
echo 'export GROQ_API_KEY="sua_chave_aqui"' >> ~/.bashrc
source ~/.bashrc
```

## Verificar se está configurado

```powershell
# Windows PowerShell
echo $env:GROQ_API_KEY

# Linux/macOS
echo $GROQ_API_KEY
```

## Uso

Após configurar a variável de ambiente, todas as funcionalidades de IA estarão disponíveis:

```powershell
snip.exe ai-create "Tópico"
snip.exe ai-code "descrição" --lang "go"
snip.exe project ai-create "Meu Projeto"
snip.exe checklist ai-create "Preparação" --items 5
```

