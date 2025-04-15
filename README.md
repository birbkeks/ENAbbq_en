<div align="center">
<h1>TRANSLATING ENA: DREAM BBQ</h1>
Translator repository for ENA: Dream BBQ from BARBEQUE TEAM
</div>

>[!CAUTION]
>
>THIS IS A ENGLISH TRANSLATION OF THE [ENAbbq_ru](https://github.com/bazelik-null/ENAbbq_rus) REPO <br>
>I DID NOT MADE ANY OF THIS, I ONLY TRANSLATED THE REPO TO ENGLISH SO OTHERS COULD UNDERSTAND IT BETTER. <br>
>IF YOU HAVE ANY PROBLEMS WITH THE TRANSLATOR, GO TO THE ORIGINAL REPO AND REPORT IT THERE. <br>
>
>I DO NOT KNOW ANY RUSSIAN, I USED GOOGLE TRANSLATE FOR ALL THE RUSSIAN TEXTS.

## Contents
- [Description](#Description)
- [Installation](#Installation)
- [Compiling](#Compiling)
- [Team](#Team)
- [License](#License)
- [Channel](#Channel)


## Description
Translator for the game ENA: Dream BBQ translated entirely manually. Includes translation of (almost) all dialogues, interface and main menu.

In this Translator you will not find voice acting other than Rnglish, which would kill the essence of the characters - the diversity of foreign languages. Our project was aimed at translation to other languages and adaptation of jokes, trying to convey the meaning of certain wordplays, puns and expressions as much as possible.

We use a custom installer based on the [fyne](https://fyne.io/) framework for extensibility, beautiful design, transparency and convenience of our project.

>[!CAUTION]
>
>The Translator is in beta version. If you notice mistakes, bugs or other problems, create a bug report in [Issues](https://github.com/birbkeks/ENAbbq_en/issues).

Enjoy!

## Installation
### Automatically (Recommended)
1. Download the installer from [Releases](https://github.com/birbkeks/ENAbbq_en/releases/).
2. Run the installer and install the Translator.
   **P.S. On Linux, you need to go to "Compatibility" in the game properties and select Proton!**

### Manually
1. Download all files (except SteamIcon.png and yarnmeta) from resources.
2. Move **resources.assets font_res.resS font_modern.resS** and everything from the folder **levels** to the ```/ENA Dream BBQ/ENA-4-DreamBBQ_Data/``` directory.
3. Move **JoelG.ENA4.dll** and **YarnSpinner.Unity.dll** to the ```/ENA Dream BBQ/ENA-4-DreamBBQ_Data/Managed/``` directoy.
4. Move **catalog.json** to the ```/ENA Dream BBQ/ENA-4-DreamBBQ_Data/StreamingAssets/aa/``` directoy.
5. Move **yarndialogue_assets_all_*.bundle** to the ```/ENA Dream BBQ/ENA-4-DreamBBQ_Data/StreamingAssets/aa/StandaloneWindows64/``` directory.

### Hardpack
**What is a hardpack?** A hardpack is a packed build of the entire game. It is needed if the migration of changes to a new game update is delayed. *It is not recommended to install hardpacks just like that. It is best to use them only in cases where the update of the Translator is delayed*
1. Go to the [Releases](https://github.com/birbkeks/ENAbbq_en/releases/), follow the link to the TG channel with hardpacks and download the latest archive.
2. In Steam, click on the gear icon in the ENA: Dream BBQ tab, go to "Manage" and click on "Browse Local Files"
3. Unzip the contents of the archive into the game folder.

## Compiling
If you decide to compile the installer yourself, follow these steps:
1. Clone the repository:
```bash
git clone https://github.com/birbkeks/ENAbbq_en
```
**On Linux:**
  1. Install [go](https://go.dev/) and mingw-w64-gcc for cross-compilation.
  2. Install fyne:
  ```bash
  go get fyne.io/fyne/v2@latest
  ```
  3. Go to the ``./src/`` directoy
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

      P.S. For cross-compilation, read [this](https://docs.fyne.io/started/cross-compiling) guide.

## Team
**Our Team:**

@fiverebbles - Technical part of the Russifier, team coordination.

@anpatu -  Editorial, assistance in decision making.

@felesneveve - Translation, editing, running of the TG channel.

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
*1. Question:* Why aren't all elements translated? (Eg. 3D text in cutscenes)

***Answer:*** There are no plans to translate them - we consider this an interference with the author's vision and visual style of the work.

*2. Question:* I want to contribute! How can I do this?

***Answer:*** You can do this in the following way: Spread the translation, report errors, follow updates!

*3. Question:* Do you have anything besides Telegram and GitHub?

***Answer:*** Yes, we have a [Steam Guide](https://steamcommunity.com/sharedfiles/filedetails/?id=3453809143)!

## License
This project is licensed under the open license [MIT](https://mit-license.org/). You are free to use, modify and distribute this Russifier in accordance with the terms of the license.

## Channel:
<div align="center">
<img src="https://github.com/user-attachments/assets/d5718154-17b2-49a8-98be-c71cc5d5cacd" alt="image" width="335" height="470" />
</div>
