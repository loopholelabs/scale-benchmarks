package main

import (
	"fmt"
	"github.com/loopholelabs/scale/go/tests/harness"
	rustCompile "github.com/loopholelabs/scale/rust/compile"
	"github.com/loopholelabs/scalefile"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"os/exec"
	"path"
	"testing"
)

func RustSetup(t testing.TB, modules []*harness.Module, dependencies []*scalefile.Dependency) map[*harness.Module]string {
	cargo, err := exec.LookPath("cargo")
	require.NoError(t, err, "cargo not found in path")

	t.Cleanup(func() {
		for _, module := range modules {
			moduleDir := path.Dir(module.Path)
			err := os.RemoveAll(path.Join(moduleDir, fmt.Sprintf("%s-%s-build", module.Name, t.Name())))
			if !os.IsNotExist(err) {
				require.NoError(t, err, fmt.Sprintf("failed to remove module %s", module.Name))
			}
		}
	})

	g := rustCompile.NewGenerator()

	generated := make(map[*harness.Module]string)

	for _, module := range modules {
		_, err = os.Stat(module.Path)
		require.NoError(t, err, fmt.Sprintf("module %s not found", module.Name))

		moduleDir := path.Dir(module.Path)

		err = os.Mkdir(path.Join(moduleDir, fmt.Sprintf("%s-%s-build", module.Name, t.Name())), 0755)
		if !os.IsExist(err) {
			require.NoError(t, err, fmt.Sprintf("failed to create build directory for scale function %s", module.Name))
		}

		file, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-%s-build", module.Name, t.Name()), "lib.rs"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		require.NoError(t, err, fmt.Sprintf("failed to create lib.rs for scale function %s", module.Name))

		err = g.GenerateRsLib(file, "./scale/scale.rs", module.Signature)
		require.NoError(t, err, fmt.Sprintf("failed to generate lib.rs for scale function %s", module.Name))

		err = file.Close()
		require.NoError(t, err, fmt.Sprintf("failed to close lib.rs for scale function %s", module.Name))

		cargoFile, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-%s-build", module.Name, t.Name()), "Cargo.toml"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		require.NoError(t, err, fmt.Sprintf("failed to create Cargo.toml for scale function %s", module.Name))

		err = g.GenerateRsCargo(cargoFile, dependencies, module.Signature, module.SignaturePath)
		require.NoError(t, err, fmt.Sprintf("failed to generate lib.rs for scale function %s", module.Name))

		err = cargoFile.Close()
		require.NoError(t, err, fmt.Sprintf("failed to close Cargo.toml for scale function %s", module.Name))

		err = os.Mkdir(path.Join(moduleDir, fmt.Sprintf("%s-%s-build", module.Name, t.Name()), "scale"), 0755)
		if !os.IsExist(err) {
			require.NoError(t, err, fmt.Sprintf("failed to create scale directory for scale function %s", module.Name))
		}

		scale, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-%s-build", module.Name, t.Name()), "scale", "scale.rs"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		require.NoError(t, err, fmt.Sprintf("failed to create scale.go for scale function %s", module.Name))

		file, err = os.Open(module.Path)
		require.NoError(t, err, fmt.Sprintf("failed to open scale function %s", module.Name))

		_, err = io.Copy(scale, file)
		require.NoError(t, err, fmt.Sprintf("failed to copy scale function %s", module.Name))

		err = scale.Close()
		require.NoError(t, err, fmt.Sprintf("failed to close scale.go for scale function %s", module.Name))

		err = file.Close()
		require.NoError(t, err, fmt.Sprintf("failed to close scale function %s", module.Name))

		wd, err := os.Getwd()
		require.NoError(t, err, fmt.Sprintf("failed to get working directory for scale function %s", module.Name))

		cmd := exec.Command(cargo, "build", "--release", "--target", "wasm32-unknown-unknown", "--manifest-path", "Cargo.toml")
		cmd.Dir = path.Join(wd, moduleDir, fmt.Sprintf("%s-%s-build", module.Name, t.Name()))

		err = cmd.Run()
		require.NoError(t, err, fmt.Sprintf("wd:  %s", cmd.Dir))

		generated[module] = path.Join(cmd.Dir, "target/wasm32-unknown-unknown/release/compile.wasm")
	}

	return generated
}
