{inputs, ...}: {
  perSystem = {pkgs, ...}: let
    devour-flake = pkgs.callPackage inputs.devour-flake {};
  in {
    checks = {
      nix-build-all = pkgs.writeShellApplication {
        name = "nix-build-all";
        runtimeInputs = [
          pkgs.nix
          devour-flake
        ];
        text = ''
          # Make sure that flake.lock is sync
          nix flake lock --no-update-lock-file

          # Do a full nix build (all outputs)
          devour-flake . "$@"
        '';
      };
    };

    devshells.default = {
      commands = [
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
  };
}
