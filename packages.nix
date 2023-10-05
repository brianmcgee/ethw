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
        version = "0.2.0";

        src = lib.cleanSource ./.;

        vendorSha256 = "sha256-Q1lW3fj0D13dZ1Ci2qiiCiFJXWb7N3IaRAA7jrIoAyY=";

        ldflags = [
          "-s"
          "-w"
          "-X 'github.com/aldoborrero/ethw/internal/build.Name=${pname}'"
          "-X 'github.com/aldoborrero/ethw/internal/build.Version=${version}'"
        ];

        subPackages = ["cmd/ethw.go"];

        meta = with lib; {
          description = "ethw - ethereum wallet utils";
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
