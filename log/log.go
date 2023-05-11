package log

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	tu "github.com/thebromo/turing-machine/machine"
)

const columnWidth = 25

var (
	norm      = lipgloss.AdaptiveColor{Light: "#d4d4d5", Dark: "#d4d4d5"}
	subtle    = lipgloss.AdaptiveColor{Light: "#46484c", Dark: "#46484c"}
	highlight = lipgloss.AdaptiveColor{Light: "#e87979", Dark: "#e87979"}
	special   = lipgloss.AdaptiveColor{Light: "#37d99e", Dark: "#37d99e"}

	// List.

	list = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(subtle).
		MarginRight(2).
		Height(8).
		Width(columnWidth + 1)

	listHeader = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(subtle).
			MarginRight(2).
			Width(columnWidth).
			Render

	listHeaderActive = lipgloss.NewStyle().
				BorderStyle(lipgloss.ThickBorder()).
				BorderBottom(true).
				BorderForeground(special).
				MarginRight(2).
				Width(columnWidth).
				Foreground(special).
				Render

	listItem = lipgloss.NewStyle().PaddingLeft(2).Render

	checkMark = lipgloss.NewStyle().SetString("❯").
			Foreground(special).
			PaddingRight(1).
			String()

	listDone = func(s string) string {
		return checkMark + lipgloss.NewStyle().
			Foreground(highlight).Bold(true).
			Render(s)
	}

	machineStyle = lipgloss.NewStyle().Bold(true).
			Foreground(norm).
			PaddingTop(1).
			PaddingBottom(1).
			PaddingLeft(2).
			Margin(1).
			PaddingRight(2).BorderStyle(lipgloss.RoundedBorder()).Render
	machineStyleActive = lipgloss.NewStyle().Foreground(special).Render

	errorStyle = lipgloss.NewStyle().Bold(true).
			Foreground(subtle).Background(highlight).Render
)

func PrintMachine(machine tu.Machine, err error) {
	states := make([]string, 0)
	count := 0
	for _, v := range machine.States {
		if v.Number == machine.CurrentState.Number {
			states = append(states, printActiveState(*v, machine.Tape.Read()))
		} else {
			states = append(states, printState(*v))
		}
		count++
	}

	states = fitStates(states)

	if err == nil {
		fmt.Println(lipgloss.JoinVertical(lipgloss.Top, printTape(*machine.Tape),
			lipgloss.JoinVertical(lipgloss.Top, states...)))
	} else {
		fmt.Println(
			lipgloss.JoinVertical(lipgloss.Left, errorStyle("⚠: "+err.Error()),
				lipgloss.JoinVertical(lipgloss.Top, printTape(*machine.Tape),
					lipgloss.JoinVertical(lipgloss.Top, states...))))
	}

}

func fitStates(states []string) []string {
	count := 0
	current := make([]string, 0)
	temp := make([]string, 0)
	for _, state := range states {
		if count == 5 {
			current = append(current, lipgloss.JoinHorizontal(lipgloss.Center, temp...))
			temp = make([]string, 0)
			count = 0
		} else {
			temp = append(temp, state)
			count++
		}
	}
	return current
}

func printTape(tape tu.Tape) string {
	top := ""
	arrow := ""

	for i, v := range tape.Content {
		if i == tape.Position {
			top += machineStyleActive(" " + strconv.Itoa(v))
			arrow += machineStyleActive(" ▲")
		} else {
			top += " " + strconv.Itoa(v)
			arrow += "  "
		}

	}
	dec := printDecValue(tape.Content)
	return machineStyle(top + "\n" + arrow + "\n d: " + dec + "\n")
}

func printDecValue(tape []int) string {
	values := make([]int, 0)
	count := 0

	for _, v := range tape {
		if v == 0 {
			count++
		} else {
			values = append(values, count)
			count = 0
		}
	}
	values = append(values, count)
	result := ""
	for i, v := range values {
		if i != len(values)-1 {
			result += strconv.Itoa(v) + " , "
		} else {
			result += strconv.Itoa(v)
		}
	}

	return result
}

func printState(state tu.State) string {
	text := make([]string, 0)
	current := strconv.Itoa(state.Number)
	text = append(text, listHeader("q"+current))

	for _, v := range state.Transitions {
		read := strconv.Itoa(v.Read)
		end := strconv.Itoa(v.EndState.Number)
		write := strconv.Itoa(v.Write)
		dir := ""
		if v.Direction == tu.Left {
			dir += "L"
		} else {
			dir += "R"
		}
		text = append(text, listItem("δ(q"+current+", "+read+") = (q"+end+","+write+", "+dir+"),"))
	}

	return list.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			text...,
		),
	)
}
func printActiveState(state tu.State, input int) string {
	text := make([]string, 0)
	current := strconv.Itoa(state.Number)
	text = append(text, listHeaderActive("q"+current))

	for _, v := range state.Transitions {
		read := strconv.Itoa(v.Read)
		end := strconv.Itoa(v.EndState.Number)
		write := strconv.Itoa(v.Write)
		dir := ""
		if v.Direction == tu.Left {
			dir += "L"
		} else {
			dir += "R"
		}

		if input == v.Read {
			text = append(text, listDone("δ(q"+current+", "+read+") = (q"+end+","+write+", "+dir+"),"))

		} else {
			text = append(text, listItem("δ(q"+current+", "+read+") = (q"+end+","+write+", "+dir+"),"))
		}
	}

	return list.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			text...,
		),
	)
}
