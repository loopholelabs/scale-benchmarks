/*
	Copyright 2023 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package compile

import (
	"fmt"
	"github.com/loopholelabs/scale/go/tests/harness"
	rustCompile "github.com/loopholelabs/scale/rust/compile"
	"github.com/loopholelabs/scalefile"
	"io"
	"os"
	"os/exec"
	"path"
)

func RustSetup(modules []*harness.Module, dependencies []*scalefile.Dependency) (map[*harness.Module]string, error) {
	cargo, err := exec.LookPath("cargo")
	if err != nil {
		return nil, fmt.Errorf("cargo not found in path: %w", err)
	}

	g := rustCompile.NewGenerator()

	generated := make(map[*harness.Module]string)

	for _, module := range modules {
		_, err = os.Stat(module.Path)
		if err != nil {
			return nil, fmt.Errorf("module %s not found: %w", module.Name, err)
		}

		moduleDir := path.Dir(module.Path)

		err = os.Mkdir(path.Join(moduleDir, fmt.Sprintf("%s-%s-build", module.Name, "http")), 0755)
		if !os.IsExist(err) {
			if err != nil {
				return nil, fmt.Errorf("failed to create build directory for scale function %s: %w", module.Name, err)
			}
		}

		file, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-%s-build", module.Name, "http"), "lib.rs"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to create lib.rs for scale function %s: %w", module.Name, err)
		}

		err = g.GenerateRsLib(file, "./scale/scale.rs", module.Signature)
		if err != nil {
			return nil, fmt.Errorf("failed to generate lib.rs for scale function %s: %w", module.Name, err)
		}

		err = file.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to close lib.rs for scale function %s: %w", module.Name, err)
		}

		cargoFile, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-%s-build", module.Name, "http"), "Cargo.toml"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to create Cargo.toml for scale function %s: %w", module.Name, err)
		}

		err = g.GenerateRsCargo(cargoFile, dependencies, module.Signature, module.SignaturePath)
		if err != nil {
			return nil, fmt.Errorf("failed to generate Cargo.toml for scale function %s: %w", module.Name, err)
		}

		err = cargoFile.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to close Cargo.toml for scale function %s: %w", module.Name, err)
		}

		err = os.Mkdir(path.Join(moduleDir, fmt.Sprintf("%s-%s-build", module.Name, "http"), "scale"), 0755)
		if !os.IsExist(err) {
			if err != nil {
				return nil, fmt.Errorf("failed to create scale directory for scale function %s: %w", module.Name, err)
			}
		}

		scale, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-%s-build", module.Name, "http"), "scale", "scale.rs"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to create scale.rs for scale function %s: %w", module.Name, err)
		}

		file, err = os.Open(module.Path)
		if err != nil {
			return nil, fmt.Errorf("failed to open scale function %s: %w", module.Name, err)
		}

		_, err = io.Copy(scale, file)
		if err != nil {
			return nil, fmt.Errorf("failed to copy scale function %s: %w", module.Name, err)
		}

		err = scale.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to close scale.rs for scale function %s: %w", module.Name, err)
		}

		err = file.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to close scale function %s: %w", module.Name, err)
		}

		wd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("failed to get working directory for scale function %s: %w", module.Name, err)
		}

		cmd := exec.Command(cargo, "build", "--release", "--target", "wasm32-unknown-unknown", "--manifest-path", "Cargo.toml")
		cmd.Dir = path.Join(wd, moduleDir, fmt.Sprintf("%s-%s-build", module.Name, "http"))

		err = cmd.Run()
		if err != nil {
			return nil, fmt.Errorf("failed to compile scale function %s: %w", module.Name, err)
		}

		generated[module] = path.Join(cmd.Dir, "target/wasm32-unknown-unknown/release/compile.wasm")
	}

	return generated, nil
}
