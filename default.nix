{
  lib,
  buildGoModule,
  pkg-config,
  vips,
}:
buildGoModule {
  pname = "recueil";
  version = "0.0.1";

  src = ./.;

  vendorHash = "sha256-AXRWW6+Xx6vkixPG3ZBrfMcQA0+2wAK1qRHL5DJWpnI";

  ldflags = [
    "-s"
    "-w"
  ];

  nativeBuildInputs = [
    pkg-config
  ];

  buildInputs = [
    vips
  ];

  meta = {
    description = "My opinionated static gallery generator.";
    homepage = "https://github.com/theobori/recueil";
    license = lib.licenses.gpl3Plus;
    mainProgram = "recueil";
  };
}
