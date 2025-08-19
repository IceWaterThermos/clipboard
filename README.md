# 다중 클립보드 관리자 🚀

Windows용 다중 클립보드 관리 프로그램입니다. 키보드 단축키를 사용하여 최대 8개의 클립보드 슬롯에 텍스트를 저장하고 불러올 수 있습니다.

[![Go Version](https://img.shields.io/badge/Go-1.25+-blue.svg)](https://golang.org/)
[![Platform](https://img.shields.io/badge/Platform-Windows-green.svg)](https://www.microsoft.com/windows)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## 📋 주요 기능

- **8개의 클립보드 슬롯**: F1~F8 키로 저장, 1~8 키로 불러오기
- **직관적인 키 조합**: `Ctrl + Shift` 조합으로 간단한 사용
- **실시간 모니터링**: 백그라운드에서 키보드 입력 감지
- **안전한 메모리 관리**: 스레드 안전 클립보드 관리
- **시각적 피드백**: 저장/불러오기 시 상태 표시

## 🎯 사용법

### 기본 키 조합

| 기능 | 키 조합 | 설명 |
|------|---------|------|
| **저장** | `Ctrl + Shift + F1~F8` | 현재 클립보드 내용을 해당 슬롯에 저장 |
| **불러오기** | `Ctrl + Shift + 1~8` | 해당 슬롯의 내용을 클립보드에 복사 |
| **슬롯 보기** | `Ctrl + Shift + F9` | 저장된 모든 슬롯 내용 확인 |
| **종료** | `Ctrl + C` | 프로그램 종료 |

### 슬롯 배치도

```
F1(0) F2(1) F3(2) F4(3) F5(4) F6(5) F7(6) F8(7)
 1(0)  2(1)  3(2)  4(3)  5(4)  6(5)  7(6)  8(7)
```

### 사용 예시

1. **텍스트 저장하기**
   - 원하는 텍스트를 복사 (`Ctrl + C`)
   - `Ctrl + Shift + F1` 눌러서 슬롯 0에 저장

2. **텍스트 불러오기**
   - `Ctrl + Shift + 1` 눌러서 슬롯 0의 내용을 클립보드에 복사
   - 원하는 곳에 붙여넣기 (`Ctrl + V`)

## 🔧 설치 및 빌드

### 필수 요구사항

- **Go 1.25 이상**
- **Windows 운영체제**
- **관리자 권한** (시스템 전역 키보드 모니터링을 위해 필요)

### 소스코드에서 빌드

```bash
# 저장소 클론
git clone https://github.com/YOUR_USERNAME/clipboard-manager.git
cd clipboard-manager

# 의존성 설치
go mod tidy

# 코드 포맷팅
go fmt ./...

# 빌드
go build -o clipboard-manager.exe

# 또는 직접 실행
go run clipboard_manager.go
```

### 최적화 빌드

성능을 위한 최적화 빌드:

```bash
# 최적화된 릴리즈 빌드
go build -ldflags="-s -w" -o clipboard-manager.exe

# 크기 최적화 (UPX 압축 - 선택사항)
# UPX 설치 후: upx --best clipboard-manager.exe
```

### 빌드 플래그 설명

- `-ldflags="-s -w"`: 디버그 정보 제거로 파일 크기 축소
  - `-s`: 심볼 테이블 제거
  - `-w`: DWARF 디버그 정보 제거

## 🚀 실행 방법

### 관리자 권한으로 실행

시스템 전역 키보드 모니터링을 위해 반드시 관리자 권한이 필요합니다:

1. **명령 프롬프트를 관리자로 실행**
2. 프로그램 경로로 이동
3. 실행: `clipboard-manager.exe`

또는

1. **실행 파일 우클릭**
2. **"관리자 권한으로 실행"** 선택

### 시작프로그램 등록 (선택사항)

Windows 시작 시 자동 실행하려면:

1. `Win + R` → `shell:startup` 입력
2. 실행 파일의 바로가기를 시작프로그램 폴더에 복사
3. 바로가기 속성에서 "관리자 권한으로 실행" 체크

## 🔍 문제 해결

### 자주 발생하는 문제

| 문제 | 원인 | 해결방법 |
|------|------|----------|
| 키 조합이 작동하지 않음 | 관리자 권한 부족 | 관리자 권한으로 실행 |
| 클립보드 읽기 실패 | 다른 프로그램이 클립보드 점유 | 잠시 후 다시 시도 |
| 프로그램이 실행되지 않음 | Go 런타임 누락 | Go 설치 확인 |

### 디버깅

문제 발생 시 확인사항:

```bash
# Go 버전 확인
go version

# 의존성 확인
go mod verify

# 빌드 테스트
go build -v
```

## 🏗️ 프로젝트 구조

```
clipboard-manager/
├── clipboard_manager.go   # 메인 소스코드
├── go.mod                # Go 모듈 설정
├── go.sum                # 의존성 체크섬
├── README.md             # 한글 문서 (이 파일)
├── README_EN.md          # 영문 문서
└── CLAUDE.md             # 개발자 가이드
```

## 📦 의존성

- `github.com/atotto/clipboard`: 크로스 플랫폼 클립보드 액세스
- `github.com/go-vgo/robotgo`: 시스템 자동화 (간접 의존성)

## ⚠️ 주의사항

- **Windows 전용**: 현재 Windows API를 직접 사용하므로 다른 OS에서는 동작하지 않습니다
- **관리자 권한 필요**: 시스템 전역 키보드 모니터링을 위해 필수
- **메모리 보안**: 클립보드 내용이 메모리에 평문으로 저장됩니다
- **한글 지원**: 유니코드 텍스트 완전 지원

## 🔄 버전 히스토리

- **v1.0.0**: 초기 릴리즈
  - 8개 슬롯 클립보드 관리
  - Windows API 기반 키보드 모니터링
  - 한글 UI 지원

## 📄 라이선스

이 프로젝트는 MIT 라이선스 하에 배포됩니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 확인하세요.


---

📖 **English documentation**: [README_EN.md](README_EN.md)