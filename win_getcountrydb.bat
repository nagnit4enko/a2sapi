@echo off
wget http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.mmdb.gz
gzip -d GeoLite2-City.mmdb.gz
mkdir db
move GeoLite2-City.mmdb db/