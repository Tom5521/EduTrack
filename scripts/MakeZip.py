import shutil
import requests,os


if not os.path.exists("tmp"):
    os.mkdir("tmp")

if not os.path.exists("tmp/opengl32.7z"):
    print("dowloading opengl32.dll...")
    url = 'https://downloads.fdossena.com/geth.php?r=mesa64-latest'
    r = requests.get(url)
    with open('tmp/opengl32.7z', 'wb') as f:
        f.write(r.content)

if not os.path.exists("tmp/opengl"):
    print("unzipping 7z file...")
    os.mkdir("tmp/opengl")
    os.system("7z e -o\"tmp/opengl\" tmp/opengl32.7z")

print("Compressing for windows...")
if os.path.exists("builds/EduTrack-win64.zip"):
    os.remove("builds/EduTrack-win64.zip")
if not os.path.exists("EduTrack-win64/"):
    os.mkdir("EduTrack-win64/")
shutil.copy2("builds/EduTrack.exe","EduTrack-win64/")
shutil.copy2("tmp/opengl/opengl32.dll","EduTrack-win64/")
os.system("zip -r builds/EduTrack-win64.zip EduTrack-win64/")
shutil.rmtree("EduTrack-win64/")

print("Compressing for linux...")
if os.path.exists("builds/EduTrack-linux64.zip"):
    os.remove("builds/EduTrack-linux64.zip")
if not os.path.exists("EduTrack-linux64"):
    os.mkdir("EduTrack-linux64")
shutil.copy2("builds/EduTrack_Linux","EduTrack-linux64")
os.system("zip -r builds/EduTrack-linux64.zip files EduTrack-linux64/")
shutil.rmtree("EduTrack-linux64")

  

