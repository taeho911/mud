#!/usr/bin/env python3
# coding=utf-8

# reference: https://realpython.com/python-sockets/
# about GIL: https://realpython.com/python-gil/


import os, socket, selectors, types

HOST = os.getenv('ANAL_HOST') if os.getenv('ANAL_HOST') != None else '127.0.0.1'
PORT = int(os.getenv('ANAL_PORT')) if os.getenv('ANAL_PORT') != None else 19011
BUF_SIZE = 4096

def accept(sock):
    conn, addr = sock.accept()
    print('accepted connection from', addr)
    conn.setblocking(False)
    data = types.SimepleNamespace(addr=addr, inb=b'', outb=b'')
    events = selectors.EVENT_READ | selectors.EVENT_WRITE
    sel.register(conn, events, data=data)

def service_connection(key, mask):
    sock = key.fileobj
    data = key.data
    if mask & selectors.EVENT_READ:
        recv_data = sock.recv(4096)
        if recv_data:
            data.outb += recv_data
        else:
            print('closing connection to', data.addr)
            sel.unregister(sock)
            sock.close()
    if mask & selectors.EVENT_WRITE:
        if data.outb:
            print('echoing', repr(data.outb), 'to', data.addr)
            sent = sock.send(data.outb)
            data.outb = data.outb[sent:]

if __name__ == '__main__':
    sel = selectors.DefaultSelector()
    lsock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    lsock.bind((HOST, PORT))
    lsock.listen()

    print('listening on', (HOST, PORT))
    lsock.setblocking(False)
    sel.register(lsock, selectors.EVENT_READ, data=None)

    while True:
        events = sel.select(timeout=None)
        for key, mask in events:
            accept(key.fileobj)
        else:
            service_connection(key, mask)

    # with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as soc:
    #     soc.bind((HOST, PORT))
    #     soc.listen()
    #     conn, addr = soc.accept()
    #     with conn:
    #         print(f'Connected by {addr}')
    #         while True:
    #             data = conn.recv(BUF_SIZE)
    #             if not data:
    #                 break
    #             conn.sendall(data)
