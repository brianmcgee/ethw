{
  perSystem = {config, ...}: {
    treefmt.config = {
      inherit (config.flake-root) projectRootFile;
      flakeFormatter = true;
      programs = {
        alejandra.enable = true;
        deadnix.enable = true;
        gofumpt.enable = true;
        mdformat.enable = true;
        prettier.enable = true;
        shfmt.enable = true;
        terraform.enable = true;
      };
      settings.formatter = {
        prettier.excludes = ["*.md"];
      };
    };

    devshells.default.commands = [
      {
        category = "Tools";
        name = "fmt";
        help = "Format the source tree";
        command = "nix fmt";
      }
    ];
  };
}
