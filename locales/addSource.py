import os

# This script adds an empty key for translation to all files in the .po folder.
# Reading from the last_add.txt file


def get_line(file_path, line_number):
    try:
        with open(file_path, 'r') as file:
            lines = file.readlines()
            if 1 <= line_number <= len(lines):
                specific_line = lines[line_number - 1].strip()
                return specific_line
            else:
                return f"Error: Line number {line_number} is out of range."
    except FileNotFoundError:
        return f"Error: File not found - {file_path}"
    except Exception as e:
        return f"An error occurred: {e}"


folder_path = "./po/"
files_in_folder = os.listdir(folder_path)


to_add_file = "last_add.txt"

src = get_line(to_add_file,1)
msgid = get_line(to_add_file,2)
text_to_append = f"""
#: {src}
msgid "{msgid}"
msgstr ""
"""


for file_name in files_in_folder:
    file_path = os.path.join(folder_path, file_name)
    if os.path.isfile(file_path):
        with open(file_path, 'a') as file:
            file.write(text_to_append + '\n')

