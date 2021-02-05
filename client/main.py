import socket
import requests
from warnings import warn
import json
from abc import ABC, abstractmethod

class Client(ABC):

    @abstractmethod
    def get_data(self, msg):
        pass

class SocketClient(Client):
    
    def __init__(self):
        self.host = "localhost"
        self.port = 8000


    def get_data(self, msg):
        self.client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        # connect
        try:
            self.client.connect((self.host, self.port))
        except Exception as e:
            warn(e)
            print("Connection Failed")
            return
        # send
        self.client.sendall(msg.encode())
        # receive
        byte = self.client.recv(4096)
        return byte.decode()

class RestClient(Client):
    
    def __init__(self):
        self.host = "localhost"
        self.port= "5000"

    def get_data(self, msg):
        if msg == "config":
            url = f"http://{self.host}:{self.port}/api/config"
        data = requests.get(url)
        if data.status_code != 200:
            print(data.text)
            return 
        else:
            data = data.json()
        return data


class ConfigClient:

    def __init__(self, conn_type: "REST"):
    
        if conn_type == "REST":
            self.conn_type = RestClient()

        else:  # SOCKET
            self.conn_type = SocketClient()
        
    def send_msg(self, msg):
        
        result = self.conn_type.get_data(msg)   
        return result


if __name__ == "__main__":

    print(ConfigClient("REST").send_msg("config"))
    print(ConfigClient("SOCKET").send_msg("config"))