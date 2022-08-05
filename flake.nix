{
  description = "iter";

  inputs = {
    nixpkgs.url = "nixpkgs/nixpkgs-unstable";
    utils.url = "github:numtide/flake-utils";
    gow-src = {
      url = "github:mitranim/gow";
      flake = false;
    };
    yaegi-src = {
      url = "github:traefik/yaegi";
      flake = false;
    };
  };

  outputs = { self, nixpkgs, utils, gow-src, yaegi-src }:
    utils.lib.eachDefaultSystem (system: with import nixpkgs
      {
        overlays = [
          (final: prev: {
            gow = final.buildGoModule rec {
              pname = "gow";
              version = "0.1.0";
              src = gow-src;
              vendorSha256 = "o6KltbjmAN2w9LMeS9oozB0qz9tSMYmdDW3CwUNChzA=";
            };
            yaegi-rlwrapped = final.buildGoModule rec {
              pname = "yaegi-rlwrapped";
              version = yaegi-src.shortRev;
              src = yaegi-src;
              vendorSha256 = "pQpattmS9VmO3ZIQUFn66az8GSmB4IvYhTTCFn6SUmo=";
              GOROOT = "${final.go_1_18}/share/go";
              fixupPhase = ''
                mv $out/bin/yaegi $out/bin/.yaegi-wrapped
                printf "#!${final.bash}/bin/bash\n${final.rlwrap}/bin/rlwrap $out/bin/.yaegi-wrapped \"\$@\"\n" > $out/bin/yaegi
                chmod +x $out/bin/yaegi
              '';
            };
          })
        ]; inherit system;
      }; rec {
      devShells = rec {
        ci = mkShell { packages = [ go mdsh revive ]; };

        default = mkShell {
          packages = ci.nativeBuildInputs ++ [ gopls gow yaegi-rlwrapped ];
        };
      };
    });
}
