{
  inputs = {
    nixpkgs.url = "nixpkgs/nixpkgs-unstable";
    utils.url = "github:numtide/flake-utils";
    godoc-coverage.url = "github:mtoohey31/godoc-coverage";
    gow-src = {
      url = "github:mitranim/gow";
      flake = false;
    };
  };

  outputs = { self, nixpkgs, utils, gow-src, ... }@inputs:
    utils.lib.eachDefaultSystem (system:
      with import nixpkgs
        {
          overlays = [
            (final: prev: {
              gow = final.buildGo118Module rec {
                pname = "gow";
                version = "0.1.0";
                src = gow-src;
                vendorSha256 = "o6KltbjmAN2w9LMeS9oozB0qz9tSMYmdDW3CwUNChzA=";
              };
            })
            inputs.godoc-coverage.overlays.default
          ];
          inherit system;
        }; {
        devShells = rec {
          ci = mkShell { packages = [ go_1_18 godoc-coverage mdsh ]; };

          default = mkShell {
            packages = ci.nativeBuildInputs ++ [ gopls gow ];
          };
        };
      });
}
