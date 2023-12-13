'''
  Copyright Tom5521(c) - All Rights Reserved.
 
  This project is licensed under the MIT License.
'''

import platform
import Build
import MakeWinZip
import os

def MakeRelease():
    print("Compiling for linux...")
    Build.BuildForLinux()
    print("Compiling for windows...")
    Build.BuildForWindows()
    print("Making windows zip")
    MakeWinZip.MakeWinZip()
    if platform.system() == "unix" or "linux":
        os.system("notify-send \"Release compilation ended susefully\"")

if __name__ == "__main__":
    MakeRelease()

