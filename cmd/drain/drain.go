package drain

import (
	"time"

	"github.com/lunarway/dextre/pkg/drain"
	"github.com/lunarway/dextre/pkg/kubernetes"
	"github.com/spf13/cobra"
)

var (
	nodeName    string
	gracePeriod time.Duration
)

// NewCommand sets up the move command
func NewCommand(kubectl *kubernetes.Client) *cobra.Command {
	c := &cobra.Command{
		Use:   "drain",
		Short: "",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return drain.Run(kubectl, nodeName, gracePeriod)
		},
	}
	c.Flags().StringVar(&nodeName, "node", "", "The node that lander should drain in a safe manner (required)")
	c.MarkFlagRequired("node")
	c.Flags().DurationVar(&gracePeriod, "grace-period", (30 * time.Second), "pod grace-period")

	return c
}