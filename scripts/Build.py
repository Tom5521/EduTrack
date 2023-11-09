'''
  Copyright Tom5521(c) - All Rights Reserved.
 
  This project is licensed under the MIT License.
'''

import os
import platform
import shutil
import sys

OS = platform.system()

def BuildForWindows():
    if not os.path.exists("builds"):
        os.mkdir("builds")

    # Check the platform (Windows or not).
    if OS != "Windows":
        os.environ["GGO_ENABLED"] = "1"
        os.environ["CC"] = "/usr/bin/x86_64-w64-mingw32-gcc"
        os.environ["CXX"] = "/usr/bin/x86_64-w64-mingw32-c++"
        os.environ["GOOS"] = "windows"

    # Package the application for Windows using fyne package.
    os.system("fyne package --os windows --exe builds/EduTrack.exe --release")


def BuildForLinux():
    if not os.path.exists("builds"):
        os.mkdir("builds")

    if OS != "Windows":
        os.system("fyne package --os linux --release")
    else:
        os.system("sudo fyne-cross -os linux -release")

    if os.path.exists("EduTrack.tar.xz"):
        shutil.move("EduTrack.tar.xz","builds/EduTrack-linux64.tar.xz")

if __name__ == "__main__":
    if len(sys.argv) == 1:
        exit()
    elif sys.argv[1] == "win":
        print("Compiling for windows...")
        BuildForWindows()
    elif sys.argv[1] == "linux":
        print("Compiling for linux...")
        BuildForLinux()
    
