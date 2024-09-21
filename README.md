# crgenerator

<a href="#">
 <img src="https://img.shields.io/badge/license-MIT-blue.svg">
</a>
<a href="?tab=readme-ov-file#contribution">
 <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square">
</a>
<br><br>
A tool to analyze git commits, extract JIRA numbers and pull issue details for Change Request process

## Setup ##
1. install
```
    brew install crgenerator
```
2. config
```
    echo 'export JIRA_BASE_URL=???' >> ~/.zshrc
    echo 'export JIRA_USER_NAME=???' >> ~/.zshrc
    echo 'export JIRA_API_TOKEN=???' >> ~/.zshrc
    source ~/.zshrc
```
3. run (from repository directory, and if no end commit is provided it will default to the lastest one)
```
    crgenerator $START_COMMIT $END_COMMIT
```

sample output (with JIRA issue number, description and URL)
![output](https://github.com/user-attachments/assets/92ec4ca9-3757-4378-a84c-5c48943a915e)



## Contribution ##
Your contributions are always welcome!

## License ##
This work is licensed under [MIT](https://opensource.org/licenses/MIT).
