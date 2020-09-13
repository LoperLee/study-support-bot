# study-support-bot

해당 프로젝트는 모각코의 커밋 관리 및 참여율 관리를 위한 슬랙 봇 프로젝트입니다.

필요에 따라 누구나 **수정, 배포, 테스트**할 수 있습니다.

## 기능

해당 봇은 매일 자정 등록된 참여자들의 Github 커밋을 확인합니다.

또한 모각코 모임 생성시 참여/불참 여부를 손쉽게 체크할 수 있습니다.

## 빌드 방법

```
# git clone https://github.com/LoperLee/study-support-bot.git
# go build
# ./study-support-bot
```

## 의존성

* [yaml](https://github.com/go-yaml/yaml)
* [slack](https://github.com/slack-go/slack)
* [cron](https://github.com/robfig/cron)
* [gorm](https://gorm.io)

## 라이선스

[MIT License](https://github.com/LoperLee/study-support-bot/blob/master/LICENSE)