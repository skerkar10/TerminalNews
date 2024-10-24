package main

import (
  //GOLANG and local
  "fmt"
  "os"
  "TerminalNews/ScrapePage"
  "TerminalNews/OpenURL"

  //Web Scraping
  "github.com/gocolly/colly"
  "github.com/gocolly/colly/extensions"


  //TUI
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/bubbles/table"
  "github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
  BorderStyle(lipgloss.NormalBorder()).
  BorderForeground(lipgloss.Color("240"))

type model struct {
  table table.Model
  stories []ScrapePage.Story
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "q", "ctrl+c":
      return m, tea.Quit
    case "enter":
      cursorVal := m.table.Cursor()
      urlVal := m.stories[cursorVal].Url
      return m, func() tea.Msg {
        OpenURL.OpenLink(urlVal)
        return nil
      }
    }
  }
  m.table, cmd = m.table.Update(msg)
  return m, cmd
}

func (m model) View() string {
  return baseStyle.Render(m.table.View()) + "\n"
}

func main() {
  webScraper := colly.NewCollector(
    colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
    colly.Async(true),
    )

  extensions.RandomUserAgent(webScraper)
  extensions.Referer(webScraper)

  stories := ScrapePage.ScrapeNews(webScraper)


  columns := []table.Column{
    {Title: "Title", Width: 200},
  }

  rows := []table.Row{}

  for _, story := range stories {
    rows = append(rows, table.Row{story.Name})
  }

  t := table.New(
    table.WithColumns(columns),
    table.WithRows(rows),
    table.WithFocused(true),
    table.WithHeight(7),
    )

  s := table.DefaultStyles()
  s.Header = s.Header.
    BorderStyle(lipgloss.NormalBorder()).
    BorderForeground(lipgloss.Color("240")).
    BorderBottom(true).
    Bold(false)
  s.Selected = s.Selected.
    Foreground(lipgloss.Color("229")).
    Background(lipgloss.Color("57")).
    Bold(false)
  t.SetStyles(s)

  m := model{
    table: t,
    stories: stories,
  }

  if _, err := tea.NewProgram(m).Run(); err != nil {
    fmt.Println("Error running program:", err)
    os.Exit(1)
  }
}
