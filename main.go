package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/chai2010/webp"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Image Converter to WebP")
	myWindow.Resize(fyne.NewSize(800, 600))

	// 진행 상황을 표시할 레이블
	statusLabel := widget.NewLabel("이미지 파일들을 포함한 폴더를 선택하여 WebP로 변환하세요.")

	uploadButton := widget.NewButton("폴더 선택", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil || list == nil {
				return
			}

			// 진행 상황 초기화
			statusLabel.SetText("변환 시작...")

			// 이미지 파일 찾기 및 변환
			go func() {
				processedCount := 0
				errorCount := 0

				err := filepath.Walk(list.Path(), func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}

					// 파일이 아니면 건너뜀
					if info.IsDir() {
						return nil
					}

					// 파일 확장자 확인
					ext := strings.ToLower(filepath.Ext(path))
					if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
						err := convertToWebP(path)
						if err != nil {
							errorCount++
							fmt.Printf("Error converting %s: %v\n", path, err)
						} else {
							processedCount++
						}

						// UI 업데이트
						statusLabel.SetText(fmt.Sprintf("처리중... %d개 완료, %d개 실패", processedCount, errorCount))
					}
					return nil
				})

				// 최종 결과 표시
				if err != nil {
					dialog.ShowError(err, myWindow)
					statusLabel.SetText("변환 중 오류가 발생했습니다.")
				} else {
					message := fmt.Sprintf("변환 완료!\n총 %d개의 이미지 처리됨, %d개 실패", processedCount, errorCount)
					dialog.ShowInformation("완료", message, myWindow)
					statusLabel.SetText("변환이 완료되었습니다. 다른 폴더를 선택하려면 '폴더 선택' 버튼을 클릭하세요.")
				}
			}()

		}, myWindow)
	})

	myWindow.SetContent(container.NewVBox(
		statusLabel,
		uploadButton,
	))

	myWindow.ShowAndRun()
}

func convertToWebP(inputFilePath string) error {
	// 입력 파일 열기
	file, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("파일 열기 실패: %v", err)
	}
	defer file.Close()

	// 이미지 디코딩
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("이미지 디코딩 실패: %v", err)
	}

	outputFilePath := strings.TrimSuffix(inputFilePath, filepath.Ext(inputFilePath)) + ".webp"

	// WebP 파일 생성
	outFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("출력 파일 생성 실패: %v", err)
	}
	defer outFile.Close()

	// WebP로 인코딩
	err = webp.Encode(outFile, img, &webp.Options{Lossless: true})
	if err != nil {
		return fmt.Errorf("WebP 인코딩 실패: %v", err)
	}

	return nil
}
