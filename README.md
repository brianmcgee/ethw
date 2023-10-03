# ethw - Ethereum Wallet Generator

This Go application is designed to generate Ethereum wallets using deterministic seeds or BIP-39 mnemonics. It offers both command-line arguments and a TOML configuration file for seed input. The application outputs wallet information including private key, public key, and Ethereum address either in plain text or JSON format.

**NOTE**: Use this tool mainly for developing and for testing!

## Command-line Arguments

```bash
Usage: ethw [<seeds> ...]

Arguments:
  [<seeds> ...]    Deterministic seeds or BIP-39 mnemonics to generate keys.

Flags:
  -h, --help                  Show context-sensitive help.
  -j, --json                  Output results in JSON format
  -c, --config-file=STRING    Path to TOML config file containing seeds
```

## Usage

To generate a wallet from seed:

```bash
$ ethw "crouch apology feel panda curtain remind text dignity knee empty sibling radar"
```

To generate multiple wallets from seeds (also you can add aliases to generated wallets by using `|`):

```bash
$ ethw "Hermione Granger|crouch apology feel panda curtain remind text dignity knee empty sibling radar" "Harry Potter | radar sibling empty knee dignity text remind curtain panda feel apology crouch"
```

You can also skip specifying the whole seed:

```bash
$ ethw "Ron Weasly|crouch"
```

To display the output in `JSON` format (useful for concatenating the output with `jq`):

```bash
$ ethw -j "crouch apology feel panda curtain remind text dignity knee empty sibling radar"
```

## License

Please refer to the LICENSE file for information on how the code in this repository is licensed.
