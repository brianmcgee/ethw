{
  description = "ethw / Ethereum Wallet Generator";

  nixConfig = {
    extra-substituters = [
      "https://nix-community.cachix.org"
    ];
    extra-trusted-public-keys = [
      "nix-community.cachix.org-1:mB9FSh9qf2dCimDSUo8Zy7bkq5CX+/rkCWyvRCYg3Fs="
    ];
  };

  inputs = {
    # nixpkgs
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";

    # flake-parts
    flake-parts = {
      url = "github:hercules-ci/flake-parts";
      inputs.nixpkgs-lib.follows = "nixpkgs";
    };
    flake-root.url = "github:srid/flake-root";

    # Utils
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    devshell = {
      url = "github:numtide/devshell";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    flake-compat = {
      url = "github:nix-community/flake-compat";
      flake = false;
    };
    haumea = {
      url = "github:nix-community/haumea/v0.2.2";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = inputs @ {
    flake-parts,
    haumea,
    ...
  }: let
    localInputs = haumea.lib.load {
      src = ./.;
      loader = haumea.lib.loaders.path;
    };
  in
    flake-parts.lib.mkFlake
    {
      inherit inputs;
    } {
      imports = [
        inputs.devshell.flakeModule
        inputs.flake-parts.flakeModules.easyOverlay
        inputs.flake-root.flakeModule
        inputs.treefmt-nix.flakeModule
        localInputs.packages
      ];

      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];

      debug = false;

      perSystem = {
        self',
        pkgs,
        config,
        ...
      }: {
        # shell
        devshells.default = {
          name = "ethw";
          env = [
            {
              name = "GOROOT";
              value = pkgs.go + "/share/go";
            }
            {
              name = "LD_LIBRARY_PATH";
              value = "$DEVSHELL_DIR/lib";
            }
          ];
          packages = with pkgs; [
            go
            go-tools
            delve
            golangci-lint
          ];
          commands = [
            {
              category = "dev";
              package = self'.packages.ethw;
            }

            {
              category = "nix";
              name = "fmt";
              help = "Format the source tree";
              command = "nix fmt";
            }

            {
              name = "check";
              help = "run all linters and build all packages";
              category = "checks";
              command = "nix flake check";
            }
            {
              name = "fix";
              help = "Remove unused nix code";
              category = "checks";
              command = "${pkgs.deadnix}/bin/deadnix -e $PRJ_ROOT";
            }
          ];
        };

        # formatter
        treefmt.config = {
          inherit (config.flake-root) projectRootFile;
          flakeFormatter = true;
          flakeCheck = true;
          programs = {
            alejandra.enable = true;
            deadnix.enable = true;
            deno.enable = true;
            gofumpt.enable = true;
            mdformat.enable = true;
            shfmt.enable = true;
          };
          settings.formatter = {
            deno.excludes = ["*.md"];
          };
        };

        # checks
      };
    };
}
