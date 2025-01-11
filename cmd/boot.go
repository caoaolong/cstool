/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cstool/conf"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var config = conf.BootConfig{
	Name: "boot",
}

// bootCmd represents the boot command
var bootCmd = &cobra.Command{
	Use:   config.Name,
	Short: "查看引导扇区编码结构",
	Run:   command,
}

func readFile() []byte {
	// 打开文件
	file, err := os.Open(config.File)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()
	// 创建一个缓冲区
	buffer := make([]byte, 512)
	// 读取前512个字节
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		log.Fatal(err)
		return nil
	}
	return buffer
}

var columns = 16

func drawTable(buffer []byte) *tview.Table {
	table := tview.NewTable().SetBorders(true).SetSelectable(true, true).
		SetSelectedStyle(tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorNone).Bold(true))
	table.SetBackgroundColor(tcell.ColorDefault)
	var sectorSize = 512
	var rows = sectorSize / columns
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			cell := tview.NewTableCell(fmt.Sprintf(" %02X ", buffer[r*columns+c]))
			var idx = r*columns + c
			for k, v := range conf.MBR {
				if idx >= k[0] && idx < k[1] {
					cell.SetTextColor(tcell.NewHexColor(v.Color))
				}
			}
			table.SetCell(r, c, cell)
		}
		var value strings.Builder
		value.WriteByte(' ')
		for i := 0; i < columns; i++ {
			var v = buffer[r*columns+i]
			if v < 127 && v > 31 {
				value.WriteByte(v)
			} else {
				value.WriteByte('.')
			}
		}
		value.WriteByte(' ')
		table.SetCellSimple(r, columns, value.String())
		table.SetCellSimple(r, columns+1, fmt.Sprintf("%04X", r*columns))
	}
	table.SetBordersColor(tcell.ColorBlack)
	return table
}

func command(cmd *cobra.Command, args []string) {
	data := readFile()
	table := drawTable(data)
	app := tview.NewApplication()
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			app.Stop()
		}
		return event
	})
	label := tview.NewTable()
	label.SetBackgroundColor(tcell.ColorDefault).SetBorder(true)
	label.SetCellSimple(0, 0, "Field:")
	label.SetCellSimple(0, 1, "--")
	label.SetCellSimple(1, 0, "Offset:")
	label.SetCellSimple(1, 1, "--")
	label.SetCellSimple(2, 0, "Length:")
	label.SetCellSimple(2, 1, "--")
	table.SetSelectedFunc(func(row, column int) {
		var idx = row*columns + column
		for k, v := range conf.MBR {
			if idx >= k[0] && idx < k[1] {
				label.SetCellSimple(0, 1, v.Desc)
				label.SetCellSimple(1, 1, fmt.Sprintf("%X", k[0]))
				label.SetCellSimple(2, 1, fmt.Sprintf("%d bytes", k[1]-k[0]))
				break
			}
		}
	})
	flex := tview.NewFlex().
		AddItem(label, 0, 1, false).
		AddItem(table, 0, 2, true)
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(bootCmd)
	// Flags
	bootCmd.Flags().StringVarP(&config.Type, "type", "t", "floopy", "磁盘格式 <disk | floppy>")
	bootCmd.Flags().StringVarP(&config.File, "file", "f", "", "文件路径")
}
