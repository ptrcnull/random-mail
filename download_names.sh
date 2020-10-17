#!/usr/bin/env bash

set -e

mkdir -p data/

curl https://api.dane.gov.pl/media/resources/20200330/Wykaz_nazwisk_%C5%BCe%C5%84skich_uwzgl_os__zmar%C5%82e_2020-01-22.csv -o data/surnames_female.csv
curl https://api.dane.gov.pl/media/resources/20200330/Wykaz_nazwisk_m%C4%99skich_uwzgl_os__zmar%C5%82e_2020-01-22.csv -o data/surnames_male.csv
curl https://api.dane.gov.pl/media/resources/20200130/lista_imion_%C5%BCe%C5%84skich_os_%C5%BCyj%C4%85ce_2020-01-21.csv -o data/names_female.csv
curl https://api.dane.gov.pl/media/resources/20200130/lista_imion_m%C4%99skich_os_%C5%BCyj%C4%85ce_2020-01-21.csv -o data/names_male.csv
