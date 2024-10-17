# Project Overview for Concierge

## Project Structure
```
Concierge/
├── docker-compose/
|   └──configs/
├── docs/
|    └──images/
|   └──sprint_0/
|   └──sprint_1/
└── src/
    ├── accounts/
    ├── amenities/
    ├── api_help/
    ├── guest_webapp
    ├── incident_reports/     
    ├── permissions
    ├── sessions
    ├── staff_webapp
    └── webapp
```
## Branching Strategy
Before you branch, create an issue for what you are working on.
We'll use the issue # in the branch name and snake_case.

the format is:
```
{branch-type}/{feature-name}/{issue-#}
ex. feature/account_server/9354093
```

| Branch Type | Description                                                          |
| ---------- | ----------------------------------------------------------------     |
| `feature`  | use this when changes are related to a development of a feature.     |
| `dev_task` | use this when changes are related to a specified dev task            |
| `docs`     | use this when changes are documentation related                      |
| `refactor` | use this when refactor the code base                                 |

## We Have Style

We will use the following style guides on these following languages:
- [Python](https://google.github.io/styleguide/pyguide.html)
- [C#](https://google.github.io/styleguide/csharp-style.html)
- [JavaScript](https://google.github.io/styleguide/jsguide.html)
- [TypeScript](https://google.github.io/styleguide/tsguide.html)
