<p align=center>
  <br>
  <img width="500px" src="Myimx-Title.png"/>
  <br>
  <span><b>A simple ASCII art generator built in Go</b></span>
  <br>
</p>

<p align="center">
  <a href="#installation">Installation</a>
  &nbsp;&nbsp;&nbsp;•&nbsp;&nbsp;&nbsp;
  <a href="#usage">Usage</a>
  &nbsp;&nbsp;&nbsp;•&nbsp;&nbsp;&nbsp;
  <a href="#contribute">Contribute</a>
</p>
<br>

<a name="installation"></a>
## Installation
> [!NOTE]
> Myimx at its current version is only available to __windows__ users. <br>
> Support for MacOS / Linux will probably be available sometime in the future.

To install the latest version of __Myimx__ run:
```powershell
Invoke-WebRequest -Uri 'https://raw.githubusercontent.com/JimmyTarson12/Myimx/main/install-myimx.ps1' -OutFile "$env:TEMP\install-myimx.ps1" -UseBasicParsing; Set-ExecutionPolicy Bypass -Scope Process -Force; & "$env:TEMP\install-myimx.ps1"
```

To install an earlier version of Myimx run the [install-myimx.ps1](https://github.com/JimmyTarson12/Myimx/blob/main/install-myimx.ps1) script with a different supported version in the `$exeUrl` variable

<a name="usage"></a>
## Usage
Print an ASCII art from [github](https://github.com/search?q=repo%3AJimmyTarson12%2FMyimx+path%3Ainternal%2Fart+path%3A*.md&type=code) or [locally](#local)
```bash
myimx <art-name>
```
List all local ASCII art you have installed:
```bash
myimx list
```
Display all commands:
```bash
myimx help
```

<a name="local"></a>
## Create your own ASCII art
> [!WARNING]
> This feature is in Beta, there may be issues

1 Create a new Markdown file with your art:
```powershell
mkdir -Force "$env:USERPROFILE\AppData\Local\Programs\myimx\art"

# Open it in notepad
notepad "$env:USERPROFILE\AppData\Local\Programs\myimx\art\fish.md"
```
2 Add your ASCII art to the file:
```
><((°>
```
3 Save the file and close the editor

4 Now you can use your new art:
```bash
myimx fish
```

<a name="contribute"></a>
## How to contribute

1. Fork this repository
2. Create a file `<art-name>.md` in the `./internal/art` directory
3. Paste your ASCII art into that file and then commit the changes
4. Start a pull request

When the pull request is accepted you will be able to use your art:
```bash
myimx <art-name>
```
