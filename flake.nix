{
  description = "iter";

  inputs = {
    nixpkgs.url = "nixpkgs/nixpkgs-unstable";
    utils.url = "github:numtide/flake-utils";
    gow-src = {
      url = "github:mitranim/gow";
      flake = false;
    };
    perf-src = {
      url = "git+https://go.googlesource.com/perf";
      flake = false;
    };
  };

  outputs = { self, nixpkgs, utils, gow-src, perf-src }:
    utils.lib.eachDefaultSystem (system: with import nixpkgs
      {
        overlays = [
          (final: _: {
            benchstat = final.buildGoModule {
              pname = "benchstat";
              version = perf-src.shortRev;
              src = perf-src;
              subPackages = [ "cmd/benchstat" ];
              vendorSha256 = "ZIkH3LBzrvqWEN6m4fpU2cmOXup9LLU3FiFooJJtiOk=";
            };
            gow = final.buildGoModule {
              pname = "gow";
              version = gow-src.shortRev;
              src = gow-src;
              vendorSha256 = "o6KltbjmAN2w9LMeS9oozB0qz9tSMYmdDW3CwUNChzA=";
            };
          })
        ]; inherit system;
      }; {
      devShells = rec {
        ci = mkShell { packages = [ benchstat go mdsh revive ]; };

        default = mkShell {
          packages = ci.nativeBuildInputs ++ [ gopls gow ];
        };
      };
    });
}
