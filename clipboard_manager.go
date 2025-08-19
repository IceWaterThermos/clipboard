package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type ClipboardManager struct {
	clipboards map[int]string
	mutex      sync.RWMutex
}

func NewClipboardManager() *ClipboardManager {
	return &ClipboardManager{
		clipboards: make(map[int]string),
	}
}

func (cm *ClipboardManager) SaveToSlot(slot int, content string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	cm.clipboards[slot] = content
	fmt.Printf("✅ 클립보드 슬롯 %d에 저장됨: %s\n", slot, truncateString(content, 50))
}

func (cm *ClipboardManager) LoadFromSlot(slot int) (string, bool) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	content, exists := cm.clipboards[slot]
	return content, exists
}

func (cm *ClipboardManager) ShowSlots() {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	fmt.Println("\n=== 현재 저장된 슬롯들 ===")
	fmt.Println("F1(0) F2(1) F3(2) F4(3) F5(4) F6(5) F7(6) F8(7)")
	fmt.Println(" 1(0)  2(1)  3(2)  4(3)  5(4)  6(5)  7(6)  8(7)")
	fmt.Println("-----------------------------------------------")

	if len(cm.clipboards) == 0 {
		fmt.Println("저장된 슬롯이 없습니다.")
	} else {
		for i := 0; i < 8; i++ {
			if content, exists := cm.clipboards[i]; exists {
				fmt.Printf("슬롯 %d: %s\n", i, truncateString(content, 50))
			}
		}
	}
	fmt.Println("=========================================\n")
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

// Windows API 함수들
var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)

// Windows 가상 키 코드
const (
	VK_CONTROL = 0x11
	VK_SHIFT   = 0x10

	// 함수키 F1~F9
	VK_F1 = 0x70
	VK_F2 = 0x71
	VK_F3 = 0x72
	VK_F4 = 0x73
	VK_F5 = 0x74
	VK_F6 = 0x75
	VK_F7 = 0x76
	VK_F8 = 0x77
	VK_F9 = 0x78

	// 숫자키 1~8
	VK_1 = 0x31
	VK_2 = 0x32
	VK_3 = 0x33
	VK_4 = 0x34
	VK_5 = 0x35
	VK_6 = 0x36
	VK_7 = 0x37
	VK_8 = 0x38
)

// 키가 눌렸는지 확인하는 함수
func isKeyPressed(vkCode int) bool {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(vkCode))
	// 최상위 비트가 1이면 키가 눌린 상태
	return (ret & 0x8000) != 0
}

// 키 이름을 반환하는 헬퍼 함수
func getKeyName(vkCode int) string {
	switch vkCode {
	case VK_CONTROL:
		return "Ctrl"
	case VK_SHIFT:
		return "Shift"
	case VK_F1:
		return "F1"
	case VK_F2:
		return "F2"
	case VK_F3:
		return "F3"
	case VK_F4:
		return "F4"
	case VK_F5:
		return "F5"
	case VK_F6:
		return "F6"
	case VK_F7:
		return "F7"
	case VK_F8:
		return "F8"
	case VK_F9:
		return "F9"
	case VK_1:
		return "1"
	case VK_2:
		return "2"
	case VK_3:
		return "3"
	case VK_4:
		return "4"
	case VK_5:
		return "5"
	case VK_6:
		return "6"
	case VK_7:
		return "7"
	case VK_8:
		return "8"
	default:
		return fmt.Sprintf("키(%d)", vkCode)
	}
}

