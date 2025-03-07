# Backend Server Technical Specs

## Business Goal: 
A language learning school wants to build a prototype of learning portal which will act as three things:
- Inventory of possible vocabulary that can be learned
- Act as a  Learning record store (LRS), providing correct and wrong score on practice vocabulary
- A unified launchpad to launch different learning apps

## Technical Requirements

- The backend will be built using Go
- The database will be SQLite3
- The API will be built using Gin
- The API will always return JSON
- There will be no authentication or authorization
- Everything will be treated as a single user

## Database Schema
Our databse will be a single sqlite3 database called `words.db` that will be in the root of the project folder of `backend_go`

We have the following tables:
-   words - stored vocabulary words
    - id integer primary key

    - japanese string
    - romaji string
    - english string
    - parts json

-   words_groups - join table for words and groups many to many
    - id integer
    - word_id integer
    - group_id integer

-   groups - thematic groups of words
    - id integer
    - name string

-   study_sessions - records of study sessions grouping word_review_items
    - id integer
    - group_id integer
    - study_activity_id integer
    - correct boolean
    - created_at datetime

-   study_activities - a specific study activity, linking a study_session to group
    - id integer
    - study_session_id integer
    - group_id integer
    - created_at datetime

-   word_review_items - a record of word practice, determining if the word was       correct or not
    - word_id integer
    - study_session_id integer
    - correct boolean
    - created_at datetime

### API Endpoints

### GET /api/dashboard/last_study_session
Returns information about the most recent study session.
#### JSON Response
```json
{
    "id": 1,
    "group_id": 1,
    "study_activity_id": 1,
    "created_at": "2024-03-20T15:04:05Z",
    "group_name": "Basic Greetings"
}
```

### GET /api/dashboard/study_progress
Returns study progress statistics.
Please note that the frontend will determine progress bar based on total words studied and total available words.
#### JSON Response
```json
{
    "total_words_studied": 3,
    "total_words_studied": 124
}
```

### GET /api/dashboard/quick_stats
Returns quick overview statistics.
#### JSON Response
```json
{
    "success_rate": 0.80,
    "total_study_sessions": 4,
    "total_active_groups": 3,
    "streak_days": 4
}
```

### GET /api/study_activities/:id
Returns information about a specific study activity.
#### JSON Response
```json
{
    "id": 1,
    "name": "Flashcards",
    "thumbnail_url": "https://example.com/flashcards.png",
    "description": "Practice vocabulary with flashcards",
    "launch_url": "https://example.com/launch/flashcards"
}
```

### GET /api/study_activities/:id/study_sessions
#### JSON Response
```json
{
    "items": [
        {
            "id": 1,
            "activity_name": "Flashcards",
            "group_name": "Basic Greetings",
            "start_time": "2024-03-20T15:04:05Z",
            "end_time": "2024-03-20T15:14:05Z",
            "review_items_count": 20
        }
    ],
    "pagination": {
        "current_page": 1,
        "total_pages": 5,
        "total_items": 100,
        "items_per_page": 20
    }
}
```

### POST /api/study_activities
#### Request Params: 
- group_id integer, 
- study_activity_id integer   
### JSON Response:
```json
{
    "id": 1,
    "group_id": 1
}
```

### GET /api/words
    - pagination with 100 items per page
#### JSON Response
```json
{
    "items": [
        {
            "id": 1,
            "japanese": "こんにちは",
            "romaji": "konnichiwa",
            "english": "hello",
            "correct_count": 10,
            "wrong_count": 2
        }
    ],
    "pagination": {
        "current_page": 1,
        "total_pages": 10,
        "total_items": 1000,
        "items_per_page": 100
    }
}
```



### GET /api/words/:id
#### JSON Response
```json
{
    "id": 1,
    "japanese": "こんにちは",
    "romaji": "konnichiwa",
    "english": "hello",
    "correct_count": 10,
    "wrong_count": 2,
    "groups": [
        {
            "id": 1,
            "name": "Basic Greetings"
        }
    ]
}
```

