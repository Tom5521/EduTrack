'''
  Copyright Tom5521(c) - All Rights Reserved.
 
  This project is licensed under the MIT License.
'''

import os

def MakeBundle():

  # Define the route for the icon bundle file.
  route = "./iconloader/bundle.go"

  # Function to regenerate an icon bundle.
  def ReGenerateBundle(name, file):
      os.system(f"fyne bundle --pkg iconloader --name {name} -o {route} {file}")

  # Function to append icons to the existing bundle.
  def AppendBundle(name, file):
      os.system(f"fyne bundle --pkg iconloader --name {name} -o {route} -append {file}")

  # Regenerate the "Placeholder" icon bundle.
  ReGenerateBundle("Placeholder", "./Assets/Placeholder.png")

  # Append other icons to the bundle.

  # App Icon
  AppendBundle("App_light", "./Assets/Icon.png")
  AppendBundle("App_dark","./Assets/Icon-Dark.png")

  # Dark Icons
  AppendBundle("Save_Dark", "./Assets/Icons-Dark/Save.png")
  AppendBundle("Dev_Dark", "./Assets/Icons-Dark/Dev.png")
  AppendBundle("Install_Dark", "./Assets/Icons-Dark/Install.png")
  AppendBundle("Info_Dark", "./Assets/Icons-Dark/Info.png")
  AppendBundle("Error_Dark", "./Assets/Icons-Dark/Error.png")
  AppendBundle("Restart_Dark", "./Assets/Icons-Dark/Restart.png")
  AppendBundle("Download_Dark", "./Assets/Icons-Dark/Download.png")
  AppendBundle("Uninstall_Dark", "./Assets/Icons-Dark/Uninstall.png")
  AppendBundle("TemplateUser_Dark", "./Assets/Icons-Dark/UserTemplate.png")

  # Light Icons
  AppendBundle("Save_Light", "./Assets/Icons-Light/Save.png")
  AppendBundle("Dev_Light", "./Assets/Icons-Light/Dev.png")
  AppendBundle("Install_Light", "./Assets/Icons-Light/Install.png")
  AppendBundle("Info_Light", "./Assets/Icons-Light/Info.png")
  AppendBundle("Error_Light", "./Assets/Icons-Light/Error.png")
  AppendBundle("Restart_Light", "./Assets/Icons-Light/Restart.png")
  AppendBundle("Download_Light", "./Assets/Icons-Light/Download.png")
  AppendBundle("Uninstall_Light", "./Assets/Icons-Light/Uninstall.png")
  AppendBundle("TemplateUser_Light", "./Assets/Icons-Light/UserTemplate.png")

if __name__ == "__main__":
   MakeBundle() 

