# 🪟 Instalação do Snip no Windows

## 📋 Pré-requisitos

### 1. Instalar Go

1. **Baixe o Go para Windows:**
   - Acesse: https://go.dev/dl/
   - Baixe o instalador `.msi` (ex: `go1.21.x.windows-amd64.msi`)

2. **Execute o instalador:**
   - Siga o assistente de instalação
   - O Go será instalado em `C:\Program Files\Go` por padrão
   - O instalador adiciona automaticamente ao PATH

3. **Verifique a instalação:**
   Abra um **novo** PowerShell e execute:
   ```powershell
   go version
   ```
   
   Você deve ver algo como: `go version go1.21.x windows/amd64`

## 🚀 Instalação do Snip

### Método 1: Compilar a partir do Código (Recomendado)

Você já tem o código! Siga estes passos:

1. **Abra o PowerShell** no diretório do projeto:
   ```powershell
   cd C:\repositorio\SnipAI\SnipAI
   ```

2. **Instale as dependências:**
   ```powershell
   go mod download
   ```

3. **Compile o projeto:**
   ```powershell
   go build -o snip.exe main.go
   ```

4. **Teste:**
   ```powershell
   .\snip.exe --help
   ```

5. **Adicione ao PATH (opcional mas recomendado):**

   **Opção A - Copiar para pasta no PATH:**
   ```powershell
   # Criar pasta
   New-Item -ItemType Directory -Force -Path "C:\Program Files\Snip"
   
   # Copiar executável
   Copy-Item snip.exe "C:\Program Files\Snip\snip.exe"
   
   # Adicionar ao PATH (execute como Admin)
   [Environment]::SetEnvironmentVariable("Path", $env:Path + ";C:\Program Files\Snip", "User")
   ```

   **Opção B - Adicionar diretório atual ao PATH:**
   ```powershell
   # Permanente (execute como Admin)
   [Environment]::SetEnvironmentVariable("Path", $env:Path + ";C:\repositorio\SnipAI\SnipAI", "User")
   ```

   **Opção C - Criar alias no PowerShell:**
   Edite seu perfil (`notepad $PROFILE`) e adicione:
   ```powershell
   function snip { & "C:\repositorio\SnipAI\SnipAI\snip.exe" $args }
   ```

### Método 2: Usando Scoop (Se já tiver instalado)

```powershell
scoop bucket add snip https://github.com/matheuzgomes/Snip
scoop install snip
```

### Método 3: Download do Binário Pré-compilado

1. Acesse: https://github.com/matheuzgomes/Snip/releases
2. Baixe `snip_Windows_x86_64.zip`
3. Extraia e adicione ao PATH

## ⚙️ Configuração da API Groq (Opcional)

Para usar as funcionalidades de IA:

1. **Obtenha sua API key:**
   - Acesse: https://console.groq.com/
   - Crie uma conta e gere uma API key

2. **Configure a variável de ambiente:**

   **Temporário (apenas nesta sessão):**
   ```powershell
   $env:GROQ_API_KEY="sua-chave-aqui"
   ```

   **Permanente:**
   ```powershell
   [Environment]::SetEnvironmentVariable("GROQ_API_KEY", "sua-chave-aqui", "User")
   ```

   **Nota:** Se não configurar, o Snip usará uma chave padrão (não recomendado para produção).

## ✅ Verificação

Após instalar, teste:

```powershell
# Ver ajuda
snip --help

# Criar uma nota de teste
snip create "Minha Primeira Nota" --message "Olá, Snip!"

# Listar notas
snip list

# Testar IA (se configurado)
snip ai "Hello, AI!"
```

## 🐛 Solução de Problemas

### "go: command not found"
- Instale o Go primeiro (veja seção Pré-requisitos)
- **Reinicie o PowerShell** após instalar
- Verifique: `go version`

### "snip: command not found"
- Use o caminho completo: `.\snip.exe`
- Ou adicione ao PATH (veja passo 5 da instalação)

### Erro ao compilar
```powershell
# Limpe e tente novamente
go clean -modcache
go mod download
go build -o snip.exe main.go
```

### Erro de permissão
- Execute o PowerShell como Administrador
- Ou instale em um diretório onde você tem permissão

### Erro: "cannot find package"
```powershell
go mod tidy
go mod download
go build -o snip.exe main.go
```

## 📚 Próximos Passos

1. ✅ Instale o Go
2. ✅ Compile o Snip
3. ✅ Configure a API Groq (opcional)
4. 📝 Comece a criar notas!
5. 🤖 Explore as funcionalidades de IA

Veja `AI_FEATURES.md` para mais informações sobre IA.

## 🎯 Comandos Rápidos

```powershell
# Criar nota
snip create "Título" --message "Conteúdo"

# Criar nota com IA
snip ai-create "Tópico" --tag "tag1"

# Buscar
snip find "termo"

# Busca com IA
snip ai-search "termo"

# Perguntar à IA
snip ai-ask "pergunta"

# Gerar código
snip ai-code "descrição" --lang "go"

# Ver todas as opções
snip --help
```

