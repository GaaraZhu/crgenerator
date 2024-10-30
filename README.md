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
3. run - go to project directory, check JIRA issues between commits, tags, branches etc
```
    crgenerator $START $END
```
![issues_between_commits](https://github.com/user-attachments/assets/7914f4e2-936d-4148-9a0b-bdd38b78643b)



## Contribution ##
Your contributions are always welcome!

## License ##
This work is licensed under [MIT](https://opensource.org/licenses/MIT).
