let

  pkgs = import (builtins.fetchTarball {
    url = "https://github.com/NixOS/nixpkgs/archive/21.05.tar.gz";
  }) {};

  betaPkgs = import (builtins.fetchTarball {
    url = "https://github.com/NixOS/nixpkgs/archive/nixpkgs-unstable.tar.gz";
  }) {};

in

  pkgs.stdenv.mkDerivation rec {

    pname = "taste";
    version = "0.1";

    buildInputs = (with pkgs; [
      # javascript
      nodejs-14_x
    ]) ++ (with betaPkgs; [
      nodePackages.zx
    ]);

    buildPhase = ''
    '';
    installPhase = ''
    '';
    shellHook = ''
    '';

    # envs
    CABIN_ROOT = builtins.toString ./.;
    PATH       = "${CABIN_ROOT}/bin:${CABIN_ROOT}/sbin:${builtins.getEnv "PATH"}";
  }
