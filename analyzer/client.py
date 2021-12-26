#!/usr/bin/env python3
# coding=utf-8

import os, socket

HOST = os.getenv('ANAL_HOST') if os.getenv('ANAL_HOST') != None else '127.0.0.1'
PORT = int(os.getenv('ANAL_PORT')) if os.getenv('ANAL_PORT') != None else 8081
BUF_SIZE = 4096

if __name__ == '__main__':
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as soc:
        soc.connect((HOST, PORT))
        soc.sendall(b'Hello world')
        data = soc.recv(BUF_SIZE)
    print(f'Recieved {data}')
