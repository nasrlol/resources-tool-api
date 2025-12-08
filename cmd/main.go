package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
	"os"

	router "synf/internal/api/routes"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pressly/goose/v3"
)

var (
	embedMigrations embed.FS
	PORT            string
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func commands() model {
	return model{
		choices:  []string{"start", "migrate", "reset-migrations"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	PORT = ":8500"

	r := router.InitRestRoutes()
	log.Fatal(http.ListenAndServe(PORT, r))

	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Start the application\n\n"

	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

func migrate() {
	var db *sql.DB

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}

// Reset the database
func resetMigrations() {
	var db *sql.DB

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	if err := goose.Down(db, "migrations"); err != nil {
		panic(err)
	}
}

func RawConnect(host string, port string) {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		fmt.Println("Connecting error:", err)
	}
	if conn != nil {
		defer func(conn net.Conn) {
			_ = conn.Close()
		}(conn)
		fmt.Println("Opened", net.JoinHostPort(host, port))
	}
}

func main() {
		p := tea.NewProgram(commands())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}

}
