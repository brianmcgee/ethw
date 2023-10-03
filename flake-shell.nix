{
  perSystem = {
    self',
    pkgs,
    ...
  }: {
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
      ];
    };
  };
}
