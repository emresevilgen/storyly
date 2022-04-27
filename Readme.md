# To Run

Copy the data folder in postgresql folder at the root of the project into ~/Desktop/temp

Pull the docker images below and run

- $ docker pull postgres
- $ docker run --name postgresql -e POSTGRES_USER=myusername -e POSTGRES_PASSWORD=mypassword -p 5432:5432 -v ~
  /Desktop/temp/data:/var/lib/postgresql/data -d postgres

- $ docker pull dpage/pgadmin4:latest
- $ docker run --name my-pgadmin -p 82:80 -e 'PGADMIN_DEFAULT_EMAIL=user@domain.local' -e '
  PGADMIN_DEFAULT_PASSWORD=postgresmaster' -d dpage/pgadmin4

- $ docker build . -t storyly
- $ docker run --name storyly -p 8086:8086 storyly

Swagger: http://localhost:8086/swagger/index.html
Example Curl:

curl -X 'GET' \
'http://localhost:8086/stories/token_1' \
-H 'accept: application/json' \
-H 'correlationId: 1' \
-H 'agentName: 2' \
-H 'executorUser: 3'

# Before Enhancement

ab -c 100 -n 10000 -H "correlationId:test" -H "agentName:test" -H "executorUser:
test" http://localhost:8086/stories/token_1

- Concurrency Level:      100
- Time taken for tests:   31.355 seconds
- Complete requests:      10000
- Failed requests:        0
- Total transferred:      2650000 bytes
- HTML transferred:       1410000 bytes
- Requests per second:    318.93 [#/sec] (mean)
- Time per request:       313.551 [ms] (mean)
- Time per request:       3.136 [ms] (mean, across all concurrent requests)
- Transfer rate:          82.53 [Kbytes/sec] received

Connection Times (ms)
min mean[+/-sd] median max Connect:        0 3 8.2 1 231 Processing:     3 309 246.6 261 1909 Waiting:        3 307
246.3 261 1909 Total:          3 312 249.8 262 1917

Percentage of the requests served within a certain time (ms)
50% 262 66% 340 75% 408 80% 454 90% 638 95% 810 98% 929 99% 1182 100% 1917 (longest request)

# After Enhancement

ab -c 100 -n 10000 -H "correlationId:test" -H "agentName:test" -H "executorUser:
test" http://localhost:8086/stories/token_1

- Concurrency Level:      100
- Time taken for tests:   2.448 seconds
- Complete requests:      10000
- Failed requests:        0
- Total transferred:      2650000 bytes
- HTML transferred:       1410000 bytes
- Requests per second:    4084.52 [#/sec] (mean)
- Time per request:       24.483 [ms] (mean)
- Time per request:       0.245 [ms] (mean, across all concurrent requests)
- Transfer rate:          1057.03 [Kbytes/sec] received

Connection Times (ms)
min mean[+/-sd] median max Connect:        0 10 10.3 10 265 Processing:     1 14 53.7 9 754 Waiting:        1 14 53.7 8
754 Total:          5 24 54.6 18 762

Percentage of the requests served within a certain time (ms)
50% 18 66% 19 75% 20 80% 21 90% 24 95% 29 98% 35 99% 266 100% 762 (longest request)