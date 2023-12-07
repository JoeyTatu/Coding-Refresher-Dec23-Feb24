import cv2
import time
import os
import win32gui
import win32con

path = os.getcwd()
fourcc = 'MJPG'


def minimiseWindow():
    window = win32gui.GetForegroundWindow()
    win32gui.ShowWindow(window, win32con.SW_MINIMIZE)


def cctv():
    video = cv2.VideoCapture(200)
    video.set(3, 640)
    video.set(4, 480)
    width = video.get(3)
    height = video.get(4)
    # print("Video resolution is set to:", width, "by", height)
    print('''Help:
          1. Press ESC to quit
          2. Press M to minimise''')
    # Explicitly specify the codec
    # value

    date_time = time.strftime("Recording_%d-%b-%Y_%H:%M:%S")
    filename = "Recordings_" + date_time + ".mjpg"
    output = cv2.VideoWriter(os.path.join(
        path, filename), cv2.CAP_FFMPEG, 0, (640, 480))

    while video.isOpened():
        check, frame = video.read()
        if check == True:
            frame = cv2.flip(frame, 1)
            t = time.ctime()

            cv2.rectangle(frame, (5, 5, 100, 20), (255, 255, 255), cv2.FILLED)
            cv2.putText(frame, "Camera 1", (20, 20),
                        cv2.FONT_HERSHEY_DUPLEX, 0.5, (5, 5, 5), 1)
            cv2.putText(frame, t, (420, 460),
                        cv2.FONT_HERSHEY_DUPLEX, 0.5, (5, 5, 5), 1)
            cv2.imshow('CCTV Camera', frame)
            output.write(frame)

            if cv2.waitKey(1) == 27:
                print("Video recording saved in current directory")
                break
            elif cv2.waitKey(1) == ord("m"):
                minimiseWindow()
        else:
            print("Unable to open camera!")
            break

    video.release()
    output.release()
    cv2.destroyAllWindows()


print("*" * 80 + "\n" + " " * 30 + "Welcome to CCTV Software\n" + "*" * 80)
ask = int(input("Do you want to open the CCTV?\n1. Yes\n2. No\n>>>>>"))
if ask == 1:
    cctv()
elif ask == 2:
    print("Goodbye!")
    exit()
