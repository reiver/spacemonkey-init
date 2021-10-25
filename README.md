# spacemonkey-init

Initialize spacemonkey instillation.

For more about spacemonkey see https://github.com/reiver/spacemonkey

## Usage

```
spacemonkey init
```

```
spacemonkey init -v
```

## Instillation

Usually, this will result in the creation of the following directories:

* `~/.local/share/gemini-protocol`
* `~/.local/share/rook-protocol`

However, the `$XDG_DATA_HOME`, `$XDG_DATA_DIRS`, and `$HOME` environment-variables can affect the exact locations of those `gemini-protocol/`,and `rook-protocol/` directories.
