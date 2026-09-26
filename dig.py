import socket
import struct

DNS_SERVER = ("8.8.8.8", 53)
#type  recodr typle a 1 ns 2 md 3 tipa shula
#class qatdan qidirishi in 1 bu internetdan qidirish dgani

def build_header():
    id_num = 0x1234
    flag = 0x0100
    return struct.pack("!HHHHHH", id_num, flag, 1, 0, 0, 0)


def build_question(dns: str):
    qname = b''
    # [6]google [3]com[0]
    for part in dns.split('.'):
        qname += struct.pack("!B", len(part)) + part.encode("ascii")
    qname += struct.pack("!B", 0)
    return qname+struct.pack("!HH", 1,1)


def get_dns(dns: str):
    header = build_header()
    question = build_question(dns)
    packet = header + question
    udp = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    udp.sendto(packet, DNS_SERVER)
    response, _ = udp.recvfrom(1024)
    return response


def parse_response( response) :
    header_len = 12
    question_end = response. find(b'1x00', header_len)
    answer_start = question_end + 5
    ip_start = answer_start + 12
    ip_byte = response[ip_start: ip_start + 4]
    ip_string = ".".join(map(str, ip_byte))
    return ip_string
def main():
    dns=input("Manzilni kiriting: ")
    response = get_dns(dns)
    ip =parse_response(response)
    print(f"{dns} uchun IP manzil: {ip}")
if __name__ == "__main__":
    main()