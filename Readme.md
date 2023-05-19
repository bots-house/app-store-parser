# App store parser

Clone of the [app-store-parser](https://github.com/facundoolano/app-store-scraper)

- [Quick example](#quick-example)
- [Features](#features)
  - [App](#app)
  - [Similar](#similar)
  - [List](#list)
  - [Search](#search)
  - [Developer](#developer)
  - [Ratings](#ratings)
  - [Reviews](#reviews)
  - [Privacy](#privacy)
  - [Suggest](#suggest)

# Quick example

```go
package main

import (
    "context"
    "log"

    asp "github.com/bots-house/app-store-parser"
)

func main() {
    collector := asp.New()

    app, err := collector.App(context.Background(), asp.AppSpec{
        AppID: "com.netflix.Netflix",
    })
    if err != nil {
        log.Fatal(err)
    }

    log.Println(app)
}
```

# Features

## App

Method for parsing app data from app store

**Parameters**

- `id` - app id such **553834731**
- `app-id` - platform readable id such **com.netflix.Netflix**
- `lang`
- `country` in ISO format
- `ratings` - if `true` will return total ratings of app and histogram

**id or app_id required**

**Example result**

```json
{
    "id": 363590051,
    "app_id": "com.netflix.Netflix",
    "title": "Netflix",
    "url": "https://apps.apple.com/us/app/netflix/id363590051?uo=4",
    "description": "~~app description~~",
    "genres": [
    "Entertainment",
    "Lifestyle"
    ],
    "genre_ids": [
        "6016",
        "6012"
    ],
    "primary_genre": "Entertainment",
    "primary_genre_id": 6016,
    "content_rating": "12+",
    "size": "156941312",
    "required_os_version": "15.0",
    "released": "2010-04-01T20:41:34Z",
    "updated": "2023-05-15T16:00:32Z",
    "release_notes": "Keep your app updated to get the latest Netflix experience on your iPhone and iPad.\n\nAnd in this release, we fixed bugs and made performance improvements. Just for you.",
    "version": "15.31.0",
    "currency": "USD",
    "free": true,
    "developer_id": 363590054,
    "developer": "Netflix, Inc.",
    "developer_url": "https://apps.apple.com/us/developer/netflix-inc/id363590054?uo=4",
    "developer_website": "http://www.netflix.com",
    "score": 3.71896,
    "reviews": 348265,
    "current_version_score": 3.71896,
    "current_version_reviews": 348265,
    "icon": "https://is3-ssl.mzstatic.com/image/thumb/Purple116/v4/87/87/7c/87877ccb-957f-662c-7963-b1b6dea3ab5e/AppIcon-0-0-1x_U007emarketing-0-0-0-10-0-0-0-85-220.png/512x512bb.jpg",
    ...
}
```

---

## Similar

Method which work as app but return apps that similar for requested id or app_id


**Parameters**

- `id` - app id such **553834731**
- `app-id` - platform readable id such **com.netflix.Netflix**
- `lang`
- `country` in ISO format
- `ratings` - if `true` will return total ratings of app and histogram

**id or app_id required**

---

## List

Method which return list of apps

**Parameters**

- `count`
- `country` in ISO format
- `collection` - possible [values](#list-collections)
- `category` - possible [values](#list-categories)

---

## Search

Method for searching with query. Return apps matched search result.

**Parameters**

- `query` - search query [required]
- `lang`
- `country` in ISO format
- `count`
- `page`
- `ids-only` - if `true` will return only matched numeric app ids

---

## Developer

Method which return apps for requested developer.

**Parameters**

- `id` - app id such **297606954** [required]
- `lang`
- `country` in ISO format

---

## Ratings

Method which parse app ratings

**Parameters**

- `id` - app id such **553834731** [required]
- `country` in ISO format

**Example result**

```json
{
  "total": 2864240,
  "histogram": {
    "1": 51249,
    "2": 30639,
    "3": 98589,
    "4": 343333,
    "5": 2340430
  }
}
```

---

## Reviews

Method return app reviews

**Parameters**

- `id` - app id such **553834731** 
- `app_id`
- `country` in ISO format
- `page`
- `sort` - possible [values](#reviews-sort)


**id or app_id required**

**Example result**

```json
[
  {
    "id": "9939621285",
    "title": "What is wrong with the game",
    "content": "Please fix the bug since yesterday it always mess up my game everytime I click the add it turn out black screen and everything is gone so i have to refresh and my progress on game is gone and I have to start over and lose a life.",
    "user_name": "nyllen",
    "user_url": "https://itunes.apple.com/us/reviews/id392839919",
    "version": "1.252.1.1",
    "score": "3",
    "url": "https://itunes.apple.com/us/review?id=553834731&type=Purple%20Software",
    "updated": "2023-05-18T05:45:37-07:00"
  },
  {
    "id": "9938788139",
    "title": "Freezing up",
    "content": "Not sure if I need an update or what because it keeps freezing or stalling on me when I lose to get the extra turns and when I get a prize before the game what’s going on my phone has been updated but 😢",
    "user_name": "convincedsofty2",
    "user_url": "https://itunes.apple.com/us/reviews/id1310134879",
    "version": "1.252.1.1",
    "score": "1",
    "url": "https://itunes.apple.com/us/review?id=553834731&type=Purple%20Software",
    "updated": "2023-05-18T00:06:37-07:00"
  },
  ...
]
```

---

## Privacy

Method which parse app privacy settings
**Parameters**

- `id` - app id such **553834731** [required]

**Example result**

```json
[
  {
    "categories": [
      {
        "category": "Purchases",
        "types": [
          "Purchase History"
        ],
        "identifier": "PURCHASES"
      },
      {
        "category": "Location",
        "types": [
          "Coarse Location"
        ],
        "identifier": "LOCATION"
      },
      {
        "category": "Contact Info",
        "types": [
          "Email Address",
          "Name"
        ],
        "identifier": "CONTACT_INFO"
      },
      {
        "category": "User Content",
        "types": [
          "Gameplay Content"
        ],
        "identifier": "USER_CONTENT"
      },
      {
        "category": "Identifiers",
        "types": [
          "User ID",
          "Device ID"
        ],
        "identifier": "IDENTIFIERS"
      },
      {
        "category": "Usage Data",
        "types": [
          "Product Interaction",
          "Advertising Data",
          "Other Usage Data"
        ],
        "identifier": "USAGE_DATA"
      }
    ],
    "description": "The following data may be used to track you across apps and websites owned by other companies:",
    "identifier": "DATA_USED_TO_TRACK_YOU",
    "type": "Data Used to Track You",
    "purposes": []
  },
  ...
]
```

---

## Suggest

Method which takes search query and return app variants

**Parameters**

- `query` [required]
- `country` in ISO format

**Example result**

```json
[
  {
    "term": "vrbo",
    "url": "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search?clientApplication=Software&media=software&src=hint&submit=edit&term=vrbo"
  },
  {
    "term": "vr games",
    "url": "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search?clientApplication=Software&media=software&src=hint&submit=edit&term=vr%20games"
  },
  {
    "term": "vr",
    "url": "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search?clientApplication=Software&media=software&src=hint&submit=edit&term=vr"
  },
  {
    "term": "vrv",
    "url": "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search?clientApplication=Software&media=software&src=hint&submit=edit&term=vrv"
  },
  {
    "term": "vroom",
    "url": "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search?clientApplication=Software&media=software&src=hint&submit=edit&term=vroom"
  },
  {
    "term": "vrchat",
    "url": "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search?clientApplication=Software&media=software&src=hint&submit=edit&term=vrchat"
  },
  {
    "term": "vrbo owner",
    "url": "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search?clientApplication=Software&media=software&src=hint&submit=edit&term=vrbo%20owner"
  },
  {
    "term": "vr apps for iphone",
    "url": "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search?clientApplication=Software&media=software&src=hint&submit=edit&term=vr%20apps%20for%20iphone"
  },
  {
    "term": "vrv: anime, game videos & more",
    "url": "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search?clientApplication=Software&media=software&src=hint&submit=edit&term=vrv%3A%20anime%2C%20game%20videos%20%26%20more"
  },
  {
    "term": "vrbo.com",
    "url": "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search?clientApplication=Software&media=software&src=hint&submit=edit&term=vrbo.com"
  }
]
```
Each link redirect to app store.

---

---

## List collections

- "TOP_MAC"
- "TOP_FREE_MAC"
- "TOP_GROSSING_MAC"
- "TOP_PAID_MAC"
- "NEW_IOS"
- "NEW_FREE_IOS"
- "NEW_PAID_IOS"
- "TOP_FREE_IOS"
- "TOP_FREE_IPAD"
- "TOP_GROSSING_IOS"
- "TOP_GROSSING_IPAD"
- "TOP_PAID_IOS"
- "TOP_PAID_IPAD"

## List categories

- "BOOKS"
- "BUSINESS"
- "CATALOGS"
- "EDUCATION"
- "ENTERTAINMENT"
- "FINANCE"
- "FOOD_AND_DRINK"
- "GAMES"
- "GAMES_ACTION"
- "GAMES_ADVENTURE"
- "GAMES_ARCADE"
- "GAMES_BOARD"
- "GAMES_CARD"
- "GAMES_CASINO"
- "GAMES_DICE"
- "GAMES_EDUCATIONAL"
- "GAMES_FAMILY"
- "GAMES_MUSIC"
- "GAMES_PUZZLE"
- "GAMES_RACING"
- "GAMES_ROLE_PLAYING"
- "GAMES_SIMULATION"
- "GAMES_SPORTS"
- "GAMES_STRATEGY"
- "GAMES_TRIVIA"
- "GAMES_WORD"
- "HEALTH_AND_FITNESS"
- "LIFESTYLE"
- "MAGAZINES_AND_NEWSPAPERS"
- "MAGAZINES_ARTS"
- "MAGAZINES_AUTOMOTIVE"
- "MAGAZINES_WEDDINGS"
- "MAGAZINES_BUSINESS"
- "MAGAZINES_CHILDREN"
- "MAGAZINES_COMPUTER"
- "MAGAZINES_FOOD"
- "MAGAZINES_CRAFTS"
- "MAGAZINES_ELECTRONICS"
- "MAGAZINES_ENTERTAINMENT"
- "MAGAZINES_FASHION"
- "MAGAZINES_HEALTH"
- "MAGAZINES_HISTORY"
- "MAGAZINES_HOME"
- "MAGAZINES_LITERARY"
- "MAGAZINES_MEN"
- "MAGAZINES_MOVIES_AND_MUSIC"
- "MAGAZINES_POLITICS"
- "MAGAZINES_OUTDOORS"
- "MAGAZINES_FAMILY"
- "MAGAZINES_PETS"
- "MAGAZINES_PROFESSIONAL"
- "MAGAZINES_REGIONAL"
- "MAGAZINES_SCIENCE"
- "MAGAZINES_SPORTS"
- "MAGAZINES_TEENS"
- "MAGAZINES_TRAVEL"
- "MAGAZINES_WOMEN"
- "MEDICAL"
- "MUSIC"
- "NAVIGATION"
- "NEWS"
- "PHOTO_AND_VIDEO"
- "PRODUCTIVITY"
- "REFERENCE"
- "SHOPPING"
- "SOCIAL_NETWORKING"
- "SPORTS"
- "TRAVEL"
- "UTILITIES"
- "WEATHER"

## Reviews sort
- "RECENT"
- "HELPFUL"