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

  vendorHash = "sha256-EZvrGnZ3NGoG7mAEE3sapclkQId4nTaQqaRenL0r64A=";

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
    description = "";
    homepage = "https://github.com/theobori/recueil";
    license = lib.licenses.gpl3Plus;
    mainProgram = "recueil";
  };
}
