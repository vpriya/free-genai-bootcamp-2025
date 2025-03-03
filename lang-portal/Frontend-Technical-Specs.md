# Frontend Technical Specs

## Pages

### Dashboard `/dashboard`

#### Purpose
The purpose of this page is to provide a summary of learning and act as the default page when the user visits the web-app

#### Components
- Last Study Session
    shows last activity used
    shows when last activity used
    summarizes wrong vs correct from last activity
    has a link to the group

- Study Progress
    - total words study eg. 3/124
        - across all study session show the total words studied out of all possible words in our database
    - display a mastery progress eg. 0%

- Quick Stats
    - success rate eg. 80%
    - total study sessions eg. 4
    - total active groups eg. 3
    - study streak eg. 4 days
- Start Studying Button
    - goes to study activities page


We will need following API endpoints to power this page:
- GET /dashboard/last_study_session
- GET /dashboard/study_progress
- GET /dashboard/quick_stats

### Study Activities Index `/study-activities`

#### Purpose
The purpose of this page is to show a collection of study activities with a thumbnail and its name, to either launch or view the study activity.

#### Components
- Study Activity Card
    - show a thumbnail of the study activity
    - the name of the study activity
    - a launch button to take us to the launch page
    - the view page to view more information about past study sessions for this study activity


#### Needed API Endpoints

- GET /study_activities

### Study Activity Show `/study-activity/:id`

#### Purpose
The purpose of this page is to show a detailed view of a study activity and its past study sessions.

#### Components
    
    - Name of the study activity
    - Thumbnail of study activity
    - Description of the study activity
    - Launch button
    - Study activities Paginated list
        - id
        - activity name
        - group name
        - start time
        - end time (inferred by the last word_review_item submitted)

#### Needed API Endpoints
- GET /api/study_activities/:id
- GET /api/study_activities/:id/study_sessions

### Study Activities Launch `/study-activities/:id/launch`
#### Purpose
The purpose of this page is to launch a study activity and start the timer.

#### Components
- Name of the study activity
- Launch Form
    - select field for group
    - launch now button

## Behavior
After the form is submitted a new tab opens with the study activity based on its URL provided in the database.

Also after the form is submitted the page will redirect to the study session show page.

#### Needed API Endpoints   
- POST /api/study_activities

### Words Index `/words`

#### Purpose
The purpose of this page is to show all words in our database.
#### Components
- Paginated Word List
    - Columns
        - Japanese
        - Romaji
        - English meaning
        - Correct Count
        - wrong count
    - Pagination with 100 items per page
    - Clcking the Japanese word will take us to the word show page

            
    
#### Needed API Endpoints
    -  GET /api/words

### Word Show `/words/:id` 
#### Purpose
The purpose of this page is to show information about a specific word.
#### Components
    - Japanese
    - Romaji
    - English
    - Study Statistics
        - Correct Count
        - wrong count
    - Word Groups 
        - show a series pf pills eg. tags
        - when group name is clicked it will take us to the group show page

#### Needed API Endpoints
- GET /api/words/:id
    
### Word Groups Index `/groups`
#### Purpose
The purpose of this page is to show a list of groups in our database.

#### Components
    - Paginated Group List
        -Columns
            - Group Name
            - Word Count
        - clicking the group name will take us to the group show page

#### Needed API Endpoints
- GET /api/groups

### Group Show `/groups/:id`
#### Purpose
The purpose of this page is to show information about a specific group.

#### Components 
- Group Name
- Group Statistics
    - Total Word Count
- Words in Group (Paginated List of Words)
    - Should use the same component as the words index page
- Study Sessions (Paginated List of Study Sessions)
    - Should use the same component as the study sessions index page

#### Needed API Endpoints
- GET /api/groups/:id (the name and groups stats)
- GET /api/groups/:id/words
- GET /api/groups/:id/study_sessions

## Study Sessions Index `/study-sessions`
#### Purpose
The purpose of this page is to show a list of study sessions in our database.

#### Components
- Paginated Study Session List
    - Columns
        - Id
        - Activity Name
        - Group Name
        - Start Time
        - End Time
        - Number of Review Items
    - Clicking on the study session will take us to the study session show page

#### Needed API Endpoints
- GET /api/study_sessions

### Study Session Show `/study-sessions/:id`
#### Purpose
The purpose of this page is to show information about a specific study session.

#### Components
- Study Session Details
    - Activity Name
    - Group Name
    - Start Time
    - End Time
    - Number of Review Items
- Words Review Items (Paginated List of Words)
    - Should use the same component as the words index page

#### Needed API Endpoints
- GET /api/study_sessions/:id
- GET /api/study_sessions/:id/words

### Settings Page `/settings`
#### Purpose
The purpose of this page is to make configurations to the study portal.

#### Components
- Theme Selection eg. Light, Dark, System Default
- Reset History Button
    - this will delete all the study sessions and word review items
- Full Reset Button
    - this will drop all tables and re-create with the seed data

#### Needed API Endpoints
- POST /api/reset_history
- POST /api/full_reset
















keeping following for copy and paste
#### Purpose
#### Components
#### Needed API Endpoints
