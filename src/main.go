package main

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Init app
	a := app.New()
	w := a.NewWindow("Russifier installer for ENA: Dream BBQ")
	w.Resize(fyne.NewSize(800, 600))

	mainLabel := widget.NewLabel("Please wait while we load installer resources...")
	loadingWidget := widget.NewActivity()
	loadingWidget.Start()

	init := container.New(layout.NewCenterLayout(),
		container.New(layout.NewVBoxLayout(),
			mainLabel,
			loadingWidget,
		),
	)

	w.SetContent(init)
	w.Show()

	// Download files in goroutine
	go func() {
		err := download()
		if err != 0 {
			w.SetContent(pageERR(w, err))
		} else {
			w.SetContent(page0(w))
		}
	}()

	fyne.CurrentApp().Run()
}

func page0(w fyne.Window) *fyne.Container {
	mainLabel := widget.NewLabel("Welcome to the Russifier installer for ENA: Dream BBQ")

	teamLabel := canvas.NewText("  by BARBEQUE TEAM", color.RGBA{169, 169, 169, 255})
	teamLabel.TextSize = 12

	errorLabel := canvas.NewText("", color.RGBA{255, 0, 0, 255})

	btnContinue := widget.NewButton("Continue", func() {
		w.SetContent(pageInstall(w))
	})

	page0 := container.New(layout.NewCenterLayout(),
		container.New(layout.NewVBoxLayout(),
			mainLabel,
			teamLabel,
			btnContinue,
			errorLabel,
		),
	)

	// Check integrity of downloaded files
	checkIntegrity(btnContinue, errorLabel)

	return page0
}

func pageERR(_ fyne.Window, err int) *fyne.Container {
	errLabel := canvas.NewText("[FATL]: A critical error occurred while loading files.", color.RGBA{255, 0, 0, 255})
	errCode := canvas.NewText("[FATL]: Error "+fmt.Sprint(err), color.RGBA{255, 0, 0, 255})

	buttonClose := widget.NewButtonWithIcon("Close", theme.WindowCloseIcon(), func() {
		fyne.CurrentApp().Quit()
	})

	pageERRContainer := container.New(layout.NewCenterLayout(),
		container.New(layout.NewVBoxLayout(),
			errLabel,
			errCode,
			buttonClose,
		),
	)
	return pageERRContainer
}

func pageInstall(w fyne.Window) *fyne.Container {
	var path string
	var steamIcon fyne.Resource

	appDir, _ := os.Getwd()
	currentTheme := fyne.CurrentApp().Settings().Theme()

	// Load an icon depending on theme
	if currentTheme == theme.DarkTheme() {
		steamIcon, _ = fyne.LoadResourceFromPath(filepath.Join(appDir, "resources", "steamIconW.png"))
	} else {
		steamIcon, _ = fyne.LoadResourceFromPath(filepath.Join(appDir, "resources", "steamIcon.png"))
	}

	label := widget.NewLabel("Select the path to the game. If it is installed in the standard path, click on the Steam button.")
	labelPath := widget.NewLabel("")
	errorLabel := canvas.NewText("", color.RGBA{255, 0, 0, 255})

	btnContinue := widget.NewButtonWithIcon("Install", theme.DownloadIcon(), func() {
		w.SetContent(pageEnd(path))
	})
	btnContinue.Disable()

	btnSteam := widget.NewButtonWithIcon("Steam", steamIcon, func() {
		// Choose default path to game depending on OS
		if runtime.GOOS == "windows" {
			path = filepath.Join("C:\\", "Program Files (x86)", "Steam", "steamapps", "common", "ENA Dream BBQ")
		} else {
			homeDir := os.Getenv("HOME")
			path = filepath.Join(homeDir, ".steam", "root", "steamapps", "common", "ENA Dream BBQ")
		}
		// Check if there is executable game file
		checkExecutable(path, btnContinue, errorLabel)
		// Display chosen path
		labelPath.SetText("The chosen path: " + path)
	})

	btnBrowse := widget.NewButtonWithIcon("Open", theme.SearchIcon(), func() {
		browseFile(w, func(selectedPath string) {
			path = selectedPath
			// Check if there is executable game file
			checkExecutable(path, btnContinue, errorLabel)
			// Display chosen path
			labelPath.SetText("The chosen path: " + path)
		})
	})

	pageInstall := container.New(layout.NewCenterLayout(),
		container.New(layout.NewVBoxLayout(),
			label,
			btnSteam,
			btnBrowse,
			labelPath,
			errorLabel,
			btnContinue,
		),
	)

	return pageInstall
}

func checkIntegrity(btnContinue *widget.Button, errorLabel *canvas.Text) {
	appDir, _ := os.Getwd()
	resourcesPath := filepath.Join(appDir, "resources", "meta.json")
	if _, err := os.Stat(resourcesPath); os.IsNotExist(err) {
		btnContinue.Disable()
		errorLabel.Text = "[FATL]: \"resources\" not found."
		errorLabel.Refresh()
	} else {
		btnContinue.Enable()
		errorLabel.Text = ""
		errorLabel.Refresh()
	}
}

func checkExecutable(selectedPath string, btnContinue *widget.Button, errorLabel *canvas.Text) {
	executablePath := filepath.Join(selectedPath, "ENA-4-DreamBBQ.exe")
	if _, err := os.Stat(executablePath); os.IsNotExist(err) {
		btnContinue.Disable()
		errorLabel.Text = "[ERROR]: \"ENA-4-DreamBBQ.exe\" not found, select the folder with the game executable file"
		errorLabel.Refresh()
	} else {
		btnContinue.Enable()
		errorLabel.Text = ""
		errorLabel.Refresh()
	}
}

func browseFile(w fyne.Window, onPathSelected func(string)) {
	dialog.ShowFolderOpen(func(folder fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		if folder != nil {
			onPathSelected(folder.Path())
		}
	}, w)
}

func pageEnd(path string) *fyne.Container {
	appDir, _ := os.Getwd()

	err := install(path)

	os.RemoveAll(filepath.Join(appDir, "resources"))

	if err != nil {
		errLabel := canvas.NewText("[FATL]: A critical error occurred while injecting assets.", color.RGBA{255, 0, 0, 255})
		errCode := canvas.NewText("[FATL]: Error "+fmt.Sprint(err), color.RGBA{255, 0, 0, 255})

		buttonClose := widget.NewButtonWithIcon("Close", theme.WindowCloseIcon(), func() {
			fyne.CurrentApp().Quit()
		})

		pageEndContainer := container.New(layout.NewCenterLayout(),
			container.New(layout.NewVBoxLayout(),
				errLabel,
				errCode,
				buttonClose,
			),
		)
		return pageEndContainer
	} else {
		label := widget.NewLabel("Thank you for installing")
		buttonClose := widget.NewButtonWithIcon("Close", theme.WindowCloseIcon(), func() {
			fyne.CurrentApp().Quit()
		})

		pageEndContainer := container.New(layout.NewCenterLayout(),
			container.New(layout.NewVBoxLayout(),
				label,
				buttonClose,
			),
		)
		return pageEndContainer
	}
}
