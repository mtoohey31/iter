{
  inputs = {
    nixpkgs.url = "nixpkgs/nixpkgs-unstable";
    utils.url = "github:numtide/flake-utils";
    gow-src = {
      url = "github:mitranim/gow";
      flake = false;
    };
  };

  outputs = { self, nixpkgs, utils, gow-src }:
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
          ];
          inherit system;
        }; {
        devShells.default = mkShell { packages = [ go_1_18 gopls gow ]; };
      });
}
