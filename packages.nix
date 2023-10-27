{
  perSystem = {
    lib,
    pkgs,
    self',
    ...
  }: {
    packages = rec {
      ethw = pkgs.buildGoModule rec {
        pname = "ethw";
        version = "0.3.0";

        src = lib.cleanSource ./.;

        vendorSha256 = "sha256-7+f5cAPjl/t+GfXJMLLNXmOF/uqO9NdwHkx3ktaUTg8=";

        ldflags = [
          "-s"
          "-w"
          "-X 'github.com/aldoborrero/ethw/internal/build.Name=${pname}'"
          "-X 'github.com/aldoborrero/ethw/internal/build.Version=${version}'"
        ];

        subPackages = ["cmd/ethw.go"];

        meta = with lib; {
          description = "ethw - Ethereum Wallet Generator";
          homepage = "https://github.com/aldoborrero/ethw";
          license = licenses.mit;
          mainProgram = "ethw";
          maintainers = with maintainers; [aldoborrero];
        };
      };

      default = ethw;
    };

    overlayAttrs = self'.packages;
  };
}
