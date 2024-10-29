# 프로젝트 설명
이미지 (jpg, jpeg, png) 를 webp 로 변환해주는 프로젝트입니다.  
폴더를 선택하면 위의 3개 이미지 파일들을 찾아 webp 로 변환해줍니다.

# 빌드
1. 현재 시스템에 맞는 버전 빌드
```bash
make build
```

2. 특정 시스템용 빌드
```bash
# Mac M1/M2/M3용
make build-mac-arm

# Mac Intel용
make build-mac-intel

# Windows용
make build-windows
```

3. 빌드 파일 정리
```bash
make clean
```

4. 시스템 정보 확인
```bash
make info
```

5. 도움말 보기
```bash
make help
```