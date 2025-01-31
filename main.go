package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/marcuschui2022/blog-aggregator/internal/config"
	"github.com/marcuschui2022/blog-aggregator/internal/database"
	"html"
	"io"
	"log"
	"net/http"
	"os"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	//init cfg
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error connecting db: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	dbQueries := database.New(db)

	appState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	//init cmds
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)

	//get cli Args
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	//init cmd
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	//execute cmd
	err = cmds.run(appState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(
		"User-Agent",
		"gator")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			_ = fmt.Errorf("error closing body: %v", err)
		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var feed RSSFeed
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := 0; i < len(feed.Channel.Item); i++ {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	return &feed, nil
}

func handlerAgg(s *state, cmd command) error {
	//_ = s
	//if len(cmd.Args) != 1 {
	//	return fmt.Errorf("usage: %v <url>", cmd.Name)
	//
	//}
	//_ = cmd.Args[0]

	ctx := context.Background()
	feed, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}
	fmt.Println(feed)

	return nil
}
