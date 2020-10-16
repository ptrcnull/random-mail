# random-mail

Note: this project doesn't provide mailboxes; its only purpose is to generate semi-legit-looking emails

## Files needed

- [surnames_male.csv](https://dane.gov.pl/pl/dataset/568,nazwiska-wystepujace-w-rejestrze-pesel)
- [surnames_female.csv](https://dane.gov.pl/pl/dataset/568,nazwiska-wystepujace-w-rejestrze-pesel)
- [names_male.csv](https://dane.gov.pl/pl/dataset/1667,lista-imion-wystepujacych-w-rejestrze-pesel-osoby-zyjace)
- [names_female.csv](https://dane.gov.pl/pl/dataset/1667,lista-imion-wystepujacych-w-rejestrze-pesel-osoby-zyjace)

## Using

- `go run main.go` or `docker-compose up -d`
- `GET /`
- `GET /?count=10`
- `GET /?count=10&gender=male`
