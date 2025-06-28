<div align="center">
<img src="https://github.com/user-attachments/assets/568d0943-f1b1-4daa-8aae-0f670cf9e4ea" alt="image" width="400" height="400" />
<h1>RUSSIFIER ENA: DREAM BBQ</h1>
Repository of the Russifier for ENA: Dream BBQ by BARBEQUE TEAM
</div>

# WARNING
I am NOT a part of the "Russifier" project, I DO NOT know any Russian, I used Google Translate to translate all the Russian texts to English as a base for other people to use and translate the game in their own language. If you have any problems with the installation or with the "Russifier", report it to [ENAbbq_rus repo](https://github.com/bazelik-null/ENAbbq_rus/issues). If you have problems with the English text not being translated correctly, report it at [Issues](https://github.com/birbkeks/ENAbbq_en/issues).

## Contents
- [Description](#description)
- [Installation](#installation)
- [Compilation](#compilation)
- [Command](#command)
- [License](#license)
- [Channel](#channel)

![image](https://github.com/user-attachments/assets/1d688944-ed03-49d0-92c1-c87d71244569)

## Description
A Russifier for the game ENA: Dream BBQ translated entirely by hand. Includes translation of all dialogues, interface and main menu.

In this Russifier you will not find Russian voice acting that would kill the essence of the characters - a variety of foreign languages. Our project was aimed at Russian localization and adaptation of jokes, trying to convey the meaning of certain plays on words, puns and expressions as much as possible.

We use a self-written installer based on the [fyne](https://fyne.io/) framework for extensibility, beautiful design, transparency and convenience of our project.

>[!CAUTION]
>
>If you notice any inaccuracies, bugs or other problems, please create a bug report at [Issues](https://github.com/birbkeks/ENAbbq_en/issues).

Enjoy the game!

## Installation
### Automatic (recommended)
1. Download the installer from the [Releases] tab (https://github.com/birbkeks/ENAbbq_en/releases/).
2. Run the installer and install the Russifier.
**P.S. On Linux, you need to go to "Compatibility" in the game properties and select Proton!**

### Manually
1. Download all files (except SteamIcon.png and yarnmeta) from resources.
2. Move **resources.assets font_res.resS font_modern.resS** and everything from the **levels** folder to the ```/ENA Dream BBQ/ENA-4-DreamBBQ_Data/``` directory
3. Move **JoelG.ENA4.dll** and **YarnSpinner.Unity.dll** to the ```/ENA Dream BBQ/ENA-4-DreamBBQ_Data/Managed/``` directory
4. Move **catalog.json** to the ```/ENA Dream BBQ/ENA-4-DreamBBQ_Data/StreamingAssets/aa/``` directory
5. Move **yarndialogue_assets_all_*.bundle** to the ```/ENA Dream directory BBQ/ENA-4-DreamBBQ_Data/StreamingAssets/aa/StandaloneWindows64/```

### Hardpack
**What is a hardpack?** A hardpack is a packed build of the entire game. It is needed if the migration of changes to a new game update is delayed. *Hardpacks are not recommended to be installed just like that. It is best to use them only in cases where the update of the Russifier is delayed*
1. Go to the [Releases] tab (https://github.com/birbkeks/ENAbbq_en/releases/), follow the link to the TG channel with hardpacks and download the latest archive.
2. In Steam, click on the gear icon in the tab with ENA: Dream BBQ, go to "Management" and click on "View local files"
3. Unzip the contents of the archive into the game folder.

## Compilation
If you decide to compile the installer yourself, follow these steps:
1. Clone the repository:
```bash
git clone https://github.com/bazelik-null/ENAbbq_rus
```
**On Linux:**
  1. Install [go](https://go.dev/) and mingw-w64-gcc for cross-compilation.
  2. Install fyne:
  ```bash
  go get fyne.io/fyne/v2@latest
  ```
  3. Change to ``./src/`` directory
  4. Compile:
  ```bash
  make linux ; make windows
  ```
  5. Move the executables to the root of the repository directory.

**On Windows**
  1. Install [go](https://go.dev/).
  2. Install fyne
  ```bash
  go get fyne.io/fyne/v2@latest
  ```
  3. Go to the ``./src/windows/`` directory
  4. Compile for windows:
  ```bash
  go build -ldflags="-s -w -H=windowsgui" -o Installer-Windows.exe *.go
  ```
  5. Move the executables to the root of the repository directory.

  P.S. For cross-compilation, read the guide [here](https://docs.fyne.io/started/cross-compiling).

## Team
**Our team: (Russifier)**

@fiverebbles - Technical part of the Russifier, team coordination.

@anpatu - Editing, help in decision-making.

@felesneveve - Translation, editing, TG channel management.

@big_fan_of_kiwi - Translation.

@bagaxdd - Translation.

@hugginggg - Translation.

@PhONaRr - Translation.

@DetectivePrince - Translation.

@nucl3arsnake - Translation, translation of the entire interface.

@dorime_lolk4075 - Translation.

@kaimanDoppel - Translation.

@SalvetEna - Translation.

@freeze2222 - Font.

## F.A.Q

*1. Question:* Why aren't all elements translated? (Ex. 3D text in cutscenes)

***Answer:*** There are no plans to translate them - we consider this an interference with the author's vision and visual style of the work.

*2. Question:* I want to contribute! How can I do this?

***Answer:*** You can do this in the following ways: Spread the translation, report bugs, follow updates!

*3. Question:* Do you have anything besides Telegram and GitHub?

***Answer:*** Yes, we have a [Guide on Steam (Russifier)](https://steamcommunity.com/sharedfiles/filedetails/?id=3453809143)!

## License
This project is licensed under the open license [MIT](https://mit-license.org/). You are free to use, modify and distribute this Russifier in accordance with the terms of the license.

## Channel:
<div align="center">
<img src="https://github.com/user-attachments/assets/d5718154-17b2-49a8-98be-c71cc5d5cacd" alt="image" width="335" height="470" />
</div>
