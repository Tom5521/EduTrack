'''
  Copyright Tom5521(c) - All Rights Reserved.
 
  This project is licensed under the MIT License.
'''

import platform
import os
import shutil

def BuildForLinux():
    OS = platform.system()

    if not os.path.exists("builds"):
        os.mkdir("builds")

    if OS != "Windows":
        os.system("fyne package --os linux --release")
    else:
        os.system("sudo fyne-cross -os linux")

    if os.path.exists("EduTrack.tar.xz"):
        shutil.move("EduTrack.tar.xz","builds/EduTrack-linux64.tar.xz")

if __name__ == "__main__":
    BuildForLinux()
    
