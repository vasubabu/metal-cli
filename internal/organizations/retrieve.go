// Copyright © 2018 Jasmin Gacic <jasmin@stackpointcloud.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package organizations

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	var organizationID string
	// retrieveOrganizationCmd represents the retrieveOrganization command
	retrieveOrganizationCmd := &cobra.Command{
		Use:     `get -i <organization_UUID>`,
		Aliases: []string{"list"},
		Short:   "Retrieves a list of organizations or the details of the specified organization.",
		Long:    "Retrieves a list of organizations or the details of the specified organization. Details of an organization are only available to its members.",
		Example: `  # Retrieves list of the current user's organizations:
  metal organization get

  # Retrieves details of an organization:
  metal organization get -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			//listOpts := c.Servicer.ListOptions(nil, nil)
			if organizationID != "" {
				//getOpts := &packngo.GetOptions{Includes: listOpts.Includes, Excludes: listOpts.Excludes}

				org, _, err := c.Service.FindOrganizationById(context.Background(), organizationID).Execute()
				if err != nil {
					return fmt.Errorf("Could not get Organization: %w", err)
				}

				data := make([][]string, 1)

				data[0] = []string{org.GetId(), org.GetName(), org.GetCreatedAt().GoString()}
				header := []string{"ID", "Name", "Created"}

				return c.Out.Output(org, header, &data)
			} else {
				orgs, _, err := c.Service.FindOrganizations(context.Background()).Execute()
				if err != nil {
					return fmt.Errorf("Could not list Organizations: %w", err)
				}

				data := make([][]string, len(orgs.GetOrganizations()))

				for i, p := range orgs.GetOrganizations() {
					data[i] = []string{p.GetId(), p.GetId(), p.GetCreatedAt().GoString()}
				}
				header := []string{"ID", "Name", "Created"}

				return c.Out.Output(orgs, header, &data)
			}
		},
	}

	retrieveOrganizationCmd.Flags().StringVarP(&organizationID, "organization-id", "i", "", "The UUID of an organization.")
	return retrieveOrganizationCmd
}