### GET /api/groups
    - pagination with 100 items per page
#### JSON Response
```json
{
    "items": [
        {
            "id": 1,
            "name": "Basic Greetings",
            "word_count": 20
        }
    ],
    "pagination": {
        "current_page": 1,
        "total_pages": 5,
        "total_items": 10,
        "items_per_page": 100
    }
}
```

### GET /api/groups/:id
#### JSON Response
```json
{
    "id": 1,
    "name": "Basic Greetings",
    "stats": {
        "total_word_count": 20
    }
}
```

### GET /api/groups/:id/words
#### JSON Response
```json
{
    "items": [
        {
            "id": 1,
            "japanese": "こんにちは",
            "romaji": "konnichiwa",
            "english": "hello",
            "correct_count": 10,
            "wrong_count": 2
        }
    ],
    "pagination": {
        "current_page": 1,
        "total_pages": 10,
        "total_items": 20,
        "items_per_page": 100
    }
}
```

### GET /api/groups/:id/study_sessions
#### JSON Response
```json
{
    "items": [
        {
            "id": 1,
            "activity_name": "Flashcards",
            "start_time": "2024-03-20T15:04:05Z",
            "end_time": "2024-03-20T15:14:05Z",
            "review_items_count": 20
        }
    ],
    "pagination": {
        "current_page": 1,
        "total_pages": 5,
        "total_items": 50
    }
}
```

### GET /api/study_sessions
    - pagination with 100 items per page
#### JSON Response
```json
{
    "items": [
        {
            "id": 1,
            "activity_name": "Flashcards",
            "group_name": "Basic Greetings",
            "start_time": "2024-03-20T15:04:05Z",
            "end_time": "2024-03-20T15:14:05Z",
            "review_items_count": 20
        }
    ],
    "pagination": {
        "current_page": 1,
        "total_pages": 10,
        "total_items": 1000,
        "items_per_page": 100
    }
}
```

### GET /api/study_sessions/:id
#### JSON Response
```json
{
    "id": 1,
    "activity_name": "Flashcards",
    "group_name": "Basic Greetings",
    "start_time": "2024-03-20T15:04:05Z",
    "end_time": "2024-03-20T15:14:05Z",
    "review_items_count": 20
}
```

### GET /api/study_sessions/:id/words
#### JSON Response
```json
{
    "items": [
        {
            "id": 1,
            "japanese": "こんにちは",
            "romaji": "konnichiwa",
            "english": "hello",
            "correct_count": 10,
            "wrong_count": 2
        }
    ],
    "pagination": {
        "current_page": 1,
        "total_pages": 10,
        "total_items": 20,
        "items_per_page": 100
    }
}
```

### POST /api/reset_history
#### JSON Response
```json
{
    "success": true,
    "message": "Study history has been reset successfully"
}
```

### POST /api/full_reset
#### JSON Response
```json
{
    "success": true,
    "message": "All data has been reset successfully"
}
```

### POST /api/study_sessions/:study_session_id/words/:word_id/review
#### Request Params: 
- id (study_session_id) integer
- word_id integer
- correct boolean
#### Request Payload
```json
{
    "correct": true
}
```
#### JSON Response
```json
{
    "success": true,
    "word_id": 1,
    "study_session_id": 123,
    "correct": true,
    "created_at": "2024-03-20T15:04:05Z"
}
```
## Mage Tasks
Mage is a task runner for Go.
Lets list out possible tasks we need for our lang portal.

### Initialize Database
This task will initialize the sqllite databse called `words.db`

### Migrate Database
This task will run a series of migrations sql files on the database

Migrations live in the `migrations` folder.
The migration files will be run in order of their file name.
The file names should look like this:

```sql
0001_init.sql
0002_create_words_table.sql
0003_create_groups_table.sql
```

### Seed Data
This task will import json files and transform them into target data for our database.

All seed files live in the `seeds` folder.
All seed files should be loaded.

```json
[
  {
    "kanji": "払う",
    "romaji": "harau",
    "english": "to pay",
  },
]
```







