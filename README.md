# go-htmx API

## Requirements

- `bun`
- `go`
- `gow`
- `docker`
- `docker-compose`
- `dbmate`

## Local Setup

```bash
# Install dependencies
bun install

# Run the server
bun dev
```

## Migrations

We use `dbmate`, install with `brew install dbmate`.

Run migrations

```bash
dbmate up
```

Create new migration

```bash
dbmate new <migration_name>
```

## Reloading

There is a simple websocket setup to do reloading of the frontend without manual refreshes

## Tailwind Syntax Highlighting

Make sure to add this to VSCode/GoLand's settings

Then any `Class` will get autocomplete, also any string with "tw X" at the start

```json
{
  "tailwindCSS.experimental.classRegex": [
    ["Class\\(([^)]*)\\)", "[\"`]([^\"`]*)[\"`]"], // Class("...") or Class(`...`)
    ["Classes\\(([^)]*)\\)", "[\"`]([^\"`]*)[\"`]"], // Classes("...") or Classes(`...`)
    ["Class\\{([^)]*)\\}", "[\"`]([^\"`]*)[\"`]"], // Class{"..."} or Class{`...`}
    ["Classes\\{([^)]*)\\}", "[\"`]([^\"`]*)[\"`]"], // Classes{"..."} or Classes{`...`}
    ["Class:\\s*[\"`]([^\"`]*)[\"`]"], // Class: "..." or Class: `...`
    ["Classes:\\s*[\"`]([^\"`]*)[\"`]"] // Classes: "..." or Classes: `...`
  ]
}
```

If you want to do this with nvim do this:

```lua
  {
    "neovim/nvim-lspconfig",
    config = function()
      local lspconfig = require("lspconfig")

      local capabilities = require("cmp_nvim_lsp").default_capabilities()

      lspconfig.tailwindcss.setup({
        capabilities = capabilities,
        filetypes = { "html", "typescript", "javascript", "javascriptreact", "typescriptreact", "svelte", "vue", "go" }, -- Added "go" here
        settings = {
          tailwindCSS = {
            experimental = {
              classRegex = {
                "Class\\(([^)]*)\\)",
                '["`]([^"`]*)["`]', -- Class("...") or Class(`...`)
                "Classes\\(([^)]*)\\)",
                '["`]([^"`]*)["`]', -- Classes("...") or Classes(`...`)
                "Class\\{([^)]*)\\}",
                '["`]([^"`]*)["`]', -- Class{"..."} or Class{`...`}
                "Classes\\{([^)]*)\\}",
                '["`]([^"`]*)["`]', -- Classes{"..."} or Classes{`...`}
                'Class:\\s*["`]([^"`]*)["`]', -- Class: "..." or Class: `...`
                ':\\s*["`]([^"`]*)["`]', -- Classes: "..." or Classes: `...`
              },
            },
          },
        },
      })
-- rest
```
