{ lib, buildGoModule }:
buildGoModule {
  pname = "recueil";
  version = "0.0.1";

  src = ./.;

  vendorHash = "sha256-E8X9MgleXQOVnCE52xNxvPEh9ZwInrVuXrAHjSQyBrs=";

  ldflags = [
    "-s"
    "-w"
  ];

  meta = {
    description = "";
    homepage = "https://github.com/theobori/recueil";
    license = lib.licenses.gpl3Plus;
    mainProgram = "recueil";
  };
}
