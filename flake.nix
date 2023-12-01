{
  description = "A simple development environment for Go";

  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }: {

    devShell.x86_64-linux = with nixpkgs.legacyPackages.x86_64-linux; mkShell {
      buildInputs = [
        go
        gnumake
      ];

      shellHook = ''
        export GOPATH=$(pwd)/.go
        export PATH=$GOPATH/bin:$PATH
      '';
    };
  };
}
