// Copyright © 2021-2022 Jonas Muehlmann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/spf13/cobra"
)

func WithBookmarkTypeCommand() CliOption {
	return func(cli *Cli) (err error) {
		cli.BookmarkTypeCmd = &cobra.Command{
			Use:   "type",
			Short: "Manage types of bntp bookmarks",
			Long:  `A longer description`,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				return nil
			},
		}

		cli.BookmarkTypeAddCmd = &cobra.Command{
			Use:   "add TYPE...",
			Short: "Add bntp bookmark types",
			Long:  `A longer description`,
			Args:  cobra.ArbitraryArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				err := cli.BNTPBackend.BookmarkManager.AddType(context.Background(), args)

				return err
			},
		}

		cli.BookmarkTypeEditCmd = &cobra.Command{
			Use:   "edit OLD_NAME NEW_NAME",
			Short: "Change a bntp bookmark type",
			Long:  `A longer description`,
			Args:  cobra.ExactArgs(2),
			RunE: func(cmd *cobra.Command, args []string) error {
				err := cli.BNTPBackend.BookmarkManager.UpdateType(context.Background(), args[0], args[1])

				return err
			},
		}

		cli.BookmarkTypeRemoveCmd = &cobra.Command{
			Use:   "remove TYPE...",
			Short: "Remove bntp bookmark types",
			Long:  `A longer description`,
			Args:  cobra.ArbitraryArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				err := cli.BNTPBackend.BookmarkManager.DeleteType(context.Background(), args)

				return err
			},
		}

		cli.BookmarkTypeListCmd = &cobra.Command{
			Use:   "list",
			Short: "List bntp bookmark types",
			Long:  `A longer description`,
			Args:  cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				types, err := cli.BNTPBackend.BookmarkManager.GetAllTypes(context.Background())

				fmt.Fprintln(cli.RootCmd.OutOrStdout(), strings.Join(types, "\n"))

				return err
			},
		}

		cli.BookmarkCmd.AddCommand(cli.BookmarkTypeCmd)
		cli.BookmarkTypeCmd.AddCommand(cli.BookmarkTypeAddCmd)
		cli.BookmarkTypeCmd.AddCommand(cli.BookmarkTypeEditCmd)
		cli.BookmarkTypeCmd.AddCommand(cli.BookmarkTypeRemoveCmd)
		cli.BookmarkTypeCmd.AddCommand(cli.BookmarkTypeListCmd)

		return
	}
}
