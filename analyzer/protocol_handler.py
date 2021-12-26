#!/usr/bin/env python3
# coding=utf-8

# Reference: https://dev.to/pybash/making-a-custom-protocol-handler-and-uri-scheme-part-1-37mh

import urllib.request as urlreq

class pymudHandler(urlreq.BaseHandler):
    def pymud_open(self, req):
        url = req.get_full_url()
        print('pymud protocol requested')
        print(url)
        return url

if __name__ == '__main__':
    opener = urlreq.build_opener(pymudHandler)
    urlreq.install_opener(opener)
    urlreq.urlopen('pymud://hello')
