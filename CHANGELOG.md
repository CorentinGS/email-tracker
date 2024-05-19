## 1.0.0-a4 (2024-05-07)

### 💚👷 CI & Build

- **devcontainer**: remove devbox for devcontainer

### 📌➕⬇️ ➖⬆️  Dependencies

- **gomod**: update deps

### 🔥⚰️  Clean up

- remove old assets

### 🚀 Deployments

- **docker**: fix dockerfile

## 1.0.0-a3 (2024-05-04)

### ✨ Features

- **matches**: display the upcoming matches by rounds

### ⚡️ Performance

- improve tournaments
- improve tournaments

### BREAKING CHANGE

- moved main to ./cmd/v1 and changed major versions

### fix

- **deps**: update module github.com/go-playground/validator/v10 to v10.20.0
- **deps**: update module github.com/bytedance/gopkg to v0.0.0-20240419070415-fefc805d4d2a
- **deps**: update module github.com/labstack/echo/v4 to v4.12.0

### 💄🚸 UI & UIX

- **views**: link in tournament list

### 💚👷 CI & Build

- **templ**: build templ files

### 📌➕⬇️ ➖⬆️  Dependencies

- **golang**: update deps

## 1.0.0-a2 (2024-04-15)

### ✨ Features

- **matches**: display the upcoming matches by rounds
- **matches**: match result
- **tournaments**: fill the matches from lichess' tournaments information

### 🐛🚑️ Fixes

- **precommit**: improve pre-commit

### BREAKING CHANGE

- need database migration with breaking tables

### 🎨🏗️ Style & Architecture

- **gci**: beautify the code

### 💄🚸 UI & UIX

- **views**: link in tournament list

### 💚👷 CI & Build

- **golang-ci**: fix golang-ci warnings

### 📝💡 Documentation

- update changelog

### 🔐🚧📈✏️ 💩👽️🍻💬🥚🌱🚩🥅🩺 Others

- **lichess**: fetch tournaments from lichess and store in the database
- **services**: generate DI for services

### 🗃️ Database

- **sqlc**: create queries and add players table

### 🚨 Linting

- fix lint

## 1.0.0-a1 (2024-04-15)

### ✨ Features

- **matches**: match result
- **tournaments**: fill the matches from lichess' tournaments information

### 💚👷 CI & Build

- **golang-ci**: fix golang-ci warnings

### 🔐🚧📈✏️ 💩👽️🍻💬🥚🌱🚩🥅🩺 Others

- **lichess**: fetch tournaments from lichess and store in the database

## 1.0.0-a0 (2024-04-15)

### BREAKING CHANGE

- need database migration with breaking tables

### 🔐🚧📈✏️ 💩👽️🍻💬🥚🌱🚩🥅🩺 Others

- **services**: generate DI for services

### 🗃️ Database

- **sqlc**: create queries and add players table

### 🚨 Linting

- fix lint

## 0.0.2 (2024-04-12)

### 🐛🚑️ Fixes

- **precommit**: improve pre-commit

## 0.0.1 (2024-04-12)

### ✨ Features

- **auth**: jwt keys

### 🎨🏗️ Style & Architecture

- **gci**: beautify the code

### 📌➕⬇️ ➖⬆️  Dependencies

- **go.mod**: update golang deps
