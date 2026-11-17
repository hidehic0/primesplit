{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    git-hooks-nix = {
      url = "github:cachix/git-hooks.nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };
  outputs =
    inputs@{
      nixpkgs,
      git-hooks-nix,
      flake-parts,
      ...
    }:
    let
      systems = [
        "aarch64-darwin"
        "aarch64-linux"
        "x86_64-darwin"
        "x86_64-linux"
      ];
    in
    flake-parts.lib.mkFlake { inherit inputs; } {
      inherit systems;
      imports = [
        inputs.git-hooks-nix.flakeModule
      ];
      perSystem =
        {
          config,
          self',
          inputs',
          pkgs,
          lib,
          ...
        }:
        let
          # You should change package name and version
          executable = pkgs.buildGoModule {
            name = "primesplit";
            src = ./.;
            vendorHash = "sha256-7zrX8WG5u11xc6BDMn5wyHk5npn1IQoljZrR7GMD5CM=";
          };
        in
        {
          # pre-commit config
          pre-commit.settings.hooks = {
            gitleaks = {
              enable = true;
              entry = "${pkgs.gitleaks}/bin/gitleaks protect --staged";
            };
          };

          # devshells
          devShells.default = pkgs.mkShell {
            shellHook = ''
              ${config.pre-commit.shellHook}
            '';
            packages = with pkgs; [ gitleaks ];
          };

          packages.default = executable;
        };
    };
}
