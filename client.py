import time
import RPi.GPIO as GPIO
import socket

target_ip = "192.168.1.9"
target_port = 5000
buffer_size = 4096

GPIO.setmode(GPIO.BCM)

GPIO.setup(18, GPIO.IN, pull_up_down=GPIO.PUD_UP)

udp_client = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

sw_status = 1

while True:
    try:
        sw_status = GPIO.input(18)
            if sw_status == 0:
            # close

            else:
            # open
            udp_client.sendto(b'OpenTheDoor', (target_ip, target_port))
            while sw_status == 1:
            sw_status = GPIO.input(18)
            time.sleep(1)

    except:
        udp_client.close()
        break

GPIO.cleanup()
print("end")
