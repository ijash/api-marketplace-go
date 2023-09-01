#!/bin/bash
# https://github.com/zhao-lin-li/postgresql-with-docker-compose

# docker up and assign command variable
alias pgup='docker-compose up -d'

# docker down
alias pgdn='docker-compose down'

# or, use psql client
alias pgconnect='docker-compose run database bash'



