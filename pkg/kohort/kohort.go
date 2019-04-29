package kohort

import (
	"fmt"
	"github.com/n3wscott/kohort/pkg/kohort/resources"
	"github.com/spf13/cobra"
	"os"
	"path"
	"runtime"
	"strings"
)

// TODO: this is just a demo.
// If you build this, then run it, ko will ship the latest code, not the binary.
// Can ko take in golang build parameters ?

func podYaml(filename string) string {
	gosrc := os.Getenv("GOSRC")
	if gosrc == "" {
		gosrc = os.Getenv("GOPATH") + "/src"
	}

	if !strings.HasSuffix(gosrc, "/") {
		gosrc += "/"
	}
	image := strings.TrimPrefix(filename, gosrc)

	image = image[:strings.LastIndex(image, "/")]

	yaml := resources.Pod(&resources.PodSpec{
		Name:      "demo",
		Namespace: "default",
		Image:     image,
	})

	return yaml
}

func RunMaybe(fns ...ResourceFn) {

	_, filename, _, _ := runtime.Caller(1)

	var cmdKo = &cobra.Command{
		Use:   "promo",
		Short: "Enable ko deployment",
		Long:  `TODO.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			// default apply
			// default steve
			// default shipit

			var out string
			var err error
			verb := "apply"
			if len(fns) > 0 {
				out, err = Run(&KoOptions{
					Manifests: fns[0](), // TODO
					Verb:      verb,
				})
			} else {
				out, err = Run(&KoOptions{
					Manifests: podYaml(filename),
					Verb:      verb,
				})
			}

			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(out)
		},
	}

	var cmdDry = &cobra.Command{
		Use:   "dry",
		Short: "Display yaml to be used",
		Long:  `todo`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if len(fns) > 0 {
				for _, fn := range fns {
					fmt.Println(fn())
				}
			} else {
				fmt.Println(podYaml(filename))
			}
		},
	}

	var applyArgs string
	var cmdApply = &cobra.Command{
		Use:   "apply",
		Short: "Apply using ko",
		Long:  `todo`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("applying...")
			fmt.Println("with ", applyArgs)

			var out string
			var err error
			verb := "apply"
			if len(fns) > 0 {
				out, err = Run(&KoOptions{
					Manifests: fns[0](), // TODO
					Verb:      verb,
				})
			} else {
				out, err = Run(&KoOptions{
					Manifests: podYaml(filename),
					Verb:      verb,
				})
			}

			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(out)
		},
	}
	cmdApply.Flags().StringVar(&applyArgs, "args", "", "pass args as a string")

	var cmdDelete = &cobra.Command{
		Use:   "delete",
		Short: "Delete using ko",
		Long:  `todo`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("deleting...")

			var out string
			var err error
			verb := "delete"
			if len(fns) > 0 {
				out, err = Run(&KoOptions{
					Manifests: fns[0](), // TODO
					Verb:      verb,
				})
			} else {
				out, err = Run(&KoOptions{
					Manifests: podYaml(filename),
					Verb:      verb,
				})
			}

			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(out)
		},
	}

	koEnabled := true
	_, program := path.Split(os.Args[0])
	var rootCmd = &cobra.Command{
		Use:  program,
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			koEnabled = false
		},
	}
	rootCmd.AddCommand(cmdKo)
	cmdKo.AddCommand(cmdDry, cmdApply, cmdDelete)

	// Run it, maybe.
	err := rootCmd.Execute()

	if koEnabled {
		if err != nil {
			fmt.Printf("error: %v", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
}