func main() {
	fmt.Println("🚀 다중 클립보드 관리자가 시작되었습니다!")
	fmt.Println("\n사용법:")
	fmt.Println("  Ctrl + Shift + F1~F8: 현재 클립보드 내용을 해당 슬롯에 저장")
	fmt.Println("  Ctrl + Shift + 1~8: 해당 슬롯의 내용을 클립보드에 붙여넣기")
	fmt.Println("  Ctrl + Shift + F9: 저장된 모든 슬롯 보기")
	fmt.Println("  Ctrl + C: 프로그램 종료")
	fmt.Println("\n📋 슬롯 배치:")
	fmt.Println("  F1(0) F2(1) F3(2) F4(3) F5(4) F6(5) F7(6) F8(7)")
	fmt.Println("   1(0)  2(1)  3(2)  4(3)  5(4)  6(5)  7(6)  8(7)")
	fmt.Println("\n💡 키 조합 팁:")
	fmt.Println("  • 저장: Ctrl + Shift + F1 (함수키 사용)")
	fmt.Println("  • 불러오기: Ctrl + Shift + 1 (숫자키 사용)")
	fmt.Println("  • 간단하고 직관적인 조합입니다")
	fmt.Println("\n⚠️  관리자 권한으로 실행하세요!")

	manager := NewClipboardManager()

	// 종료 신호 처리
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// 키보드 상태를 폴링하는 고루틴
	go func() {
		ticker := time.NewTicker(30 * time.Millisecond) // 30ms로 더 빠르게
		defer ticker.Stop()

		var lastAction time.Time
		//var prevKeys = make(map[int]bool) // 이전 키 상태만 추적

		fmt.Println("\n🔍 키보드 감지를 시작합니다...")
		fmt.Println("💡 테스트: Ctrl + Shift + F1로 저장, Ctrl + Shift + 1로 불러오기!")

		for {
			select {
			case <-ticker.C:
				// 현재 모든 키 상태 확인
				currentStates := make(map[int]bool)

				// 수정자 키들
				currentStates[VK_CONTROL] = isKeyPressed(VK_CONTROL)
				currentStates[VK_SHIFT] = isKeyPressed(VK_SHIFT)
				currentStates[VK_F9] = isKeyPressed(VK_F9)

				// 함수키 F1~F8 (저장용)
				fKeys := []struct {
					vk   int
					name string
					slot int
				}{
					{VK_F1, "F1", 0}, {VK_F2, "F2", 1}, {VK_F3, "F3", 2}, {VK_F4, "F4", 3},
					{VK_F5, "F5", 4}, {VK_F6, "F6", 5}, {VK_F7, "F7", 6}, {VK_F8, "F8", 7},
				}

				// 숫자키 1~8 (불러오기용)
				numKeys := []struct {
					vk   int
					name string
					slot int
				}{
					{VK_1, "1", 0}, {VK_2, "2", 1}, {VK_3, "3", 2}, {VK_4, "4", 3},
					{VK_5, "5", 4}, {VK_6, "6", 5}, {VK_7, "7", 6}, {VK_8, "8", 7},
				}

				for _, key := range fKeys {
					currentStates[key.vk] = isKeyPressed(key.vk)
				}

				for _, key := range numKeys {
					currentStates[key.vk] = isKeyPressed(key.vk)
				}

				// 키 상태 변화 감지
				now := time.Now()

				// 새로 눌린 키들에 대한 디버깅 정보
				//for vk, pressed := range currentStates {
				//	if pressed && !prevKeys[vk] {
				//		fmt.Printf("🔍 %s 키가 눌렸습니다\n", getKeyName(vk))
				//	}
				//}

				// 조합 키 검사
				if time.Since(lastAction) > 300*time.Millisecond {

					// Ctrl + Shift + F9 (슬롯 보기)
					if currentStates[VK_CONTROL] && currentStates[VK_SHIFT] && currentStates[VK_F9] {
						manager.ShowSlots()
						lastAction = now
					}

					// Ctrl + Shift + F1~F8 (저장)
					for _, key := range fKeys {
						if currentStates[VK_CONTROL] && currentStates[VK_SHIFT] && currentStates[key.vk] {
							// 클립보드 읽기 시도 (여러 번 재시도)
							var content string
							var err error

							for retry := 0; retry < 3; retry++ {
								time.Sleep(10 * time.Millisecond)
								content, err = clipboard.ReadAll()
								if err == nil {
									break
								}
								fmt.Printf("⏳ 클립보드 읽기 재시도 %d/3...\n", retry+1)
							}

							if err != nil {
								fmt.Printf("❌ 클립보드 읽기 실패: %v\n", err)
								fmt.Printf("💡 다른 프로그램에서 텍스트를 복사한 후 다시 시도해보세요.\n")
							} else if content == "" {
								fmt.Printf("⚠️  클립보드가 비어있습니다. 텍스트를 복사한 후 다시 시도해보세요.\n")
							} else {
								manager.SaveToSlot(key.slot, content)
								fmt.Printf("💾 Ctrl+Shift+%s로 슬롯 %d에 저장완료!\n", key.name, key.slot)
								lastAction = now
							}
						}
					}

					// Ctrl + Shift + 1~8 (불러오기)
					for _, key := range numKeys {
						if currentStates[VK_CONTROL] && currentStates[VK_SHIFT] && currentStates[key.vk] {
							if content, exists := manager.LoadFromSlot(key.slot); exists {
								// 클립보드 쓰기 시도 (여러 번 재시도)
								var err error
								for retry := 0; retry < 3; retry++ {
									err = clipboard.WriteAll(content)
									if err == nil {
										break
									}
									time.Sleep(10 * time.Millisecond)
									fmt.Printf("⏳ 클립보드 쓰기 재시도 %d/3...\n", retry+1)
								}

								if err != nil {
									fmt.Printf("❌ 클립보드 쓰기 오류: %v\n", err)
								} else {
									fmt.Printf("📋 Ctrl+Shift+%s로 슬롯 %d 불러오기 완료: %s\n", key.name, key.slot, truncateString(content, 30))
									lastAction = now
								}
							} else {
								fmt.Printf("⚠️  슬롯 %d가 비어있습니다.\n", key.slot)
							}
						}
					}
				}

				// 상태 업데이트
				//prevKeys = currentStates

			case <-c:
				fmt.Println("\n👋 프로그램을 종료합니다...")
				return
			}
		}
	}()

	// 초기 테스트
	fmt.Println("\n🧪 클립보드 상태 테스트 중...")
	time.Sleep(100 * time.Millisecond)

	// 클립보드 읽기 테스트
	if content, err := clipboard.ReadAll(); err == nil {
		if content != "" {
			fmt.Printf("✅ 클립보드 읽기 성공: %s\n", truncateString(content, 30))
		} else {
			fmt.Println("ℹ️  클립보드가 비어있습니다. 텍스트를 복사한 후 테스트해보세요.")
		}
	} else {
		fmt.Printf("⚠️  클립보드 읽기 테스트 실패: %v\n", err)
		fmt.Println("💡 관리자 권한으로 실행하거나 다른 프로그램에서 텍스트를 복사해보세요.")
	}

	// 클립보드 쓰기 테스트
	testText := "클립보드 테스트"
	if err := clipboard.WriteAll(testText); err == nil {
		fmt.Println("✅ 클립보드 쓰기 성공")
		// 원래 내용이 있었다면 복원
		if content, _ := clipboard.ReadAll(); content != "" && content != testText {
			clipboard.WriteAll(content)
		}
	} else {
		fmt.Printf("⚠️  클립보드 쓰기 테스트 실패: %v\n", err)
	}

	if isKeyPressed(VK_CONTROL) && isKeyPressed(VK_SHIFT) {
		fmt.Println("✅ Ctrl + Shift 조합이 현재 눌려있습니다!")
	} else {
		fmt.Println("ℹ️  키 감지가 준비되었습니다. Ctrl + Shift + F1을 테스트해보세요.")
	}

	// 메인 스레드는 신호를 기다림
	<-c
	fmt.Println("\n프로그램이 종료되었습니다.")
}
