# health-dashboard

https://github.com/rafaeleyng/health-dashboard

## docker-compose

Install: https://devdojo.com/bobbyiliev/how-to-install-docker-and-docker-compose-on-raspberry-pi

## deploy

1. run
  ```sh
  git clone https://github.com/rafaeleyng/health-dashboard.git && cd health-dashboard
  ```

2. create a `.env` file containing:
  ```sh
  API_TZ=America/Sao_Paulo
  GIT_SYNC_USERNAME=rafaeleyng
  GIT_SYNC_PASSWORD=<use a personal token>
  GIT_SYNC_REPO=https://github.com/rafaeleyng/health-dashboard-data
  ```

3. run:
  ```sh
  make start
  ```

---
