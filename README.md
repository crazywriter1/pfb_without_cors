# pfb_without_cors
### Make sure you have git on your server

# Step by Step Intro

## 1. Step (Clone)

`git clone https://github.com/crazywriter1/pfb_without_cors `

## 2. Step (allow Port)

'sudo apt install ufw -y'
'sudo ufw allow 8080'
'sudo ufw enable'

# 3. Check Port Status

`sudo ufw status`

# 4. Change Directory (cd)

`cd pfb_without_cors`

## 5. Step (Install Packages)

`go mod tidy`

## 6. Step (Run)

`go run main.go`

# Note
### Do not close the program if you are using the website.
