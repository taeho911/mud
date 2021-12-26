#!/usr/bin/env python3
# coding=utf-8

# reference: https://realpython.com/python-sockets/
# about GIL: https://realpython.com/python-gil/


import os, socket

HOST = os.getenv('ANAL_HOST') if os.getenv('ANAL_HOST') != None else '127.0.0.1'
PORT = int(os.getenv('ANAL_PORT')) if os.getenv('ANAL_PORT') != None else 8081
BUF_SIZE = 4096

if __name__ == '__main__':
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as soc:
        soc.bind((HOST, PORT))
        soc.listen()
        conn, addr = soc.accept()
        with conn:
            print(f'Connected by {addr}')
            while True:
                data = conn.recv(BUF_SIZE)
                if not data:
                    break
                conn.sendall(data)
