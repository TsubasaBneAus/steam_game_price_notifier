// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/TsubasaBneAus/steam_game_price_notifier/app/external/discord"
	"github.com/TsubasaBneAus/steam_game_price_notifier/app/external/httpclient"
	"github.com/TsubasaBneAus/steam_game_price_notifier/app/external/notion"
	"github.com/TsubasaBneAus/steam_game_price_notifier/app/external/steam"
	"github.com/TsubasaBneAus/steam_game_price_notifier/app/interactor"
	"github.com/TsubasaBneAus/steam_game_price_notifier/config"
	"github.com/google/wire"
)

// Injectors from wire.go:

// Initialize the application
func InitializeApp(ctx context.Context) (*app, error) {
	notionConfig, err := config.NewNotionConfig(ctx)
	if err != nil {
		return nil, err
	}
	steamConfig, err := config.NewSteamConfig(ctx)
	if err != nil {
		return nil, err
	}
	httpClient := httpclient.NewHTTPClient()
	steamWishlistGetter := steam.NewSteamWishlistGetter(steamConfig, httpClient)
	steamVideoGameDetailsGetter := steam.NewSteamVideoGameDetailsGetter(steamConfig, httpClient)
	notionWishlistGetter := notion.NewNotionWishlistGetter(notionConfig, httpClient)
	notionWishlistItemCreator := notion.NewNotionWishlistItemCreator(notionConfig, httpClient)
	notionWishlistItemUpdater := notion.NewNotionWishlistItemUpdater(notionConfig, httpClient)
	discordConfig, err := config.NewDiscordConfig(ctx)
	if err != nil {
		return nil, err
	}
	videoGamePricesOnDiscordNotifier := discord.NewVideoGamePricesOnDiscordNotifier(discordConfig, httpClient)
	videoGamePricesNotifier := interactor.NewGamePricesNotifier(notionConfig, steamWishlistGetter, steamVideoGameDetailsGetter, notionWishlistGetter, notionWishlistItemCreator, notionWishlistItemUpdater, videoGamePricesOnDiscordNotifier)
	errorOnDiscordNotifier := discord.NewErrorOnDiscordNotifier(discordConfig, httpClient)
	interactorErrorOnDiscordNotifier := interactor.NewErrorOnDiscordNotifier(discordConfig, errorOnDiscordNotifier)
	mainApp := NewApp(videoGamePricesNotifier, interactorErrorOnDiscordNotifier)
	return mainApp, nil
}

// wire.go:

// A wire set for the main package
var Set = wire.NewSet(
	NewApp, config.Set, httpclient.Set, steam.Set, notion.Set, discord.Set, interactor.Set,
)
