# ethw - Ethereum Wallet Generator

This Go application is designed to generate Ethereum wallets using deterministic BIP-39 mnemonics or any arbitrary seed string as a seed.

**NOTE**: Use this tool mainly for developing and for testing! If you're seeking for a more **secure** and complete solution ready to be used on `mainnet`, consider using [ethereal](https://github.com/wealdtech/ethereal) and/or [ethdo](https://github.com/wealdtech/ethdo).

## Command-line Arguments

```console
Usage: ethw <command>

Flags:
  -h, --help                 Show context-sensitive help.
      --log-level="fatal"    Configure logging level ($LOG_LEVEL)
      --log-format="text"    Configure logging format ($LOG_FORMAT)

Commands:
  wallet create <seed> ...
    Create new Ethereum wallets

  keystore create <wallets> ...
    Manage Ethereum keystores

  keystore list
    List all wallets from the keystore

  seed create
    Create a new seed

  version
    Display the application version

Run "ethw <command> --help" for more information on a command.
```

## Usage Examples

### Wallet Generation

Below you can see some examples for generating Ethereum wallets with the `wallet create` subcommand:

#### Generate an Ethereum Wallet from a Seed

Ensure to specify the `seed` parameter:

```console
$ ethw wallet create "seed=crouch apology feel panda curtain remind text dignity knee empty sibling radar"
```

You should expect output resembling:

```console
+---+-------+--------------------------------------------+------------------------------------------------------------------+------------------------------------------------------------------------------------------------------------------------------------+
| # | ALIAS | ADDRESS                                    | PRIVATE KEY                                                      | PUBLIC KEY                                                                                                                         |
+---+-------+--------------------------------------------+------------------------------------------------------------------+------------------------------------------------------------------------------------------------------------------------------------+
| 1 |       | 0x8d86D515fbee6A364C96Cf60f3220826f13A64F3 | 1d0b0a3898ff359032970f9d831269020d78463d861b305f40b1a85bed5bcefe | 04119a43acba93317d89e4a1181cbcef1a8ac28fdee7bb0df785db2510534b4a001cff289a9b70eb8d962009490c64bc546aa1fc0c880a4d608275639cab07391c |
+---+-------+--------------------------------------------+------------------------------------------------------------------+------------------------------------------------------------------------------------------------------------------------------------+
```

#### Generate Multiple Wallets with Aliases

Generate multiple wallets at once and assign aliases for better readability:

```console
$ ethw wallet create "seed=crouch apology feel panda curtain remind text dignity knee empty sibling radar;alias=Hermione Granger" "seed=radar sibling empty knee dignity text remind curtain panda feel apology crouch;alias=Harry Potter"
```

The output table will display the aliases:

```console
+---+------------------+--------------------------------------------+------------------------------------------------------------------+------------------------------------------------------------------------------------------------------------------------------------+
| # | ALIAS            | ADDRESS                                    | PRIVATE KEY                                                      | PUBLIC KEY                                                                                                                         |
+---+------------------+--------------------------------------------+------------------------------------------------------------------+------------------------------------------------------------------------------------------------------------------------------------+
| 1 | Hermione Granger | 0x8d86D515fbee6A364C96Cf60f3220826f13A64F3 | 1d0b0a3898ff359032970f9d831269020d78463d861b305f40b1a85bed5bcefe | 04119a43acba93317d89e4a1181cbcef1a8ac28fdee7bb0df785db2510534b4a001cff289a9b70eb8d962009490c64bc546aa1fc0c880a4d608275639cab07391c |
| 2 | Harry Potter     | 0x6f339aB74be047e3C5e5a784e2D4dDB5C161a034 | 130cf1653ae56b5278203d140509306fdf2f2a619ce54a64d54b688114339c8f | 04740ae95d36f6bc8b906fd4ee56cc048a0c94a323dd9cd74505e4de30ce52f4799ddc478df118b8377b0378d870014d36ae3fa98409f0a6bfd45fc9d31e54be9b |
+---+------------------+--------------------------------------------+------------------------------------------------------------------+------------------------------------------------------------------------------------------------------------------------------------+
```

#### Generate Wallets with JSON format:

You can also generate wallets and output them in `JSON` format, useful for utilities like `jq` and `dasel`:

```console
$ ethw wallet create --json "seed=crouch apology feel panda curtain remind text dignity knee empty sibling radar"
```

The output will be like it follows:

```console
[{"alias":"Hermione Granger","address":"0x8d86D515fbee6A364C96Cf60f3220826f13A64F3","private_key":"1d0b0a3898ff359032970f9d831269020d78463d861b305f40b1a85bed5bcefe","public_key":"04119a43acba93317d89e4a1181cbcef1a8ac28fdee7bb0df785db2510534b4a001cff289a9b70eb8d962009490c64bc546aa1fc0c880a4d608275639cab07391c"}]
```

Sweet!

### Keystores

This feature allows direct generation of keystores for compatibility with Geth and other execution clients.

Wallet data format:

- `seed=<Seed>`, where `<Seed>` is the seed for generating the wallet, which could be a mnemonic or an arbitrary string.
- `password=<Password>`, where `<Password>` is the password to secure the keystore (bear in mind using passwords directly on the terminal will result in password leakage).

Some examples:

#### Create a single keystore

```console
$ ethw keystore create "seed=crouch apology feel panda curtain remind text dignity knee empty sibling radar;password=1234"
```

#### Create multiple wallets in a keystore

Same as when generating wallets, you can add multiple wallets into a single `keystore`:

```console
$ ethw keystore create "seed=crouch apology feel panda curtain remind text dignity knee empty sibling radar;password=1234" "seed=radar sibling empty knee dignity text remind curtain panda feel apology crouch;password=5678"
```

#### Overwrite existing keystore

You can nuke all the contents found in a single keystore with `--overwrite` argument:

```console
$ ethw keystore create --overwrite "seed=crouch apology feel panda curtain remind text dignity knee empty sibling radar;password=1234"
```

#### Specify a custom keystore directory

By default, `ethw` will create a keystore in the current directory where you're invoke the command, but you can easily override it with `--keystore-dir`:

```console
$ ethw keystore create --keystore-dir=./my_keystore "seed=crouch apology feel panda curtain remind text dignity knee empty sibling radar;password=1234"
```

## License

Please refer to the LICENSE file for information on how the code in this repository is licensed.
