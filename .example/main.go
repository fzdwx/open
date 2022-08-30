package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	for i := 0; i < 255; i++ {
		s := lipgloss.NewStyle().Foreground(lipgloss.Color(fmt.Sprintf("%d", i))).Render(fmt.Sprintf("%d hello world", i))
		fmt.Println(s)

	}
}
