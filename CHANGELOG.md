# Change Log

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

<a name="1.0.0"></a>
# 1.0.0 (2017-08-04)


### Bug Fixes

* add data/.keep ([376695c](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/376695c))
* boundary check ([f995e73](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/f995e73))
* return 500 Error when db connection error has occured ([7cb8998](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/7cb8998))


### Features

* 2.「国民の祝日」が日曜日に当たるときは、その日後においてその日に最も近い「国民の祝日」でない日を休日とする。 ([8e5c325](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/8e5c325))
* 3.その前日及び翌日が「国民の祝日」である日（「国民の祝日」でない日に限る。）は、休日とする。 ([fc05556](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/fc05556))
* add "debug" option ([3b6ec5a](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/3b6ec5a))
* add "from" and "to" paramater ([3f56ac2](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/3f56ac2))
* add "isOtherHolidaysStored" option ([0a6ef38](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/0a6ef38))
* add "storage" option ([ee88025](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/ee88025))
* add Day Of Week ([0740846](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/0740846))
* add Dockerfile and docker-compose.yml ([30952cb](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/30952cb))
* add health-check ([553f38b](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/553f38b))
* add sjis_csv storage option ([942a604](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/942a604))
* add Sunday ([a9d9e05](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/a9d9e05))
* allow to configure application by environment variables ([171cc4d](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/171cc4d))
* allow to configure the application port number with toml and ([34e5943](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/34e5943))
* configure rdb in config file ([0161799](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/0161799))
* configure startDate and endDate in configuration file ([030c46e](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/030c46e))
* create database and insert holidays ([3ee6385](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/3ee6385))
* get holidays from db ([97e93b3](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/97e93b3))
* read csv ([570f65b](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/570f65b))
* return JSON ([1788877](https://github.com/suzuki-shunsuke/japanese-holiday-api/commit/1788877))
