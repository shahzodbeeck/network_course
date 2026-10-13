from scapy.all import sniff
from scapy.layers.http import HTTPRequest
from scapy.packet import Raw


# sniff tutish uchun
# forge paket yasash
# attack hujum qilish
# def process_packet(packet):
#     if packet.haslayer(HTTPRequest):
#         url = packet[HTTPRequest].Host.decode() + packet[HTTPRequest].Path.decode()
#         ip = packet['IP'].src
#         method = packet[HTTPRequest].Method.decode()
#
#         print(f"Url {url} , source: {ip} , method: {method}")
#     if packet.haslayer(Raw):
#         load = packet[Raw].load.decode(errors='ignore')
#         keywords = ['username', 'login', 'user', 'password', 'pass']
#         for keyword in keywords:
#             if keyword in load:
#                 print(f"{keyword} : {load[keyword]}")
#
#
# print("[*] Sniffing boshlandi ...")
#
# sniff(iface="en0", prn=process_packet, store=False)
