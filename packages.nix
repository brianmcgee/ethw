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
        version = "0.1.0+dev";

        src = lib.cleanSource ./.;
        vendorSha256 = "sha256-WbaRsG7+6oLltMvC98DwLpBWdbRdYkPxfhZLmBXDVY4=";

        ldflags = [
          "-X 'build.Name=${pname}'"
          "-X 'build.Version=${version}'"
        ];

        meta = with lib; {
          description = "ethw - ethereum wallet generator";
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
